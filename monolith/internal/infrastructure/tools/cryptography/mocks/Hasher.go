// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// Hasher is an autogenerated mocks type for the Hasher type
type Hasher struct {
	mock.Mock
}

// GenHash provides a mocks function with given fields: in, kind
func (_m *Hasher) GenHash(in []byte, kind int) []byte {
	ret := _m.Called(in, kind)

	var r0 []byte
	if rf, ok := ret.Get(0).(func([]byte, int) []byte); ok {
		r0 = rf(in, kind)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	return r0
}

// GenHashString provides a mocks function with given fields: in, kind
func (_m *Hasher) GenHashString(in []byte, kind int) string {
	ret := _m.Called(in, kind)

	var r0 string
	if rf, ok := ret.Get(0).(func([]byte, int) string); ok {
		r0 = rf(in, kind)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

type mockConstructorTestingTNewHasher interface {
	mock.TestingT
	Cleanup(func())
}

// NewHasher creates a new instance of Hasher. It also registers a testing interface on the mocks and a cleanup function to assert the mocks expectations.
func NewHasher(t mockConstructorTestingTNewHasher) *Hasher {
	mock := &Hasher{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
