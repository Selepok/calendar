// Code generated by MockGen. DO NOT EDIT.
// Source: user.go

// Package user is a generated GoMock package.
package user

import (
	reflect "reflect"

	model "github.com/Selepok/calendar/internal/model"
	gomock "github.com/golang/mock/gomock"
)

// MockRepository is a mock of Repository interface.
type MockRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryMockRecorder
}

// MockRepositoryMockRecorder is the mock recorder for MockRepository.
type MockRepositoryMockRecorder struct {
	mock *MockRepository
}

// NewMockRepository creates a new mock instance.
func NewMockRepository(ctrl *gomock.Controller) *MockRepository {
	mock := &MockRepository{ctrl: ctrl}
	mock.recorder = &MockRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepository) EXPECT() *MockRepositoryMockRecorder {
	return m.recorder
}

// CreateUser mocks base method.
func (m *MockRepository) CreateUser(login, password, timezone string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", login, password, timezone)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockRepositoryMockRecorder) CreateUser(login, password, timezone interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockRepository)(nil).CreateUser), login, password, timezone)
}

// GetUserHashedPassword mocks base method.
func (m *MockRepository) GetUserHashedPassword(login string) (int, string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserHashedPassword", login)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetUserHashedPassword indicates an expected call of GetUserHashedPassword.
func (mr *MockRepositoryMockRecorder) GetUserHashedPassword(login interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserHashedPassword", reflect.TypeOf((*MockRepository)(nil).GetUserHashedPassword), login)
}

// GetUserIdByLogin mocks base method.
func (m *MockRepository) GetUserIdByLogin(login string) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserIdByLogin", login)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserIdByLogin indicates an expected call of GetUserIdByLogin.
func (mr *MockRepositoryMockRecorder) GetUserIdByLogin(login interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserIdByLogin", reflect.TypeOf((*MockRepository)(nil).GetUserIdByLogin), login)
}

// Update mocks base method.
func (m *MockRepository) Update(user model.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", user)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockRepositoryMockRecorder) Update(user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockRepository)(nil).Update), user)
}
