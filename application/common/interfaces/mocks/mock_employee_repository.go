// Code generated by MockGen. DO NOT EDIT.
// Source: employee_repository.go

// Package mocks is a generated GoMock package.
package mocks

import (
	entities "go-clean-architecture-sample/domain/entities"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockEmployeeRepository is a mock of EmployeeRepository interface.
type MockEmployeeRepository struct {
	ctrl     *gomock.Controller
	recorder *MockEmployeeRepositoryMockRecorder
}

// MockEmployeeRepositoryMockRecorder is the mock recorder for MockEmployeeRepository.
type MockEmployeeRepositoryMockRecorder struct {
	mock *MockEmployeeRepository
}

// NewMockEmployeeRepository creates a new mock instance.
func NewMockEmployeeRepository(ctrl *gomock.Controller) *MockEmployeeRepository {
	mock := &MockEmployeeRepository{ctrl: ctrl}
	mock.recorder = &MockEmployeeRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEmployeeRepository) EXPECT() *MockEmployeeRepositoryMockRecorder {
	return m.recorder
}

// Add mocks base method.
func (m *MockEmployeeRepository) Add(employee entities.Employee) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Add", employee)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Add indicates an expected call of Add.
func (mr *MockEmployeeRepositoryMockRecorder) Add(employee interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockEmployeeRepository)(nil).Add), employee)
}

// DeleteById mocks base method.
func (m *MockEmployeeRepository) DeleteById(id int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteById", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteById indicates an expected call of DeleteById.
func (mr *MockEmployeeRepositoryMockRecorder) DeleteById(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteById", reflect.TypeOf((*MockEmployeeRepository)(nil).DeleteById), id)
}

// GetAll mocks base method.
func (m *MockEmployeeRepository) GetAll() ([]entities.Employee, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll")
	ret0, _ := ret[0].([]entities.Employee)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll.
func (mr *MockEmployeeRepositoryMockRecorder) GetAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockEmployeeRepository)(nil).GetAll))
}

// GetById mocks base method.
func (m *MockEmployeeRepository) GetById(id int) (*entities.Employee, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetById", id)
	ret0, _ := ret[0].(*entities.Employee)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetById indicates an expected call of GetById.
func (mr *MockEmployeeRepositoryMockRecorder) GetById(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetById", reflect.TypeOf((*MockEmployeeRepository)(nil).GetById), id)
}
