// Code generated by mockery v1.0.0. DO NOT EDIT.

package mock

import mock "github.com/stretchr/testify/mock"

// Heights is an autogenerated mock type for the Heights type
type Heights struct {
	mock.Mock
}

// OnHeight provides a mock function with given fields: height, callback
func (_m *Heights) OnHeight(height uint64, callback func()) {
	_m.Called(height, callback)
}
