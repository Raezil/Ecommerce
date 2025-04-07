// Code generated by MockGen. DO NOT EDIT.
// Source: authenticator_grpc.pb.go

// Package mock_generated is a generated GoMock package.
package generated

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	grpc "google.golang.org/grpc"
)

// MockAuthClient is a mock of AuthClient interface.
type MockAuthClient struct {
	ctrl     *gomock.Controller
	recorder *MockAuthClientMockRecorder
}

// MockAuthClientMockRecorder is the mock recorder for MockAuthClient.
type MockAuthClientMockRecorder struct {
	mock *MockAuthClient
}

// NewMockAuthClient creates a new mock instance.
func NewMockAuthClient(ctrl *gomock.Controller) *MockAuthClient {
	mock := &MockAuthClient{ctrl: ctrl}
	mock.recorder = &MockAuthClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAuthClient) EXPECT() *MockAuthClientMockRecorder {
	return m.recorder
}

// Login mocks base method.
func (m *MockAuthClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginReply, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Login", varargs...)
	ret0, _ := ret[0].(*LoginReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Login indicates an expected call of Login.
func (mr *MockAuthClientMockRecorder) Login(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Login", reflect.TypeOf((*MockAuthClient)(nil).Login), varargs...)
}

// Register mocks base method.
func (m *MockAuthClient) Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterReply, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Register", varargs...)
	ret0, _ := ret[0].(*RegisterReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Register indicates an expected call of Register.
func (mr *MockAuthClientMockRecorder) Register(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Register", reflect.TypeOf((*MockAuthClient)(nil).Register), varargs...)
}

// SampleProtected mocks base method.
func (m *MockAuthClient) SampleProtected(ctx context.Context, in *ProtectedRequest, opts ...grpc.CallOption) (*ProtectedReply, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SampleProtected", varargs...)
	ret0, _ := ret[0].(*ProtectedReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SampleProtected indicates an expected call of SampleProtected.
func (mr *MockAuthClientMockRecorder) SampleProtected(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SampleProtected", reflect.TypeOf((*MockAuthClient)(nil).SampleProtected), varargs...)
}

// MockAuthServer is a mock of AuthServer interface.
type MockAuthServer struct {
	ctrl     *gomock.Controller
	recorder *MockAuthServerMockRecorder
}

// MockAuthServerMockRecorder is the mock recorder for MockAuthServer.
type MockAuthServerMockRecorder struct {
	mock *MockAuthServer
}

// NewMockAuthServer creates a new mock instance.
func NewMockAuthServer(ctrl *gomock.Controller) *MockAuthServer {
	mock := &MockAuthServer{ctrl: ctrl}
	mock.recorder = &MockAuthServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAuthServer) EXPECT() *MockAuthServerMockRecorder {
	return m.recorder
}

// Login mocks base method.
func (m *MockAuthServer) Login(arg0 context.Context, arg1 *LoginRequest) (*LoginReply, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Login", arg0, arg1)
	ret0, _ := ret[0].(*LoginReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Login indicates an expected call of Login.
func (mr *MockAuthServerMockRecorder) Login(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Login", reflect.TypeOf((*MockAuthServer)(nil).Login), arg0, arg1)
}

// Register mocks base method.
func (m *MockAuthServer) Register(arg0 context.Context, arg1 *RegisterRequest) (*RegisterReply, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Register", arg0, arg1)
	ret0, _ := ret[0].(*RegisterReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Register indicates an expected call of Register.
func (mr *MockAuthServerMockRecorder) Register(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Register", reflect.TypeOf((*MockAuthServer)(nil).Register), arg0, arg1)
}

// SampleProtected mocks base method.
func (m *MockAuthServer) SampleProtected(arg0 context.Context, arg1 *ProtectedRequest) (*ProtectedReply, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SampleProtected", arg0, arg1)
	ret0, _ := ret[0].(*ProtectedReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SampleProtected indicates an expected call of SampleProtected.
func (mr *MockAuthServerMockRecorder) SampleProtected(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SampleProtected", reflect.TypeOf((*MockAuthServer)(nil).SampleProtected), arg0, arg1)
}

// mustEmbedUnimplementedAuthServer mocks base method.
func (m *MockAuthServer) mustEmbedUnimplementedAuthServer() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "mustEmbedUnimplementedAuthServer")
}

// mustEmbedUnimplementedAuthServer indicates an expected call of mustEmbedUnimplementedAuthServer.
func (mr *MockAuthServerMockRecorder) mustEmbedUnimplementedAuthServer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "mustEmbedUnimplementedAuthServer", reflect.TypeOf((*MockAuthServer)(nil).mustEmbedUnimplementedAuthServer))
}

// MockUnsafeAuthServer is a mock of UnsafeAuthServer interface.
type MockUnsafeAuthServer struct {
	ctrl     *gomock.Controller
	recorder *MockUnsafeAuthServerMockRecorder
}

// MockUnsafeAuthServerMockRecorder is the mock recorder for MockUnsafeAuthServer.
type MockUnsafeAuthServerMockRecorder struct {
	mock *MockUnsafeAuthServer
}

// NewMockUnsafeAuthServer creates a new mock instance.
func NewMockUnsafeAuthServer(ctrl *gomock.Controller) *MockUnsafeAuthServer {
	mock := &MockUnsafeAuthServer{ctrl: ctrl}
	mock.recorder = &MockUnsafeAuthServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUnsafeAuthServer) EXPECT() *MockUnsafeAuthServerMockRecorder {
	return m.recorder
}

// mustEmbedUnimplementedAuthServer mocks base method.
func (m *MockUnsafeAuthServer) mustEmbedUnimplementedAuthServer() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "mustEmbedUnimplementedAuthServer")
}

// mustEmbedUnimplementedAuthServer indicates an expected call of mustEmbedUnimplementedAuthServer.
func (mr *MockUnsafeAuthServerMockRecorder) mustEmbedUnimplementedAuthServer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "mustEmbedUnimplementedAuthServer", reflect.TypeOf((*MockUnsafeAuthServer)(nil).mustEmbedUnimplementedAuthServer))
}
