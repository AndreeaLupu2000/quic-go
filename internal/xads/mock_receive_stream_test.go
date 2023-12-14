// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/quic-go/quic-go/internal/xads (interfaces: ReceiveStream)

// Package xads is a generated GoMock package.
package xads

import (
	reflect "reflect"
	time "time"

	gomock "github.com/golang/mock/gomock"
	protocol "github.com/quic-go/quic-go/internal/protocol"
	qerr "github.com/quic-go/quic-go/internal/qerr"
	wire "github.com/quic-go/quic-go/internal/wire"
)

// MockReceiveStream is a mock of ReceiveStream interface.
type MockReceiveStream struct {
	ctrl     *gomock.Controller
	recorder *MockReceiveStreamMockRecorder
}

// MockReceiveStreamMockRecorder is the mock recorder for MockReceiveStream.
type MockReceiveStreamMockRecorder struct {
	mock *MockReceiveStream
}

// NewMockReceiveStream creates a new mock instance.
func NewMockReceiveStream(ctrl *gomock.Controller) *MockReceiveStream {
	mock := &MockReceiveStream{ctrl: ctrl}
	mock.recorder = &MockReceiveStreamMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockReceiveStream) EXPECT() *MockReceiveStreamMockRecorder {
	return m.recorder
}

// CancelRead mocks base method.
func (m *MockReceiveStream) CancelRead(arg0 qerr.StreamErrorCode) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "CancelRead", arg0)
}

// CancelRead indicates an expected call of CancelRead.
func (mr *MockReceiveStreamMockRecorder) CancelRead(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CancelRead", reflect.TypeOf((*MockReceiveStream)(nil).CancelRead), arg0)
}

// CloseForShutdown mocks base method.
func (m *MockReceiveStream) CloseForShutdown(arg0 error) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "CloseForShutdown", arg0)
}

// CloseForShutdown indicates an expected call of CloseForShutdown.
func (mr *MockReceiveStreamMockRecorder) CloseForShutdown(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloseForShutdown", reflect.TypeOf((*MockReceiveStream)(nil).CloseForShutdown), arg0)
}

// GetWindowUpdate mocks base method.
func (m *MockReceiveStream) GetWindowUpdate() protocol.ByteCount {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetWindowUpdate")
	ret0, _ := ret[0].(protocol.ByteCount)
	return ret0
}

// GetWindowUpdate indicates an expected call of GetWindowUpdate.
func (mr *MockReceiveStreamMockRecorder) GetWindowUpdate() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWindowUpdate", reflect.TypeOf((*MockReceiveStream)(nil).GetWindowUpdate))
}

// HandleResetStreamFrame mocks base method.
func (m *MockReceiveStream) HandleResetStreamFrame(arg0 *wire.ResetStreamFrame) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HandleResetStreamFrame", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// HandleResetStreamFrame indicates an expected call of HandleResetStreamFrame.
func (mr *MockReceiveStreamMockRecorder) HandleResetStreamFrame(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HandleResetStreamFrame", reflect.TypeOf((*MockReceiveStream)(nil).HandleResetStreamFrame), arg0)
}

// HandleStreamFrame mocks base method.
func (m *MockReceiveStream) HandleStreamFrame(arg0 *wire.StreamFrame) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HandleStreamFrame", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// HandleStreamFrame indicates an expected call of HandleStreamFrame.
func (mr *MockReceiveStreamMockRecorder) HandleStreamFrame(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HandleStreamFrame", reflect.TypeOf((*MockReceiveStream)(nil).HandleStreamFrame), arg0)
}

// Read mocks base method.
func (m *MockReceiveStream) Read(arg0 []byte) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Read", arg0)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Read indicates an expected call of Read.
func (mr *MockReceiveStreamMockRecorder) Read(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Read", reflect.TypeOf((*MockReceiveStream)(nil).Read), arg0)
}

// SetReadDeadline mocks base method.
func (m *MockReceiveStream) SetReadDeadline(arg0 time.Time) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetReadDeadline", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetReadDeadline indicates an expected call of SetReadDeadline.
func (mr *MockReceiveStreamMockRecorder) SetReadDeadline(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetReadDeadline", reflect.TypeOf((*MockReceiveStream)(nil).SetReadDeadline), arg0)
}

// StreamID mocks base method.
func (m *MockReceiveStream) StreamID() protocol.StreamID {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StreamID")
	ret0, _ := ret[0].(protocol.StreamID)
	return ret0
}

// StreamID indicates an expected call of StreamID.
func (mr *MockReceiveStreamMockRecorder) StreamID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StreamID", reflect.TypeOf((*MockReceiveStream)(nil).StreamID))
}
