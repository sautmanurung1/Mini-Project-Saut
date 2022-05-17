package assignment_test

import (
	"Tugas-Mini-Project/domains"
	"Tugas-Mini-Project/entities"
	"Tugas-Mini-Project/testing/repo_test/assignment"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestCreateAssignmentRepository(t *testing.T) {
	assignmentRepository := new(assignment.AssignmentRepository)
	assignmentData := entities.Assignment{
		UserId:      1,
		Description: "Description Testing",
		Questions:   "Question Testing",
		Name:        "Name Testing",
		Title:       "Title Testing",
	}

	t.Run("Success", func(t *testing.T) {
		assignmentRepository.On("CreateAssignment", mock.Anything).Return(nil).Once()

		assignmentRepository := domains.AssignmentRepository(assignmentRepository)
		err := assignmentRepository.CreateAssignment(assignmentData)

		assert.NoError(t, err)
	})

	t.Run("Failed", func(t *testing.T) {
		assignmentRepository.On("CreateAssignment", mock.Anything).Return(errors.New("Error To Create Assignment")).Once()

		assignmentRepository := domains.AssignmentRepository(assignmentRepository)
		err := assignmentRepository.CreateAssignment(assignmentData)

		assert.Error(t, err)
	})
}

func TestGetAllAssignmentRepository(t *testing.T) {
	assignmentRepo := new(assignment.AssignmentRepository)
	assignmentData := []entities.Assignment{
		{
			UserId:      1,
			Description: "Description Testing",
			Questions:   "Question Testing",
			Name:        "Name Testing",
			Title:       "Title Testing",
		},
		{
			UserId:      2,
			Description: "Description Testing 1",
			Questions:   "Question Testing 1",
			Name:        "Name Testing 1",
			Title:       "Title Testing 1",
		},
	}

	t.Run("Success", func(t *testing.T) {
		assignmentRepo.On("GetAllAssignment").Return(assignmentData, nil).Once()

		assignmentRepo := domains.AssignmentRepository(assignmentRepo)
		getAllAssignment, err := assignmentRepo.GetAllAssignment()

		assert.Equal(t, getAllAssignment, assignmentData)
		assert.NoError(t, err)
	})

	t.Run("Failed", func(t *testing.T) {
		assignmentRepo.On("GetAllAssignment").Return(assignmentData, errors.New("Error To Get All Assginment")).Once()

		assignmentRepo := domains.AssignmentRepository(assignmentRepo)
		getAllAssignment, err := assignmentRepo.GetAllAssignment()

		assert.Equal(t, getAllAssignment, assignmentData)
		assert.Error(t, err)
	})
}

func TestGetAssignmentByIdRepository(t *testing.T) {
	assignmentRepository := new(assignment.AssignmentRepository)
	assignmentData := entities.Assignment{
		UserId:      1,
		Description: "Description Testing",
		Questions:   "Question Testing",
		Name:        "Name Testing",
		Title:       "Title Testing",
	}

	t.Run("Success", func(t *testing.T) {
		assignmentRepository.On("GetAssignmentById", mock.Anything).Return(assignmentData, nil).Once()

		assignmentRepository := domains.AssignmentRepository(assignmentRepository)
		assign, err := assignmentRepository.GetAssignmentById(int(assignmentData.ID))

		assert.Equal(t, assign, assignmentData)
		assert.NoError(t, err)
	})

	t.Run("Failed", func(t *testing.T) {
		assignmentRepository.On("GetAssignmentById", mock.Anything).Return(assignmentData, errors.New("Error To Get Assignment By Id")).Once()

		assignmentRepository := domains.AssignmentRepository(assignmentRepository)
		assign, err := assignmentRepository.GetAssignmentById(int(assignmentData.ID))

		assert.Equal(t, assign, assignmentData)
		assert.Error(t, err)
	})
}

func TestUpdateAssignmentRepository(t *testing.T) {
	assignmentRepository := new(assignment.AssignmentRepository)
	assignmentData := entities.Assignment{
		UserId:      1,
		Description: "Description Testing",
		Questions:   "Question Testing",
		Name:        "Name Testing",
		Title:       "Title Testing",
	}

	t.Run("Success", func(t *testing.T) {
		assignmentRepository.On("UpdateAssignment", mock.Anything, mock.Anything).Return(assignmentData, nil).Once()

		assignmentRepository := domains.AssignmentRepository(assignmentRepository)
		assign, err := assignmentRepository.UpdateAssignment(int(assignmentData.ID), assignmentData)

		assert.Equal(t, assign, assignmentData)
		assert.NoError(t, err)
	})

	t.Run("Failed", func(t *testing.T) {
		assignmentRepository.On("UpdateAssignment", mock.Anything, mock.Anything).Return(assignmentData, errors.New("Error To Update The Assignment")).Once()

		assignmentRepository := domains.AssignmentRepository(assignmentRepository)
		assign, err := assignmentRepository.UpdateAssignment(int(assignmentData.ID), assignmentData)

		assert.Equal(t, assign, assignmentData)
		assert.Error(t, err)
	})
}

func TestDeleteAssignmentRepository(t *testing.T) {
	assignmentRepository := new(assignment.AssignmentRepository)
	assignmentData := entities.Assignment{
		UserId:      1,
		Description: "Description Testing",
		Questions:   "Question Testing",
		Name:        "Name Testing",
		Title:       "Title Testing",
	}

	t.Run("Success", func(t *testing.T) {
		assignmentRepository.On("DeleteAssignment", mock.Anything, mock.Anything).Return(nil).Once()

		assignmentRepository := domains.AssignmentRepository(assignmentRepository)
		err := assignmentRepository.DeleteAssignment(int(assignmentData.ID), assignmentData)

		assert.NoError(t, err)
	})

	t.Run("Failed", func(t *testing.T) {
		assignmentRepository.On("DeleteAssignment", mock.Anything, mock.Anything).Return(errors.New("Error To Delete Assignment")).Once()

		assignmentRepository := domains.AssignmentRepository(assignmentRepository)
		err := assignmentRepository.DeleteAssignment(int(assignmentData.ID), assignmentData)

		assert.Error(t, err)
	})
}
