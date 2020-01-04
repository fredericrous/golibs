// Code generated by mockery v1.0.0. DO NOT EDIT.

package security

import mock "github.com/stretchr/testify/mock"

// MockRandom is an autogenerated mock type for the Random type
type MockRandom struct {
	mock.Mock
}

// GenerateRandomAlphaNum provides a mock function with given fields: length
func (_m *MockRandom) GenerateRandomAlphaNum(length uint64) string {
	ret := _m.Called(length)

	var r0 string
	if rf, ok := ret.Get(0).(func(uint64) string); ok {
		r0 = rf(length)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// GenerateRandomBytes provides a mock function with given fields: n
func (_m *MockRandom) GenerateRandomBytes(n int) ([]byte, error) {
	ret := _m.Called(n)

	var r0 []byte
	if rf, ok := ret.Get(0).(func(int) []byte); ok {
		r0 = rf(n)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(n)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GenerateRandomInt provides a mock function with given fields: n
func (_m *MockRandom) GenerateRandomInt(n int) int {
	ret := _m.Called(n)

	var r0 int
	if rf, ok := ret.Get(0).(func(int) int); ok {
		r0 = rf(n)
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

// GenerateRandomInt63 provides a mock function with given fields:
func (_m *MockRandom) GenerateRandomInt63() int64 {
	ret := _m.Called()

	var r0 int64
	if rf, ok := ret.Get(0).(func() int64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int64)
	}

	return r0
}

// GenerateRandomNum provides a mock function with given fields: n
func (_m *MockRandom) GenerateRandomNum(n uint64) string {
	ret := _m.Called(n)

	var r0 string
	if rf, ok := ret.Get(0).(func(uint64) string); ok {
		r0 = rf(n)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}
