package assignment

import (
	"Tugas-Mini-Project/domains"
	"Tugas-Mini-Project/entities"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestCreateAssignmentService(t *testing.T) {
	assignmentService := new(AssignmentService)
	assignmentData := entities.Assignment{
		UserId:      1,
		Description: "Description Testing",
		Questions:   "Question Testing",
		Name:        "Name Testing",
		Title:       "Title Testing",
	}

	t.Run("Success", func(t *testing.T) {
		assignmentService.On("CreateAssignmentService", mock.Anything).Return("Success Create Assignment", nil).Once()

		assignmentService := domains.AssignmentService(assignmentService)
		assign, err := assignmentService.CreateAssignmentService(assignmentData)

		assert.Equal(t, assign, "Success Create Assignment")
		assert.NoError(t, err)
	})

	t.Run("Failed", func(t *testing.T) {
		assignmentService.On("CreateAssignmentService", mock.Anything).Return("Failed To Create Assignment", errors.New("Error To make Unit Testing")).Once()

		assignmentService := domains.AssignmentService(assignmentService)
		assign, err := assignmentService.CreateAssignmentService(assignmentData)

		assert.Equal(t, assign, "Failed To Create Assignment")
		assert.Error(t, err)
	})
}

func TestGetAssignmentByIdService(t *testing.T) {
	assignmentService := new(AssignmentService)
	assignmentData := entities.Assignment{
		UserId:      1,
		Description: "Description Testing",
		Questions:   "Question Testing",
		Name:        "Name Testing",
		Title:       "Title Testing",
	}

	t.Run("Success", func(t *testing.T) {
		assignmentService.On("GetAssignmentByIdService", mock.Anything).Return(assignmentData, nil).Once()

		assignmentService := domains.AssignmentService(assignmentService)
		assign, err := assignmentService.GetAssignmentByIdService(int(assignmentData.ID))

		assert.Equal(t, assign, assignmentData)
		assert.NoError(t, err)
	})

	t.Run("Failed", func(t *testing.T) {
		assignmentService.On("GetAssignmentByIdService", mock.Anything).Return(assignmentData, errors.New("Error To Make Unit Testing")).Once()

		assignmentService := domains.AssignmentService(assignmentService)
		assign, err := assignmentService.GetAssignmentByIdService(int(assignmentData.ID))

		assert.Equal(t, assign, assignmentData)
		assert.Error(t, err)
	})
}

func TestGetAllAssignmentService(t *testing.T) {
	assignmentService := new(AssignmentService)
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
		assignmentService.On("GetAllAssignmentService").Return(assignmentData, nil).Once()

		assignmentService := domains.AssignmentService(assignmentService)
		getAllAssignment, err := assignmentService.GetAllAssignmentService()

		assert.Equal(t, getAllAssignment, assignmentData)
		assert.NoError(t, err)
	})

	t.Run("Failed", func(t *testing.T) {
		assignmentService.On("GetAllAssignmentService").Return(assignmentData, errors.New("Error To Unit Testing")).Once()

		assignmentService := domains.AssignmentService(assignmentService)
		getAllAssignment, err := assignmentService.GetAllAssignmentService()

		assert.Equal(t, getAllAssignment, assignmentData)
		assert.Error(t, err)
	})
}

func TestUpdateAssignmentService(t *testing.T) {
	assignmentService := new(AssignmentService)
	assignmentData := entities.Assignment{
		UserId:      1,
		Description: "Description Testing",
		Questions:   "Question Testing",
		Name:        "Name Testing",
		Title:       "Title Testing",
	}

	t.Run("Success", func(t *testing.T) {
		assignmentService.On("UpdateAssignmentService", mock.Anything, mock.Anything).Return(assignmentData, nil).Once()

		assignmentService := domains.AssignmentService(assignmentService)
		assign, err := assignmentService.UpdateAssignmentService(int(assignmentData.ID), assignmentData)

		assert.Equal(t, assign, assignmentData)
		assert.NoError(t, err)
	})

	t.Run("Failed", func(t *testing.T) {
		assignmentService.On("UpdateAssignmentService", mock.Anything, mock.Anything).Return(assignmentData, errors.New("Error To Make Unit Testing")).Once()

		assignmentService := domains.AssignmentService(assignmentService)
		assign, err := assignmentService.UpdateAssignmentService(int(assignmentData.ID), assignmentData)

		assert.Equal(t, assign, assignmentData)
		assert.Error(t, err)
	})
}

func TestDeleteAssignmentService(t *testing.T) {
	assignmentService := new(AssignmentService)
	assignmentData := entities.Assignment{
		UserId:      1,
		Description: "Description Testing",
		Questions:   "Question Testing",
		Name:        "Name Testing",
		Title:       "Title Testing",
	}

	t.Run("Success", func(t *testing.T) {
		assignmentService.On("DeleteAssignmentService", mock.Anything, mock.Anything).Return("Success Delete Assignment", nil).Once()

		assignmentService := domains.AssignmentService(assignmentService)
		assign, err := assignmentService.DeleteAssignmentService(int(assignmentData.ID), assignmentData)

		assert.Equal(t, assign, "Success Delete Assignment")
		assert.NoError(t, err)
	})

	t.Run("Failed", func(t *testing.T) {
		assignmentService.On("DeleteAssignmentService", mock.Anything, mock.Anything).Return("Failed Delete Assignment", errors.New("Error To Make Unit Testing")).Once()

		assignmentService := domains.AssignmentService(assignmentService)
		assign, err := assignmentService.DeleteAssignmentService(int(assignmentData.ID), assignmentData)

		assert.Equal(t, assign, "Failed Delete Assignment")
		assert.Error(t, err)
	})
}
