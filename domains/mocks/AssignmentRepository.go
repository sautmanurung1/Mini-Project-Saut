// Code generated by mocks v2.12.2. DO NOT EDIT.

package mocks

import (
	entities "Tugas-Mini-Project/entities"

	mock "github.com/stretchr/testify/mock"
)

// AssignmentRepository is an autogenerated mock type for the AssignmentRepository type
type AssignmentRepository struct {
	mock.Mock
}

// CreateAssignment provides a mock function with given fields: assignment
func (_m *AssignmentRepository) CreateAssignment(assignment entities.Assignment) error {
	ret := _m.Called(assignment)

	var r0 error
	if rf, ok := ret.Get(0).(func(entities.Assignment) error); ok {
		r0 = rf(assignment)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteAssignment provides a mock function with given fields: id, assign
func (_m *AssignmentRepository) DeleteAssignment(id int, assign entities.Assignment) error {
	ret := _m.Called(id, assign)

	var r0 error
	if rf, ok := ret.Get(0).(func(int, entities.Assignment) error); ok {
		r0 = rf(id, assign)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAllAssignment provides a mock function with given fields:
func (_m *AssignmentRepository) GetAllAssignment() ([]entities.Assignment, error) {
	ret := _m.Called()

	var r0 []entities.Assignment
	if rf, ok := ret.Get(0).(func() []entities.Assignment); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entities.Assignment)
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

// GetAssignmentById provides a mock function with given fields: id, assign
func (_m *AssignmentRepository) GetAssignmentById(id int, assign entities.Assignment) error {
	ret := _m.Called(id, assign)

	var r0 error
	if rf, ok := ret.Get(0).(func(int, entities.Assignment) error); ok {
		r0 = rf(id, assign)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateAssignment provides a mock function with given fields: id, assignment
func (_m *AssignmentRepository) UpdateAssignment(id int, assignment entities.Assignment) error {
	ret := _m.Called(id, assignment)

	var r0 error
	if rf, ok := ret.Get(0).(func(int, entities.Assignment) error); ok {
		r0 = rf(id, assignment)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
