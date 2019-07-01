// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import container "github.com/mesg-foundation/engine/container"
import io "io"
import mock "github.com/stretchr/testify/mock"
import swarm "github.com/docker/docker/api/types/swarm"
import types "github.com/docker/docker/api/types"

// Container is an autogenerated mock type for the Container type
type Container struct {
	mock.Mock
}

// Build provides a mock function with given fields: path
func (_m *Container) Build(path string) (string, error) {
	ret := _m.Called(path)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(path)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(path)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateNetwork provides a mock function with given fields: namespace
func (_m *Container) CreateNetwork(namespace []string) (string, error) {
	ret := _m.Called(namespace)

	var r0 string
	if rf, ok := ret.Get(0).(func([]string) string); ok {
		r0 = rf(namespace)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]string) error); ok {
		r1 = rf(namespace)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteNetwork provides a mock function with given fields: namespace
func (_m *Container) DeleteNetwork(namespace []string) error {
	ret := _m.Called(namespace)

	var r0 error
	if rf, ok := ret.Get(0).(func([]string) error); ok {
		r0 = rf(namespace)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteVolume provides a mock function with given fields: name
func (_m *Container) DeleteVolume(name string) error {
	ret := _m.Called(name)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(name)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FindContainer provides a mock function with given fields: namespace
func (_m *Container) FindContainer(namespace []string) (types.ContainerJSON, error) {
	ret := _m.Called(namespace)

	var r0 types.ContainerJSON
	if rf, ok := ret.Get(0).(func([]string) types.ContainerJSON); ok {
		r0 = rf(namespace)
	} else {
		r0 = ret.Get(0).(types.ContainerJSON)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]string) error); ok {
		r1 = rf(namespace)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindNetwork provides a mock function with given fields: namespace
func (_m *Container) FindNetwork(namespace []string) (types.NetworkResource, error) {
	ret := _m.Called(namespace)

	var r0 types.NetworkResource
	if rf, ok := ret.Get(0).(func([]string) types.NetworkResource); ok {
		r0 = rf(namespace)
	} else {
		r0 = ret.Get(0).(types.NetworkResource)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]string) error); ok {
		r1 = rf(namespace)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindService provides a mock function with given fields: namespace
func (_m *Container) FindService(namespace []string) (swarm.Service, error) {
	ret := _m.Called(namespace)

	var r0 swarm.Service
	if rf, ok := ret.Get(0).(func([]string) swarm.Service); ok {
		r0 = rf(namespace)
	} else {
		r0 = ret.Get(0).(swarm.Service)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]string) error); ok {
		r1 = rf(namespace)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListServices provides a mock function with given fields: labels
func (_m *Container) ListServices(labels ...string) ([]swarm.Service, error) {
	_va := make([]interface{}, len(labels))
	for _i := range labels {
		_va[_i] = labels[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 []swarm.Service
	if rf, ok := ret.Get(0).(func(...string) []swarm.Service); ok {
		r0 = rf(labels...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]swarm.Service)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(...string) error); ok {
		r1 = rf(labels...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListTasks provides a mock function with given fields: namespace
func (_m *Container) ListTasks(namespace []string) ([]swarm.Task, error) {
	ret := _m.Called(namespace)

	var r0 []swarm.Task
	if rf, ok := ret.Get(0).(func([]string) []swarm.Task); ok {
		r0 = rf(namespace)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]swarm.Task)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]string) error); ok {
		r1 = rf(namespace)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Namespace provides a mock function with given fields: ss
func (_m *Container) Namespace(ss []string) string {
	ret := _m.Called(ss)

	var r0 string
	if rf, ok := ret.Get(0).(func([]string) string); ok {
		r0 = rf(ss)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// ServiceLogs provides a mock function with given fields: namespace
func (_m *Container) ServiceLogs(namespace []string) (io.ReadCloser, error) {
	ret := _m.Called(namespace)

	var r0 io.ReadCloser
	if rf, ok := ret.Get(0).(func([]string) io.ReadCloser); ok {
		r0 = rf(namespace)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(io.ReadCloser)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]string) error); ok {
		r1 = rf(namespace)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SharedNetworkID provides a mock function with given fields:
func (_m *Container) SharedNetworkID() (string, error) {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// StartService provides a mock function with given fields: options
func (_m *Container) StartService(options container.ServiceOptions) (string, error) {
	ret := _m.Called(options)

	var r0 string
	if rf, ok := ret.Get(0).(func(container.ServiceOptions) string); ok {
		r0 = rf(options)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(container.ServiceOptions) error); ok {
		r1 = rf(options)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Status provides a mock function with given fields: namespace
func (_m *Container) Status(namespace []string) (container.StatusType, error) {
	ret := _m.Called(namespace)

	var r0 container.StatusType
	if rf, ok := ret.Get(0).(func([]string) container.StatusType); ok {
		r0 = rf(namespace)
	} else {
		r0 = ret.Get(0).(container.StatusType)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]string) error); ok {
		r1 = rf(namespace)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// StopService provides a mock function with given fields: namespace
func (_m *Container) StopService(namespace []string) error {
	ret := _m.Called(namespace)

	var r0 error
	if rf, ok := ret.Get(0).(func([]string) error); ok {
		r0 = rf(namespace)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// TasksError provides a mock function with given fields: namespace
func (_m *Container) TasksError(namespace []string) ([]string, error) {
	ret := _m.Called(namespace)

	var r0 []string
	if rf, ok := ret.Get(0).(func([]string) []string); ok {
		r0 = rf(namespace)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]string) error); ok {
		r1 = rf(namespace)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
