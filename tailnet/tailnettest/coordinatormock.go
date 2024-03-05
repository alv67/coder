// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/coder/coder/v2/tailnet (interfaces: Coordinator)
//
// Generated by this command:
//
//	mockgen -destination ./coordinatormock.go -package tailnettest github.com/coder/coder/v2/tailnet Coordinator
//

// Package tailnettest is a generated GoMock package.
package tailnettest

import (
	context "context"
	net "net"
	http "net/http"
	reflect "reflect"

	tailnet "github.com/coder/coder/v2/tailnet"
	proto "github.com/coder/coder/v2/tailnet/proto"
	uuid "github.com/google/uuid"
	gomock "go.uber.org/mock/gomock"
)

// MockCoordinator is a mock of Coordinator interface.
type MockCoordinator struct {
	ctrl     *gomock.Controller
	recorder *MockCoordinatorMockRecorder
}

// MockCoordinatorMockRecorder is the mock recorder for MockCoordinator.
type MockCoordinatorMockRecorder struct {
	mock *MockCoordinator
}

// NewMockCoordinator creates a new mock instance.
func NewMockCoordinator(ctrl *gomock.Controller) *MockCoordinator {
	mock := &MockCoordinator{ctrl: ctrl}
	mock.recorder = &MockCoordinatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCoordinator) EXPECT() *MockCoordinatorMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockCoordinator) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockCoordinatorMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockCoordinator)(nil).Close))
}

// Coordinate mocks base method.
func (m *MockCoordinator) Coordinate(arg0 context.Context, arg1 uuid.UUID, arg2 string, arg3 tailnet.CoordinateeAuth) (chan<- *proto.CoordinateRequest, <-chan *proto.CoordinateResponse) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Coordinate", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(chan<- *proto.CoordinateRequest)
	ret1, _ := ret[1].(<-chan *proto.CoordinateResponse)
	return ret0, ret1
}

// Coordinate indicates an expected call of Coordinate.
func (mr *MockCoordinatorMockRecorder) Coordinate(arg0, arg1, arg2, arg3 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Coordinate", reflect.TypeOf((*MockCoordinator)(nil).Coordinate), arg0, arg1, arg2, arg3)
}

// Node mocks base method.
func (m *MockCoordinator) Node(arg0 uuid.UUID) *tailnet.Node {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Node", arg0)
	ret0, _ := ret[0].(*tailnet.Node)
	return ret0
}

// Node indicates an expected call of Node.
func (mr *MockCoordinatorMockRecorder) Node(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Node", reflect.TypeOf((*MockCoordinator)(nil).Node), arg0)
}

// ServeAgent mocks base method.
func (m *MockCoordinator) ServeAgent(arg0 net.Conn, arg1 uuid.UUID, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ServeAgent", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// ServeAgent indicates an expected call of ServeAgent.
func (mr *MockCoordinatorMockRecorder) ServeAgent(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ServeAgent", reflect.TypeOf((*MockCoordinator)(nil).ServeAgent), arg0, arg1, arg2)
}

// ServeClient mocks base method.
func (m *MockCoordinator) ServeClient(arg0 net.Conn, arg1, arg2 uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ServeClient", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// ServeClient indicates an expected call of ServeClient.
func (mr *MockCoordinatorMockRecorder) ServeClient(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ServeClient", reflect.TypeOf((*MockCoordinator)(nil).ServeClient), arg0, arg1, arg2)
}

// ServeHTTPDebug mocks base method.
func (m *MockCoordinator) ServeHTTPDebug(arg0 http.ResponseWriter, arg1 *http.Request) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ServeHTTPDebug", arg0, arg1)
}

// ServeHTTPDebug indicates an expected call of ServeHTTPDebug.
func (mr *MockCoordinatorMockRecorder) ServeHTTPDebug(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ServeHTTPDebug", reflect.TypeOf((*MockCoordinator)(nil).ServeHTTPDebug), arg0, arg1)
}

// ServeMultiAgent mocks base method.
func (m *MockCoordinator) ServeMultiAgent(arg0 uuid.UUID) tailnet.MultiAgentConn {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ServeMultiAgent", arg0)
	ret0, _ := ret[0].(tailnet.MultiAgentConn)
	return ret0
}

// ServeMultiAgent indicates an expected call of ServeMultiAgent.
func (mr *MockCoordinatorMockRecorder) ServeMultiAgent(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ServeMultiAgent", reflect.TypeOf((*MockCoordinator)(nil).ServeMultiAgent), arg0)
}
