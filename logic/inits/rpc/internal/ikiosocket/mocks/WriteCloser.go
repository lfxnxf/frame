// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import context "golang.org/x/net/context"
import ikio "github.com/lfxnxf/frame/tpc/inf/ikio"
import mock "github.com/stretchr/testify/mock"

// WriteCloser is an autogenerated mock type for the WriteCloser type
type WriteCloser struct {
	mock.Mock
}

// Close provides a mock function with given fields:
func (_m *WriteCloser) Close() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Write provides a mock function with given fields: _a0, _a1
func (_m *WriteCloser) Write(_a0 context.Context, _a1 ikio.Packet) (int, error) {
	ret := _m.Called(_a0, _a1)

	var r0 int
	if rf, ok := ret.Get(0).(func(context.Context, ikio.Packet) int); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, ikio.Packet) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
