// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/nerd500/axios-cp-wing/internal/database (interfaces: Querier)
//
// Generated by this command:
//
//	mockgen -package mockdb -destination ./internal/database/mock/mock.go github.com/nerd500/axios-cp-wing/internal/database Querier
//
// Package mockdb is a generated GoMock package.
package mockdb

import (
	context "context"
	reflect "reflect"

	database "github.com/nerd500/axios-cp-wing/internal/database"
	gomock "go.uber.org/mock/gomock"
)

// MockQuerier is a mock of Querier interface.
type MockQuerier struct {
	ctrl     *gomock.Controller
	recorder *MockQuerierMockRecorder
}

// MockQuerierMockRecorder is the mock recorder for MockQuerier.
type MockQuerierMockRecorder struct {
	mock *MockQuerier
}

// NewMockQuerier creates a new mock instance.
func NewMockQuerier(ctrl *gomock.Controller) *MockQuerier {
	mock := &MockQuerier{ctrl: ctrl}
	mock.recorder = &MockQuerierMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockQuerier) EXPECT() *MockQuerierMockRecorder {
	return m.recorder
}

// CreateTask mocks base method.
func (m *MockQuerier) CreateTask(arg0 context.Context, arg1 database.CreateTaskParams) (database.Task, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTask", arg0, arg1)
	ret0, _ := ret[0].(database.Task)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateTask indicates an expected call of CreateTask.
func (mr *MockQuerierMockRecorder) CreateTask(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTask", reflect.TypeOf((*MockQuerier)(nil).CreateTask), arg0, arg1)
}

// CreateUser mocks base method.
func (m *MockQuerier) CreateUser(arg0 context.Context, arg1 database.CreateUserParams) (database.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", arg0, arg1)
	ret0, _ := ret[0].(database.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockQuerierMockRecorder) CreateUser(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockQuerier)(nil).CreateUser), arg0, arg1)
}

// GetAllAdminUsers mocks base method.
func (m *MockQuerier) GetAllAdminUsers(arg0 context.Context) ([]database.GetAllAdminUsersRow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllAdminUsers", arg0)
	ret0, _ := ret[0].([]database.GetAllAdminUsersRow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllAdminUsers indicates an expected call of GetAllAdminUsers.
func (mr *MockQuerierMockRecorder) GetAllAdminUsers(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllAdminUsers", reflect.TypeOf((*MockQuerier)(nil).GetAllAdminUsers), arg0)
}

// GetAllTasks mocks base method.
func (m *MockQuerier) GetAllTasks(arg0 context.Context) ([]database.Task, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllTasks", arg0)
	ret0, _ := ret[0].([]database.Task)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllTasks indicates an expected call of GetAllTasks.
func (mr *MockQuerierMockRecorder) GetAllTasks(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllTasks", reflect.TypeOf((*MockQuerier)(nil).GetAllTasks), arg0)
}

// GetUser mocks base method.
func (m *MockQuerier) GetUser(arg0 context.Context, arg1 string) (database.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUser", arg0, arg1)
	ret0, _ := ret[0].(database.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUser indicates an expected call of GetUser.
func (mr *MockQuerierMockRecorder) GetUser(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUser", reflect.TypeOf((*MockQuerier)(nil).GetUser), arg0, arg1)
}

// GetUserFromAuthToken mocks base method.
func (m *MockQuerier) GetUserFromAuthToken(arg0 context.Context, arg1 string) (database.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserFromAuthToken", arg0, arg1)
	ret0, _ := ret[0].(database.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserFromAuthToken indicates an expected call of GetUserFromAuthToken.
func (mr *MockQuerierMockRecorder) GetUserFromAuthToken(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserFromAuthToken", reflect.TypeOf((*MockQuerier)(nil).GetUserFromAuthToken), arg0, arg1)
}
