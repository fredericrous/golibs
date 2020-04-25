// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/qdm12/golibs/files (interfaces: FileManager)

// Package mock_files is a generated GoMock package.
package mock_files

import (
	gomock "github.com/golang/mock/gomock"
	files "github.com/qdm12/golibs/files"
	os "os"
	reflect "reflect"
)

// MockFileManager is a mock of FileManager interface
type MockFileManager struct {
	ctrl     *gomock.Controller
	recorder *MockFileManagerMockRecorder
}

// MockFileManagerMockRecorder is the mock recorder for MockFileManager
type MockFileManagerMockRecorder struct {
	mock *MockFileManager
}

// NewMockFileManager creates a new mock instance
func NewMockFileManager(ctrl *gomock.Controller) *MockFileManager {
	mock := &MockFileManager{ctrl: ctrl}
	mock.recorder = &MockFileManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockFileManager) EXPECT() *MockFileManagerMockRecorder {
	return m.recorder
}

// CreateDir mocks base method
func (m *MockFileManager) CreateDir(arg0 string, arg1 ...files.WriteOptionSetter) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateDir", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateDir indicates an expected call of CreateDir
func (mr *MockFileManagerMockRecorder) CreateDir(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateDir", reflect.TypeOf((*MockFileManager)(nil).CreateDir), varargs...)
}

// DirectoryExists mocks base method
func (m *MockFileManager) DirectoryExists(arg0 string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DirectoryExists", arg0)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DirectoryExists indicates an expected call of DirectoryExists
func (mr *MockFileManagerMockRecorder) DirectoryExists(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DirectoryExists", reflect.TypeOf((*MockFileManager)(nil).DirectoryExists), arg0)
}

// FileExists mocks base method
func (m *MockFileManager) FileExists(arg0 string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FileExists", arg0)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FileExists indicates an expected call of FileExists
func (mr *MockFileManagerMockRecorder) FileExists(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FileExists", reflect.TypeOf((*MockFileManager)(nil).FileExists), arg0)
}

// FilepathExists mocks base method
func (m *MockFileManager) FilepathExists(arg0 string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FilepathExists", arg0)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FilepathExists indicates an expected call of FilepathExists
func (mr *MockFileManagerMockRecorder) FilepathExists(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FilepathExists", reflect.TypeOf((*MockFileManager)(nil).FilepathExists), arg0)
}

// GetGroupPermissions mocks base method
func (m *MockFileManager) GetGroupPermissions(arg0 string) (bool, bool, bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetGroupPermissions", arg0)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(bool)
	ret2, _ := ret[2].(bool)
	ret3, _ := ret[3].(error)
	return ret0, ret1, ret2, ret3
}

// GetGroupPermissions indicates an expected call of GetGroupPermissions
func (mr *MockFileManagerMockRecorder) GetGroupPermissions(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGroupPermissions", reflect.TypeOf((*MockFileManager)(nil).GetGroupPermissions), arg0)
}

// GetOthersPermissions mocks base method
func (m *MockFileManager) GetOthersPermissions(arg0 string) (bool, bool, bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOthersPermissions", arg0)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(bool)
	ret2, _ := ret[2].(bool)
	ret3, _ := ret[3].(error)
	return ret0, ret1, ret2, ret3
}

// GetOthersPermissions indicates an expected call of GetOthersPermissions
func (mr *MockFileManagerMockRecorder) GetOthersPermissions(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOthersPermissions", reflect.TypeOf((*MockFileManager)(nil).GetOthersPermissions), arg0)
}

// GetOwnership mocks base method
func (m *MockFileManager) GetOwnership(arg0 string) (int, int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOwnership", arg0)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(int)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetOwnership indicates an expected call of GetOwnership
func (mr *MockFileManagerMockRecorder) GetOwnership(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOwnership", reflect.TypeOf((*MockFileManager)(nil).GetOwnership), arg0)
}

// GetUserPermissions mocks base method
func (m *MockFileManager) GetUserPermissions(arg0 string) (bool, bool, bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserPermissions", arg0)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(bool)
	ret2, _ := ret[2].(bool)
	ret3, _ := ret[3].(error)
	return ret0, ret1, ret2, ret3
}

// GetUserPermissions indicates an expected call of GetUserPermissions
func (mr *MockFileManagerMockRecorder) GetUserPermissions(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserPermissions", reflect.TypeOf((*MockFileManager)(nil).GetUserPermissions), arg0)
}

// ReadFile mocks base method
func (m *MockFileManager) ReadFile(arg0 string) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadFile", arg0)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadFile indicates an expected call of ReadFile
func (mr *MockFileManagerMockRecorder) ReadFile(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadFile", reflect.TypeOf((*MockFileManager)(nil).ReadFile), arg0)
}

// Remove mocks base method
func (m *MockFileManager) Remove(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Remove", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Remove indicates an expected call of Remove
func (mr *MockFileManagerMockRecorder) Remove(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Remove", reflect.TypeOf((*MockFileManager)(nil).Remove), arg0)
}

// SetOwnership mocks base method
func (m *MockFileManager) SetOwnership(arg0 string, arg1, arg2 int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetOwnership", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetOwnership indicates an expected call of SetOwnership
func (mr *MockFileManagerMockRecorder) SetOwnership(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetOwnership", reflect.TypeOf((*MockFileManager)(nil).SetOwnership), arg0, arg1, arg2)
}

// SetUserPermissions mocks base method
func (m *MockFileManager) SetUserPermissions(arg0 string, arg1 os.FileMode) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetUserPermissions", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetUserPermissions indicates an expected call of SetUserPermissions
func (mr *MockFileManagerMockRecorder) SetUserPermissions(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetUserPermissions", reflect.TypeOf((*MockFileManager)(nil).SetUserPermissions), arg0, arg1)
}

// Touch mocks base method
func (m *MockFileManager) Touch(arg0 string, arg1 ...files.WriteOptionSetter) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Touch", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Touch indicates an expected call of Touch
func (mr *MockFileManagerMockRecorder) Touch(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Touch", reflect.TypeOf((*MockFileManager)(nil).Touch), varargs...)
}

// WriteLinesToFile mocks base method
func (m *MockFileManager) WriteLinesToFile(arg0 string, arg1 []string, arg2 ...files.WriteOptionSetter) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "WriteLinesToFile", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// WriteLinesToFile indicates an expected call of WriteLinesToFile
func (mr *MockFileManagerMockRecorder) WriteLinesToFile(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WriteLinesToFile", reflect.TypeOf((*MockFileManager)(nil).WriteLinesToFile), varargs...)
}

// WriteToFile mocks base method
func (m *MockFileManager) WriteToFile(arg0 string, arg1 []byte, arg2 ...files.WriteOptionSetter) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "WriteToFile", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// WriteToFile indicates an expected call of WriteToFile
func (mr *MockFileManagerMockRecorder) WriteToFile(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WriteToFile", reflect.TypeOf((*MockFileManager)(nil).WriteToFile), varargs...)
}
