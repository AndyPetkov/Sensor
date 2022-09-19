// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	models "server/models"

	mock "github.com/stretchr/testify/mock"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

//var BaseExecutorDevice Repository = &Repository{}
// Create provides a mock function with given fields: device, sensor
func (_m *Repository) Create(device models.Device, sensor models.Sensor) (interface{}, error) {
	ret := _m.Called(device, sensor)

	var r0 interface{}
	if rf, ok := ret.Get(0).(func(models.Device, models.Sensor) interface{}); ok {
		r0 = rf(device, sensor)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(models.Device, models.Sensor) error); ok {
		r1 = rf(device, sensor)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: id
func (_m *Repository) Delete(id string) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAll provides a mock function with given fields:
func (_m *Repository) GetAll() (interface{}, error) {
	ret := _m.Called()

	var r0 interface{}
	if rf, ok := ret.Get(0).(func() interface{}); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetById provides a mock function with given fields: id
func (_m *Repository) GetById(id string) (interface{}, error) {
	ret := _m.Called(id)

	var r0 interface{}
	if rf, ok := ret.Get(0).(func(string) interface{}); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: id, device, sensor
func (_m *Repository) Update(id string, device models.Device, sensor models.Sensor) (interface{}, error) {
	ret := _m.Called(id, device, sensor)

	var r0 interface{}
	if rf, ok := ret.Get(0).(func(string, models.Device, models.Sensor) interface{}); ok {
		r0 = rf(id, device, sensor)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, models.Device, models.Sensor) error); ok {
		r1 = rf(id, device, sensor)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func Test(id string) {
	instance := Repository{}
	instance.GetAll()
}
