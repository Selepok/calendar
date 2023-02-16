// Code generated by MockGen. DO NOT EDIT.
// Source: ./jwt.go

// Package auth is a generated GoMock package.
package auth

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockTokenAuthentication is a mock of TokenAuthentication interface.
type MockTokenAuthentication struct {
	ctrl     *gomock.Controller
	recorder *MockTokenAuthenticationMockRecorder
}

// MockTokenAuthenticationMockRecorder is the mock recorder for MockTokenAuthentication.
type MockTokenAuthenticationMockRecorder struct {
	mock *MockTokenAuthentication
}

// NewMockTokenAuthentication creates a new mock instance.
func NewMockTokenAuthentication(ctrl *gomock.Controller) *MockTokenAuthentication {
	mock := &MockTokenAuthentication{ctrl: ctrl}
	mock.recorder = &MockTokenAuthenticationMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTokenAuthentication) EXPECT() *MockTokenAuthenticationMockRecorder {
	return m.recorder
}

// GenerateToken mocks base method.
func (m *MockTokenAuthentication) GenerateToken(arg0 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateToken", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GenerateToken indicates an expected call of GenerateToken.
func (mr *MockTokenAuthenticationMockRecorder) GenerateToken(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateToken", reflect.TypeOf((*MockTokenAuthentication)(nil).GenerateToken), arg0)
}

// ValidateToken mocks base method.
func (m *MockTokenAuthentication) ValidateToken(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidateToken", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// ValidateToken indicates an expected call of ValidateToken.
func (mr *MockTokenAuthenticationMockRecorder) ValidateToken(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidateToken", reflect.TypeOf((*MockTokenAuthentication)(nil).ValidateToken), arg0)
}
