// Code generated by mockery v2.52.1. DO NOT EDIT.

package mocks

import (
	io "io"

	gcs "github.com/google/syzkaller/pkg/gcs"

	mock "github.com/stretchr/testify/mock"
)

// Client is an autogenerated mock type for the Client type
type Client struct {
	mock.Mock
}

type Client_Expecter struct {
	mock *mock.Mock
}

func (_m *Client) EXPECT() *Client_Expecter {
	return &Client_Expecter{mock: &_m.Mock}
}

// Close provides a mock function with no fields
func (_m *Client) Close() error {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Close")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Client_Close_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Close'
type Client_Close_Call struct {
	*mock.Call
}

// Close is a helper method to define mock.On call
func (_e *Client_Expecter) Close() *Client_Close_Call {
	return &Client_Close_Call{Call: _e.mock.On("Close")}
}

func (_c *Client_Close_Call) Run(run func()) *Client_Close_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Client_Close_Call) Return(_a0 error) *Client_Close_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Client_Close_Call) RunAndReturn(run func() error) *Client_Close_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteFile provides a mock function with given fields: path
func (_m *Client) DeleteFile(path string) error {
	ret := _m.Called(path)

	if len(ret) == 0 {
		panic("no return value specified for DeleteFile")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(path)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Client_DeleteFile_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteFile'
type Client_DeleteFile_Call struct {
	*mock.Call
}

// DeleteFile is a helper method to define mock.On call
//   - path string
func (_e *Client_Expecter) DeleteFile(path interface{}) *Client_DeleteFile_Call {
	return &Client_DeleteFile_Call{Call: _e.mock.On("DeleteFile", path)}
}

func (_c *Client_DeleteFile_Call) Run(run func(path string)) *Client_DeleteFile_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *Client_DeleteFile_Call) Return(_a0 error) *Client_DeleteFile_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Client_DeleteFile_Call) RunAndReturn(run func(string) error) *Client_DeleteFile_Call {
	_c.Call.Return(run)
	return _c
}

// FileExists provides a mock function with given fields: path
func (_m *Client) FileExists(path string) (bool, error) {
	ret := _m.Called(path)

	if len(ret) == 0 {
		panic("no return value specified for FileExists")
	}

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (bool, error)); ok {
		return rf(path)
	}
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(path)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(path)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Client_FileExists_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FileExists'
type Client_FileExists_Call struct {
	*mock.Call
}

// FileExists is a helper method to define mock.On call
//   - path string
func (_e *Client_Expecter) FileExists(path interface{}) *Client_FileExists_Call {
	return &Client_FileExists_Call{Call: _e.mock.On("FileExists", path)}
}

func (_c *Client_FileExists_Call) Run(run func(path string)) *Client_FileExists_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *Client_FileExists_Call) Return(_a0 bool, _a1 error) *Client_FileExists_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Client_FileExists_Call) RunAndReturn(run func(string) (bool, error)) *Client_FileExists_Call {
	_c.Call.Return(run)
	return _c
}

// FileReader provides a mock function with given fields: path
func (_m *Client) FileReader(path string) (io.ReadCloser, error) {
	ret := _m.Called(path)

	if len(ret) == 0 {
		panic("no return value specified for FileReader")
	}

	var r0 io.ReadCloser
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (io.ReadCloser, error)); ok {
		return rf(path)
	}
	if rf, ok := ret.Get(0).(func(string) io.ReadCloser); ok {
		r0 = rf(path)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(io.ReadCloser)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(path)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Client_FileReader_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FileReader'
type Client_FileReader_Call struct {
	*mock.Call
}

// FileReader is a helper method to define mock.On call
//   - path string
func (_e *Client_Expecter) FileReader(path interface{}) *Client_FileReader_Call {
	return &Client_FileReader_Call{Call: _e.mock.On("FileReader", path)}
}

func (_c *Client_FileReader_Call) Run(run func(path string)) *Client_FileReader_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *Client_FileReader_Call) Return(_a0 io.ReadCloser, _a1 error) *Client_FileReader_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Client_FileReader_Call) RunAndReturn(run func(string) (io.ReadCloser, error)) *Client_FileReader_Call {
	_c.Call.Return(run)
	return _c
}

// FileWriter provides a mock function with given fields: path, contentType, contentEncoding
func (_m *Client) FileWriter(path string, contentType string, contentEncoding string) (io.WriteCloser, error) {
	ret := _m.Called(path, contentType, contentEncoding)

	if len(ret) == 0 {
		panic("no return value specified for FileWriter")
	}

	var r0 io.WriteCloser
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string, string) (io.WriteCloser, error)); ok {
		return rf(path, contentType, contentEncoding)
	}
	if rf, ok := ret.Get(0).(func(string, string, string) io.WriteCloser); ok {
		r0 = rf(path, contentType, contentEncoding)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(io.WriteCloser)
		}
	}

	if rf, ok := ret.Get(1).(func(string, string, string) error); ok {
		r1 = rf(path, contentType, contentEncoding)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Client_FileWriter_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FileWriter'
type Client_FileWriter_Call struct {
	*mock.Call
}

// FileWriter is a helper method to define mock.On call
//   - path string
//   - contentType string
//   - contentEncoding string
func (_e *Client_Expecter) FileWriter(path interface{}, contentType interface{}, contentEncoding interface{}) *Client_FileWriter_Call {
	return &Client_FileWriter_Call{Call: _e.mock.On("FileWriter", path, contentType, contentEncoding)}
}

func (_c *Client_FileWriter_Call) Run(run func(path string, contentType string, contentEncoding string)) *Client_FileWriter_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(string), args[2].(string))
	})
	return _c
}

func (_c *Client_FileWriter_Call) Return(_a0 io.WriteCloser, _a1 error) *Client_FileWriter_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Client_FileWriter_Call) RunAndReturn(run func(string, string, string) (io.WriteCloser, error)) *Client_FileWriter_Call {
	_c.Call.Return(run)
	return _c
}

// ListObjects provides a mock function with given fields: path
func (_m *Client) ListObjects(path string) ([]*gcs.Object, error) {
	ret := _m.Called(path)

	if len(ret) == 0 {
		panic("no return value specified for ListObjects")
	}

	var r0 []*gcs.Object
	var r1 error
	if rf, ok := ret.Get(0).(func(string) ([]*gcs.Object, error)); ok {
		return rf(path)
	}
	if rf, ok := ret.Get(0).(func(string) []*gcs.Object); ok {
		r0 = rf(path)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*gcs.Object)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(path)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Client_ListObjects_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListObjects'
type Client_ListObjects_Call struct {
	*mock.Call
}

// ListObjects is a helper method to define mock.On call
//   - path string
func (_e *Client_Expecter) ListObjects(path interface{}) *Client_ListObjects_Call {
	return &Client_ListObjects_Call{Call: _e.mock.On("ListObjects", path)}
}

func (_c *Client_ListObjects_Call) Run(run func(path string)) *Client_ListObjects_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *Client_ListObjects_Call) Return(_a0 []*gcs.Object, _a1 error) *Client_ListObjects_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Client_ListObjects_Call) RunAndReturn(run func(string) ([]*gcs.Object, error)) *Client_ListObjects_Call {
	_c.Call.Return(run)
	return _c
}

// Publish provides a mock function with given fields: path
func (_m *Client) Publish(path string) error {
	ret := _m.Called(path)

	if len(ret) == 0 {
		panic("no return value specified for Publish")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(path)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Client_Publish_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Publish'
type Client_Publish_Call struct {
	*mock.Call
}

// Publish is a helper method to define mock.On call
//   - path string
func (_e *Client_Expecter) Publish(path interface{}) *Client_Publish_Call {
	return &Client_Publish_Call{Call: _e.mock.On("Publish", path)}
}

func (_c *Client_Publish_Call) Run(run func(path string)) *Client_Publish_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *Client_Publish_Call) Return(_a0 error) *Client_Publish_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Client_Publish_Call) RunAndReturn(run func(string) error) *Client_Publish_Call {
	_c.Call.Return(run)
	return _c
}

// NewClient creates a new instance of Client. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewClient(t interface {
	mock.TestingT
	Cleanup(func())
}) *Client {
	mock := &Client{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
