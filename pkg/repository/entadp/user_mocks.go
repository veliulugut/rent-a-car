// Code generated by MockGen. DO NOT EDIT.
// Source: rent-a-car/pkg/repository/entadp (interfaces: UserRepositoryInterface)

// Package entadp is a generated GoMock package.
package entadp

import (
	context "context"
	reflect "reflect"
	dto "rent-a-car/pkg/repository/dto"

	gomock "github.com/golang/mock/gomock"
)

// MockUserRepositoryInterface is a mock of UserRepositoryInterface interface.
type MockUserRepositoryInterface struct {
	ctrl     *gomock.Controller
	recorder *MockUserRepositoryInterfaceMockRecorder
}

// MockUserRepositoryInterfaceMockRecorder is the mock recorder for MockUserRepositoryInterface.
type MockUserRepositoryInterfaceMockRecorder struct {
	mock *MockUserRepositoryInterface
}

// NewMockUserRepositoryInterface creates a new mock instance.
func NewMockUserRepositoryInterface(ctrl *gomock.Controller) *MockUserRepositoryInterface {
	mock := &MockUserRepositoryInterface{ctrl: ctrl}
	mock.recorder = &MockUserRepositoryInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserRepositoryInterface) EXPECT() *MockUserRepositoryInterfaceMockRecorder {
	return m.recorder
}

// CreateUser mocks base method.
func (m *MockUserRepositoryInterface) CreateUser(arg0 context.Context, arg1 *dto.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockUserRepositoryInterfaceMockRecorder) CreateUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockUserRepositoryInterface)(nil).CreateUser), arg0, arg1)
}

// DeleteUser mocks base method.
func (m *MockUserRepositoryInterface) DeleteUser(arg0 context.Context, arg1 int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteUser", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteUser indicates an expected call of DeleteUser.
func (mr *MockUserRepositoryInterfaceMockRecorder) DeleteUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUser", reflect.TypeOf((*MockUserRepositoryInterface)(nil).DeleteUser), arg0, arg1)
}

// GetUserByID mocks base method.
func (m *MockUserRepositoryInterface) GetUserByID(arg0 context.Context, arg1 int) (*dto.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByID", arg0, arg1)
	ret0, _ := ret[0].(*dto.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByID indicates an expected call of GetUserByID.
func (mr *MockUserRepositoryInterfaceMockRecorder) GetUserByID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByID", reflect.TypeOf((*MockUserRepositoryInterface)(nil).GetUserByID), arg0, arg1)
}

// UpdatedUser mocks base method.
func (m *MockUserRepositoryInterface) UpdatedUser(arg0 context.Context, arg1 int, arg2 *dto.UserUpdate) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdatedUser", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdatedUser indicates an expected call of UpdatedUser.
func (mr *MockUserRepositoryInterfaceMockRecorder) UpdatedUser(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatedUser", reflect.TypeOf((*MockUserRepositoryInterface)(nil).UpdatedUser), arg0, arg1, arg2)
}