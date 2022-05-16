// Code generated by mocks v2.12.2. DO NOT EDIT.

package mocks

import (
	entities "Tugas-Mini-Project/entities"

	mock "github.com/stretchr/testify/mock"
)

// DiscussionsRepository is an autogenerated mock type for the DiscussionsRepository type
type DiscussionsRepository struct {
	mock.Mock
}

// CreateDiscussions provides a mock function with given fields: discussions
func (_m *DiscussionsRepository) CreateDiscussions(discussions entities.Discussions) error {
	ret := _m.Called(discussions)

	var r0 error
	if rf, ok := ret.Get(0).(func(entities.Discussions) error); ok {
		r0 = rf(discussions)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAllDiscussions provides a mock function with given fields:
func (_m *DiscussionsRepository) GetAllDiscussions() ([]entities.Discussions, error) {
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
