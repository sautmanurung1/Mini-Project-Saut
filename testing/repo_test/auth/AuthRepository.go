// Code generated by mockery v2.12.2. DO NOT EDIT.

package auth

import (
	entities "Tugas-Mini-Project/entities"

	mock "github.com/stretchr/testify/mock"
)

// AuthRepository is an autogenerated mock type for the AuthRepository type
type AuthRepository struct {
	mock.Mock
}

// Login provides a mock function with given fields: username
func (_m *AuthRepository) Login(username string) (entities.User, error) {
	ret := _m.Called(username)

	var r0 entities.User
	if rf, ok := ret.Get(0).(func(string) entities.User); ok {
		r0 = rf(username)
	} else {
		r0 = ret.Get(0).(entities.User)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(username)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Register provides a mock function with given fields: user
func (_m *AuthRepository) Register(user entities.User) error {
	ret := _m.Called(user)

	var r0 error
	if rf, ok := ret.Get(0).(func(entities.User) error); ok {
		r0 = rf(user)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}