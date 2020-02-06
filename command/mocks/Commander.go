// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	io "io"

	mock "github.com/stretchr/testify/mock"
)

// Commander is an autogenerated mock type for the Commander type
type Commander struct {
	mock.Mock
}

// Run provides a mock function with given fields: name, arg
func (_m *Commander) Run(name string, arg ...string) (string, error) {
	_va := make([]interface{}, len(arg))
	for _i := range arg {
		_va[_i] = arg[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, name)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 string
	if rf, ok := ret.Get(0).(func(string, ...string) string); ok {
		r0 = rf(name, arg...)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, ...string) error); ok {
		r1 = rf(name, arg...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Start provides a mock function with given fields: name, arg
func (_m *Commander) Start(name string, arg ...string) (io.ReadCloser, io.ReadCloser, func() error, error) {
	_va := make([]interface{}, len(arg))
	for _i := range arg {
		_va[_i] = arg[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, name)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 io.ReadCloser
	if rf, ok := ret.Get(0).(func(string, ...string) io.ReadCloser); ok {
		r0 = rf(name, arg...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(io.ReadCloser)
		}
	}

	var r1 io.ReadCloser
	if rf, ok := ret.Get(1).(func(string, ...string) io.ReadCloser); ok {
		r1 = rf(name, arg...)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(io.ReadCloser)
		}
	}

	var r2 func() error
	if rf, ok := ret.Get(2).(func(string, ...string) func() error); ok {
		r2 = rf(name, arg...)
	} else {
		if ret.Get(2) != nil {
			r2 = ret.Get(2).(func() error)
		}
	}

	var r3 error
	if rf, ok := ret.Get(3).(func(string, ...string) error); ok {
		r3 = rf(name, arg...)
	} else {
		r3 = ret.Error(3)
	}

	return r0, r1, r2, r3
}
