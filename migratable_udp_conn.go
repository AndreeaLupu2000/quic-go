package quic

import (
	"net"
	"os"
	"reflect"
	"syscall"
	"time"
)

// check at compile time that interfaces satisfies interfaces
var _ net.PacketConn = &MigratableUDPConn{}
var _ interface {
	SyscallConn() (syscall.RawConn, error)
} = &MigratableUDPConn{}
var _ interface{ SetReadBuffer(int) error } = &MigratableUDPConn{}

// MigratableUDPConn
//
// Packet connection that supports migration of IP address and UDP port
type MigratableUDPConn struct {
	internal   *net.UDPConn
	maxRetries int
	// to restore after
	deadline      *time.Time
	readDeadline  *time.Time
	writeDeadline *time.Time
	readBuffer    *int
}

func ListenMigratableUDP(network string, laddr *net.UDPAddr) (*MigratableUDPConn, error) {
	conn, err := net.ListenUDP(network, laddr)
	if err != nil {
		return nil, err
	}
	return &MigratableUDPConn{
		internal:   conn,
		maxRetries: 5,
	}, nil
}

func (m *MigratableUDPConn) ReadFrom(p []byte) (n int, addr net.Addr, err error) {
	retryCount := 0
	for {
		n, addr, err = m.internal.ReadFrom(p)
		if err != nil {
			err = m.handleError(err, &retryCount)
			if err != nil {
				return n, addr, err
			}
			continue // retry
		}
		return n, addr, err
	}
}

func (m *MigratableUDPConn) WriteTo(p []byte, addr net.Addr) (n int, err error) {
	retryCount := 0
	for {
		n, err = m.internal.WriteTo(p, addr)
		if err != nil {
			err = m.handleError(err, &retryCount)
			if err != nil {
				return n, err
			}
			continue // retry
		}
		return n, err
	}
}

// If error is returned it should no longer be retried
func (m *MigratableUDPConn) handleError(err error, retryCount *int) error {
	switch err := err.(type) {
	case *net.OpError:
		switch err := err.Err.(type) {
		case *os.SyscallError:
			switch err := err.Err.(type) {
			case syscall.Errno:
				if err.Error() == "network is unreachable" {
					// reopen and retry, because it could be caused of migration
					if *retryCount < m.maxRetries {
						err := m.Reopen()
						if err != nil {
							return err
						}
						return nil
					}
				}
			}
		default:
			// use reflect, because type is not public
			if reflect.TypeOf(err).String() == "poll.errNetClosing" {
				// retry, because it could be caused of migration
				if *retryCount < m.maxRetries {
					// give socket migration some time
					time.Sleep(10 * time.Millisecond)
					*retryCount++
					return nil
				}
			}
		}
	}
	return err
}

func (m *MigratableUDPConn) Close() error {
	return m.internal.Close()
}

func (m *MigratableUDPConn) LocalAddr() net.Addr {
	return m.internal.LocalAddr()
}

func (m *MigratableUDPConn) SetDeadline(t time.Time) error {
	m.deadline = &t
	return m.internal.SetDeadline(t)
}

func (m *MigratableUDPConn) SetReadDeadline(t time.Time) error {
	m.readDeadline = &t
	return m.internal.SetReadDeadline(t)
}

func (m *MigratableUDPConn) SetWriteDeadline(t time.Time) error {
	m.writeDeadline = &t
	return m.internal.SetWriteDeadline(t)
}

func (m *MigratableUDPConn) SetReadBuffer(bytes int) error {
	m.readBuffer = &bytes
	return m.internal.SetReadBuffer(bytes)
}

func (m *MigratableUDPConn) SyscallConn() (syscall.RawConn, error) {
	return m.internal.SyscallConn()
}

// Reopen new UDP socket on same address
func (m *MigratableUDPConn) Reopen() error {
	err := m.internal.Close()
	if err != nil {
		return err
	}
	conn, err := net.ListenUDP("udp", m.internal.LocalAddr().(*net.UDPAddr))
	if err != nil {
		return err
	}
	m.internal = conn
	err = m.applyConfig()
	if err != nil {
		return err
	}
	return nil
}

// Migrate connection to new UDP socket.
// Returns new UDP address.
func (m *MigratableUDPConn) Migrate() (*net.UDPAddr, error) {
	err := m.internal.Close()
	if err != nil {
		return nil, err
	}

	conn, err := net.ListenUDP("udp", &net.UDPAddr{IP: net.IPv4zero, Port: 0})
	if err != nil {
		return nil, err
	}
	m.internal = conn
	err = m.applyConfig()
	if err != nil {
		return nil, err
	}
	return conn.LocalAddr().(*net.UDPAddr), nil
}

// apply config to the current internal UDP connection
func (m *MigratableUDPConn) applyConfig() error {
	if m.deadline != nil {
		err := m.internal.SetDeadline(*m.deadline)
		if err != nil {
			return err
		}
	}
	if m.readDeadline != nil {
		err := m.internal.SetReadDeadline(*m.readDeadline)
		if err != nil {
			return err
		}
	}
	if m.writeDeadline != nil {
		err := m.internal.SetWriteDeadline(*m.writeDeadline)
		if err != nil {
			return err
		}
	}
	if m.readBuffer != nil {
		err := m.internal.SetReadBuffer(*m.readBuffer)
		if err != nil {
			return err
		}
	}
	return nil
}
