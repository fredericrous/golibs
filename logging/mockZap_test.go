// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/qdm12/golibs/logging (interfaces: Zap)

// Package logging is a generated GoMock package.
package logging

import (
	gomock "github.com/golang/mock/gomock"
	zapcore "go.uber.org/zap/zapcore"
	reflect "reflect"
)

// MockZap is a mock of Zap interface
type MockZap struct {
	ctrl     *gomock.Controller
	recorder *MockZapMockRecorder
}

// MockZapMockRecorder is the mock recorder for MockZap
type MockZapMockRecorder struct {
	mock *MockZap
}

// NewMockZap creates a new mock instance
func NewMockZap(ctrl *gomock.Controller) *MockZap {
	mock := &MockZap{ctrl: ctrl}
	mock.recorder = &MockZapMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockZap) EXPECT() *MockZapMockRecorder {
	return m.recorder
}

// Debug mocks base method
func (m *MockZap) Debug(arg0 string, arg1 ...zapcore.Field) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Debug", varargs...)
}

// Debug indicates an expected call of Debug
func (mr *MockZapMockRecorder) Debug(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Debug", reflect.TypeOf((*MockZap)(nil).Debug), varargs...)
}

// Error mocks base method
func (m *MockZap) Error(arg0 string, arg1 ...zapcore.Field) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Error", varargs...)
}

// Error indicates an expected call of Error
func (mr *MockZapMockRecorder) Error(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Error", reflect.TypeOf((*MockZap)(nil).Error), varargs...)
}

// Info mocks base method
func (m *MockZap) Info(arg0 string, arg1 ...zapcore.Field) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Info", varargs...)
}

// Info indicates an expected call of Info
func (mr *MockZapMockRecorder) Info(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Info", reflect.TypeOf((*MockZap)(nil).Info), varargs...)
}

// Sync mocks base method
func (m *MockZap) Sync() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Sync")
	ret0, _ := ret[0].(error)
	return ret0
}

// Sync indicates an expected call of Sync
func (mr *MockZapMockRecorder) Sync() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Sync", reflect.TypeOf((*MockZap)(nil).Sync))
}

// Warn mocks base method
func (m *MockZap) Warn(arg0 string, arg1 ...zapcore.Field) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Warn", varargs...)
}

// Warn indicates an expected call of Warn
func (mr *MockZapMockRecorder) Warn(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Warn", reflect.TypeOf((*MockZap)(nil).Warn), varargs...)
}