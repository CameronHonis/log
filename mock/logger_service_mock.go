// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/CameronHonis/log (interfaces: LoggerServiceI)
//
// Generated by this command:
//
//	mockgen -destination mock/logger_service_mock.go . LoggerServiceI
//

// Package mock_log is a generated GoMock package.
package mock_log

import (
	reflect "reflect"

	service "github.com/CameronHonis/service"
	gomock "go.uber.org/mock/gomock"
)

// MockLoggerServiceI is a mock of LoggerServiceI interface.
type MockLoggerServiceI struct {
	ctrl     *gomock.Controller
	recorder *MockLoggerServiceIMockRecorder
}

// MockLoggerServiceIMockRecorder is the mock recorder for MockLoggerServiceI.
type MockLoggerServiceIMockRecorder struct {
	mock *MockLoggerServiceI
}

// NewMockLoggerServiceI creates a new mock instance.
func NewMockLoggerServiceI(ctrl *gomock.Controller) *MockLoggerServiceI {
	mock := &MockLoggerServiceI{ctrl: ctrl}
	mock.recorder = &MockLoggerServiceIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLoggerServiceI) EXPECT() *MockLoggerServiceIMockRecorder {
	return m.recorder
}

// AddDependency mocks base method.
func (m *MockLoggerServiceI) AddDependency(arg0 service.ServiceI) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "AddDependency", arg0)
}

// AddDependency indicates an expected call of AddDependency.
func (mr *MockLoggerServiceIMockRecorder) AddDependency(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddDependency", reflect.TypeOf((*MockLoggerServiceI)(nil).AddDependency), arg0)
}

// AddEventListener mocks base method.
func (m *MockLoggerServiceI) AddEventListener(arg0 service.EventVariant, arg1 service.EventHandler) int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddEventListener", arg0, arg1)
	ret0, _ := ret[0].(int)
	return ret0
}

// AddEventListener indicates an expected call of AddEventListener.
func (mr *MockLoggerServiceIMockRecorder) AddEventListener(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddEventListener", reflect.TypeOf((*MockLoggerServiceI)(nil).AddEventListener), arg0, arg1)
}

// Config mocks base method.
func (m *MockLoggerServiceI) Config() service.ConfigI {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Config")
	ret0, _ := ret[0].(service.ConfigI)
	return ret0
}

// Config indicates an expected call of Config.
func (mr *MockLoggerServiceIMockRecorder) Config() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Config", reflect.TypeOf((*MockLoggerServiceI)(nil).Config))
}

// Dependencies mocks base method.
func (m *MockLoggerServiceI) Dependencies() []service.ServiceI {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Dependencies")
	ret0, _ := ret[0].([]service.ServiceI)
	return ret0
}

// Dependencies indicates an expected call of Dependencies.
func (mr *MockLoggerServiceIMockRecorder) Dependencies() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Dependencies", reflect.TypeOf((*MockLoggerServiceI)(nil).Dependencies))
}

// Dispatch mocks base method.
func (m *MockLoggerServiceI) Dispatch(arg0 service.EventI) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Dispatch", arg0)
}

// Dispatch indicates an expected call of Dispatch.
func (mr *MockLoggerServiceIMockRecorder) Dispatch(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Dispatch", reflect.TypeOf((*MockLoggerServiceI)(nil).Dispatch), arg0)
}

// Log mocks base method.
func (m *MockLoggerServiceI) Log(arg0 string, arg1 ...any) {
	m.ctrl.T.Helper()
	varargs := []any{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Log", varargs...)
}

// Log indicates an expected call of Log.
func (mr *MockLoggerServiceIMockRecorder) Log(arg0 any, arg1 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Log", reflect.TypeOf((*MockLoggerServiceI)(nil).Log), varargs...)
}

// LogBlue mocks base method.
func (m *MockLoggerServiceI) LogBlue(arg0 string, arg1 ...any) {
	m.ctrl.T.Helper()
	varargs := []any{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "LogBlue", varargs...)
}

// LogBlue indicates an expected call of LogBlue.
func (mr *MockLoggerServiceIMockRecorder) LogBlue(arg0 any, arg1 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LogBlue", reflect.TypeOf((*MockLoggerServiceI)(nil).LogBlue), varargs...)
}

// LogBrown mocks base method.
func (m *MockLoggerServiceI) LogBrown(arg0 string, arg1 ...any) {
	m.ctrl.T.Helper()
	varargs := []any{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "LogBrown", varargs...)
}

// LogBrown indicates an expected call of LogBrown.
func (mr *MockLoggerServiceIMockRecorder) LogBrown(arg0 any, arg1 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LogBrown", reflect.TypeOf((*MockLoggerServiceI)(nil).LogBrown), varargs...)
}

// LogCyan mocks base method.
func (m *MockLoggerServiceI) LogCyan(arg0 string, arg1 ...any) {
	m.ctrl.T.Helper()
	varargs := []any{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "LogCyan", varargs...)
}

// LogCyan indicates an expected call of LogCyan.
func (mr *MockLoggerServiceIMockRecorder) LogCyan(arg0 any, arg1 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LogCyan", reflect.TypeOf((*MockLoggerServiceI)(nil).LogCyan), varargs...)
}

// LogGreen mocks base method.
func (m *MockLoggerServiceI) LogGreen(arg0 string, arg1 ...any) {
	m.ctrl.T.Helper()
	varargs := []any{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "LogGreen", varargs...)
}

// LogGreen indicates an expected call of LogGreen.
func (mr *MockLoggerServiceIMockRecorder) LogGreen(arg0 any, arg1 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LogGreen", reflect.TypeOf((*MockLoggerServiceI)(nil).LogGreen), varargs...)
}

// LogMagenta mocks base method.
func (m *MockLoggerServiceI) LogMagenta(arg0 string, arg1 ...any) {
	m.ctrl.T.Helper()
	varargs := []any{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "LogMagenta", varargs...)
}

// LogMagenta indicates an expected call of LogMagenta.
func (mr *MockLoggerServiceIMockRecorder) LogMagenta(arg0 any, arg1 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LogMagenta", reflect.TypeOf((*MockLoggerServiceI)(nil).LogMagenta), varargs...)
}

// LogOrange mocks base method.
func (m *MockLoggerServiceI) LogOrange(arg0 string, arg1 ...any) {
	m.ctrl.T.Helper()
	varargs := []any{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "LogOrange", varargs...)
}

// LogOrange indicates an expected call of LogOrange.
func (mr *MockLoggerServiceIMockRecorder) LogOrange(arg0 any, arg1 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LogOrange", reflect.TypeOf((*MockLoggerServiceI)(nil).LogOrange), varargs...)
}

// LogRed mocks base method.
func (m *MockLoggerServiceI) LogRed(arg0 string, arg1 ...any) {
	m.ctrl.T.Helper()
	varargs := []any{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "LogRed", varargs...)
}

// LogRed indicates an expected call of LogRed.
func (mr *MockLoggerServiceIMockRecorder) LogRed(arg0 any, arg1 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LogRed", reflect.TypeOf((*MockLoggerServiceI)(nil).LogRed), varargs...)
}

// LogYellow mocks base method.
func (m *MockLoggerServiceI) LogYellow(arg0 string, arg1 ...any) {
	m.ctrl.T.Helper()
	varargs := []any{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "LogYellow", varargs...)
}

// LogYellow indicates an expected call of LogYellow.
func (mr *MockLoggerServiceIMockRecorder) LogYellow(arg0 any, arg1 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LogYellow", reflect.TypeOf((*MockLoggerServiceI)(nil).LogYellow), varargs...)
}

// OnStart mocks base method.
func (m *MockLoggerServiceI) OnStart() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "OnStart")
}

// OnStart indicates an expected call of OnStart.
func (mr *MockLoggerServiceIMockRecorder) OnStart() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OnStart", reflect.TypeOf((*MockLoggerServiceI)(nil).OnStart))
}

// RemoveEventListener mocks base method.
func (m *MockLoggerServiceI) RemoveEventListener(arg0 int) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RemoveEventListener", arg0)
}

// RemoveEventListener indicates an expected call of RemoveEventListener.
func (mr *MockLoggerServiceIMockRecorder) RemoveEventListener(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveEventListener", reflect.TypeOf((*MockLoggerServiceI)(nil).RemoveEventListener), arg0)
}

// SetParent mocks base method.
func (m *MockLoggerServiceI) SetParent(arg0 service.ServiceI) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetParent", arg0)
}

// SetParent indicates an expected call of SetParent.
func (mr *MockLoggerServiceIMockRecorder) SetParent(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetParent", reflect.TypeOf((*MockLoggerServiceI)(nil).SetParent), arg0)
}

// Start mocks base method.
func (m *MockLoggerServiceI) Start() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Start")
}

// Start indicates an expected call of Start.
func (mr *MockLoggerServiceIMockRecorder) Start() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Start", reflect.TypeOf((*MockLoggerServiceI)(nil).Start))
}
