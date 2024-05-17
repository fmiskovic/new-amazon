// Code generated by mockery v2.42.2. DO NOT EDIT.

package core

import (
	sql "database/sql"

	mock "github.com/stretchr/testify/mock"
)

// DatabaseMock is an autogenerated mock type for the Database type
type DatabaseMock struct {
	mock.Mock
}

type DatabaseMock_Expecter struct {
	mock *mock.Mock
}

func (_m *DatabaseMock) EXPECT() *DatabaseMock_Expecter {
	return &DatabaseMock_Expecter{mock: &_m.Mock}
}

// Connect provides a mock function with given fields:
func (_m *DatabaseMock) Connect() (*sql.DB, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Connect")
	}

	var r0 *sql.DB
	var r1 error
	if rf, ok := ret.Get(0).(func() (*sql.DB, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() *sql.DB); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*sql.DB)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DatabaseMock_Connect_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Connect'
type DatabaseMock_Connect_Call struct {
	*mock.Call
}

// Connect is a helper method to define mock.On call
func (_e *DatabaseMock_Expecter) Connect() *DatabaseMock_Connect_Call {
	return &DatabaseMock_Connect_Call{Call: _e.mock.On("Connect")}
}

func (_c *DatabaseMock_Connect_Call) Run(run func()) *DatabaseMock_Connect_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *DatabaseMock_Connect_Call) Return(_a0 *sql.DB, _a1 error) *DatabaseMock_Connect_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *DatabaseMock_Connect_Call) RunAndReturn(run func() (*sql.DB, error)) *DatabaseMock_Connect_Call {
	_c.Call.Return(run)
	return _c
}

// NewDatabaseMock creates a new instance of DatabaseMock. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewDatabaseMock(t interface {
	mock.TestingT
	Cleanup(func())
}) *DatabaseMock {
	mock := &DatabaseMock{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
