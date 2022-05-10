// Code generated by mockery v2.12.2. DO NOT EDIT.

package mocks

import (
	entities "Tugas-Mini-Project/entities"

	mock "github.com/stretchr/testify/mock"
)

// DiscussionsService is an autogenerated mock type for the DiscussionsService type
type DiscussionsService struct {
	mock.Mock
}

// CreateDiscussionsService provides a mock function with given fields: discussions
func (_m *DiscussionsService) CreateDiscussionsService(discussions entities.Discussions) (string, error) {
	ret := _m.Called(discussions)

	var r0 string
	if rf, ok := ret.Get(0).(func(entities.Discussions) string); ok {
		r0 = rf(discussions)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(entities.Discussions) error); ok {
		r1 = rf(discussions)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAllDiscussionsService provides a mock function with given fields:
func (_m *DiscussionsService) GetAllDiscussionsService() ([]entities.Discussions, error) {
	ret := _m.Called()

	var r0 []entities.Discussions
	if rf, ok := ret.Get(0).(func() []entities.Discussions); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entities.Discussions)
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