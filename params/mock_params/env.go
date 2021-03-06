// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/qdm12/golibs/params (interfaces: Env)

// Package mock_params is a generated GoMock package.
package mock_params

import (
	gomock "github.com/golang/mock/gomock"
	logging "github.com/qdm12/golibs/logging"
	params "github.com/qdm12/golibs/params"
	url "net/url"
	reflect "reflect"
	time "time"
)

// MockEnv is a mock of Env interface
type MockEnv struct {
	ctrl     *gomock.Controller
	recorder *MockEnvMockRecorder
}

// MockEnvMockRecorder is the mock recorder for MockEnv
type MockEnvMockRecorder struct {
	mock *MockEnv
}

// NewMockEnv creates a new mock instance
func NewMockEnv(ctrl *gomock.Controller) *MockEnv {
	mock := &MockEnv{ctrl: ctrl}
	mock.recorder = &MockEnvMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockEnv) EXPECT() *MockEnvMockRecorder {
	return m.recorder
}

// CSV mocks base method
func (m *MockEnv) CSV(arg0 string, arg1 ...params.OptionSetter) ([]string, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CSV", varargs...)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CSV indicates an expected call of CSV
func (mr *MockEnvMockRecorder) CSV(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CSV", reflect.TypeOf((*MockEnv)(nil).CSV), varargs...)
}

// CSVInside mocks base method
func (m *MockEnv) CSVInside(arg0 string, arg1 []string, arg2 ...params.OptionSetter) ([]string, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CSVInside", varargs...)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CSVInside indicates an expected call of CSVInside
func (mr *MockEnvMockRecorder) CSVInside(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CSVInside", reflect.TypeOf((*MockEnv)(nil).CSVInside), varargs...)
}

// Duration mocks base method
func (m *MockEnv) Duration(arg0 string, arg1 ...params.OptionSetter) (time.Duration, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Duration", varargs...)
	ret0, _ := ret[0].(time.Duration)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Duration indicates an expected call of Duration
func (mr *MockEnvMockRecorder) Duration(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Duration", reflect.TypeOf((*MockEnv)(nil).Duration), varargs...)
}

// Get mocks base method
func (m *MockEnv) Get(arg0 string, arg1 ...params.OptionSetter) (string, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Get", varargs...)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockEnvMockRecorder) Get(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockEnv)(nil).Get), varargs...)
}

// Inside mocks base method
func (m *MockEnv) Inside(arg0 string, arg1 []string, arg2 ...params.OptionSetter) (string, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Inside", varargs...)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Inside indicates an expected call of Inside
func (mr *MockEnvMockRecorder) Inside(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Inside", reflect.TypeOf((*MockEnv)(nil).Inside), varargs...)
}

// Int mocks base method
func (m *MockEnv) Int(arg0 string, arg1 ...params.OptionSetter) (int, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Int", varargs...)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Int indicates an expected call of Int
func (mr *MockEnvMockRecorder) Int(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Int", reflect.TypeOf((*MockEnv)(nil).Int), varargs...)
}

// IntRange mocks base method
func (m *MockEnv) IntRange(arg0 string, arg1, arg2 int, arg3 ...params.OptionSetter) (int, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2}
	for _, a := range arg3 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "IntRange", varargs...)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IntRange indicates an expected call of IntRange
func (mr *MockEnvMockRecorder) IntRange(arg0, arg1, arg2 interface{}, arg3 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2}, arg3...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IntRange", reflect.TypeOf((*MockEnv)(nil).IntRange), varargs...)
}

// ListeningAddress mocks base method
func (m *MockEnv) ListeningAddress(arg0 string, arg1 ...params.OptionSetter) (string, string, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListeningAddress", varargs...)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListeningAddress indicates an expected call of ListeningAddress
func (mr *MockEnvMockRecorder) ListeningAddress(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListeningAddress", reflect.TypeOf((*MockEnv)(nil).ListeningAddress), varargs...)
}

// ListeningPort mocks base method
func (m *MockEnv) ListeningPort(arg0 string, arg1 ...params.OptionSetter) (uint16, string, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListeningPort", varargs...)
	ret0, _ := ret[0].(uint16)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListeningPort indicates an expected call of ListeningPort
func (mr *MockEnvMockRecorder) ListeningPort(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListeningPort", reflect.TypeOf((*MockEnv)(nil).ListeningPort), varargs...)
}

// LogEncoding mocks base method
func (m *MockEnv) LogEncoding(arg0 string, arg1 ...params.OptionSetter) (logging.Encoding, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "LogEncoding", varargs...)
	ret0, _ := ret[0].(logging.Encoding)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LogEncoding indicates an expected call of LogEncoding
func (mr *MockEnvMockRecorder) LogEncoding(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LogEncoding", reflect.TypeOf((*MockEnv)(nil).LogEncoding), varargs...)
}

// LogLevel mocks base method
func (m *MockEnv) LogLevel(arg0 string, arg1 ...params.OptionSetter) (logging.Level, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "LogLevel", varargs...)
	ret0, _ := ret[0].(logging.Level)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LogLevel indicates an expected call of LogLevel
func (mr *MockEnvMockRecorder) LogLevel(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LogLevel", reflect.TypeOf((*MockEnv)(nil).LogLevel), varargs...)
}

// OnOff mocks base method
func (m *MockEnv) OnOff(arg0 string, arg1 ...params.OptionSetter) (bool, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "OnOff", varargs...)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// OnOff indicates an expected call of OnOff
func (mr *MockEnvMockRecorder) OnOff(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OnOff", reflect.TypeOf((*MockEnv)(nil).OnOff), varargs...)
}

// Path mocks base method
func (m *MockEnv) Path(arg0 string, arg1 ...params.OptionSetter) (string, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Path", varargs...)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Path indicates an expected call of Path
func (mr *MockEnvMockRecorder) Path(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Path", reflect.TypeOf((*MockEnv)(nil).Path), varargs...)
}

// Port mocks base method
func (m *MockEnv) Port(arg0 string, arg1 ...params.OptionSetter) (uint16, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Port", varargs...)
	ret0, _ := ret[0].(uint16)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Port indicates an expected call of Port
func (mr *MockEnvMockRecorder) Port(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Port", reflect.TypeOf((*MockEnv)(nil).Port), varargs...)
}

// RootURL mocks base method
func (m *MockEnv) RootURL(arg0 string, arg1 ...params.OptionSetter) (string, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RootURL", varargs...)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RootURL indicates an expected call of RootURL
func (mr *MockEnvMockRecorder) RootURL(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RootURL", reflect.TypeOf((*MockEnv)(nil).RootURL), varargs...)
}

// URL mocks base method
func (m *MockEnv) URL(arg0 string, arg1 ...params.OptionSetter) (*url.URL, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "URL", varargs...)
	ret0, _ := ret[0].(*url.URL)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// URL indicates an expected call of URL
func (mr *MockEnvMockRecorder) URL(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "URL", reflect.TypeOf((*MockEnv)(nil).URL), varargs...)
}

// YesNo mocks base method
func (m *MockEnv) YesNo(arg0 string, arg1 ...params.OptionSetter) (bool, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "YesNo", varargs...)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// YesNo indicates an expected call of YesNo
func (mr *MockEnvMockRecorder) YesNo(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "YesNo", reflect.TypeOf((*MockEnv)(nil).YesNo), varargs...)
}
