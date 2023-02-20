// Code generated by MockGen. DO NOT EDIT.
// Source: ./events.go

// Package calendar is a generated GoMock package.
package calendar

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

// Create mocks base method.
func (m *MockRepository) Create(event model.Event) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", event)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockRepositoryMockRecorder) Create(event interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockRepository)(nil).Create), event)
}

// MockCalendar is a mock of Calendar interface.
type MockCalendar struct {
	ctrl     *gomock.Controller
	recorder *MockCalendarMockRecorder
}

// MockCalendarMockRecorder is the mock recorder for MockCalendar.
type MockCalendarMockRecorder struct {
	mock *MockCalendar
}

// NewMockCalendar creates a new mock instance.
func NewMockCalendar(ctrl *gomock.Controller) *MockCalendar {
	mock := &MockCalendar{ctrl: ctrl}
	mock.recorder = &MockCalendarMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCalendar) EXPECT() *MockCalendarMockRecorder {
	return m.recorder
}

// CreateEvent mocks base method.
func (m *MockCalendar) CreateEvent(event model.Event) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateEvent", event)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateEvent indicates an expected call of CreateEvent.
func (mr *MockCalendarMockRecorder) CreateEvent(event interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateEvent", reflect.TypeOf((*MockCalendar)(nil).CreateEvent), event)
}