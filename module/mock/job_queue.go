// Code generated by mockery v1.0.0. DO NOT EDIT.

package mock

import (
	module "github.com/onflow/flow-go/module"
	mock "github.com/stretchr/testify/mock"
)

// JobQueue is an autogenerated mock type for the JobQueue type
type JobQueue struct {
	mock.Mock
}

// Add provides a mock function with given fields: job
func (_m *JobQueue) Add(job module.Job) error {
	ret := _m.Called(job)

	var r0 error
	if rf, ok := ret.Get(0).(func(module.Job) error); ok {
		r0 = rf(job)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
