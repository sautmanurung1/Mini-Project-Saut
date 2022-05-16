package answare

import (
	"Tugas-Mini-Project/domains"
	"Tugas-Mini-Project/domains/mocks"
	"Tugas-Mini-Project/entities"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestCreateAnsware(t *testing.T) {
	answareRepository := new(mocks.AnswareRepository)
	answareData := entities.Answare{
		QuestionId:  1,
		UserId:      2,
		Questions:   "Testing Question",
		AnswareUser: "Testing Answare",
		Name:        "Testing Name",
	}

	t.Run("Success", func(t *testing.T) {
		answareRepository.On("CreateAnsware", mock.Anything).Return(nil).Once()

		answareRepository := domains.AnswareRepository(answareRepository)
		err := answareRepository.CreateAnsware(answareData)

		assert.NoError(t, err)
	})

	t.Run("Failed", func(t *testing.T) {
		answareRepository.On("CreateAnsware", mock.Anything).Return(errors.New("Error To Create Answare")).Once()

		answareRepository := domains.AnswareRepository(answareRepository)
		err := answareRepository.CreateAnsware(answareData)

		assert.Error(t, err)
	})
}

func TestGetAllAnsware(t *testing.T) {
	answareRepo := new(mocks.AnswareRepository)
	answareData := []entities.Answare{
		{
			QuestionId:  1,
			UserId:      2,
			Questions:   "Testing Question",
			AnswareUser: "Testing Answare",
			Name:        "Testing Name",
		},
		{
			QuestionId:  2,
			UserId:      3,
			Questions:   "Testing Question 1",
			AnswareUser: "Testing Answare 1",
			Name:        "Testing Name 1",
		},
	}

	t.Run("Success", func(t *testing.T) {
		answareRepo.On("GetAllAnsware").Return(answareData, nil).Once()

		answareRepo := domains.AnswareRepository(answareRepo)
		getAllAnsware, err := answareRepo.GetAllAnsware()

		assert.Equal(t, getAllAnsware, answareData)
		assert.NoError(t, err)
	})

	t.Run("Failed", func(t *testing.T) {
		answareRepo.On("GetAllAnsware").Return(answareData, errors.New("Error To Get All Answare")).Once()

		answareRepo := domains.AnswareRepository(answareRepo)
		getAllAnsware, err := answareRepo.GetAllAnsware()

		assert.Equal(t, getAllAnsware, answareData)
		assert.Error(t, err)
	})
}

func TestGetAnswareById(t *testing.T) {
	answareRepo := new(mocks.AnswareRepository)
	answareData := entities.Answare{
		QuestionId:  1,
		UserId:      2,
		Questions:   "Testing Question",
		AnswareUser: "Testing Answare",
		Name:        "Testing Name",
	}

	t.Run("Success", func(t *testing.T) {
		answareRepo.On("GetAnswareById", mock.Anything, mock.Anything).Return(nil).Once()

		answareRepo := domains.AnswareRepository(answareRepo)
		err := answareRepo.GetAnswareById(int(answareData.ID), answareData)

		assert.NoError(t, err)
	})

	t.Run("Failed", func(t *testing.T) {
		answareRepo.On("GetAnswareById", mock.Anything, mock.Anything).Return(errors.New("Error To Get Answare By Id")).Once()

		answareRepo := domains.AnswareRepository(answareRepo)
		err := answareRepo.GetAnswareById(int(answareData.ID), answareData)

		assert.Error(t, err)
	})
}

func TestUpdateAnsware(t *testing.T) {
	answareRepo := new(mocks.AnswareRepository)
	answareData := entities.Answare{
		QuestionId:  1,
		UserId:      2,
		Questions:   "Testing Question",
		AnswareUser: "Testing Answare",
		Name:        "Testing Name",
	}

	t.Run("Success", func(t *testing.T) {
		answareRepo.On("UpdateAnsware", mock.Anything, mock.Anything).Return(nil).Once()

		answareRepo := domains.AnswareRepository(answareRepo)
		err := answareRepo.UpdateAnsware(int(answareData.ID), answareData)

		assert.NoError(t, err)
	})

	t.Run("Failed", func(t *testing.T) {
		answareRepo.On("UpdateAnsware", mock.Anything, mock.Anything).Return(errors.New("Error To Update The Answare")).Once()

		answareRepo := domains.AnswareRepository(answareRepo)
		err := answareRepo.UpdateAnsware(int(answareData.ID), answareData)

		assert.Error(t, err)
	})
}

func TestDeleteAnsware(t *testing.T) {
	answareRepo := new(mocks.AnswareRepository)
	answareData := entities.Answare{
		QuestionId:  1,
		UserId:      2,
		Questions:   "Testing Question",
		AnswareUser: "Testing Answare",
		Name:        "Testing Name",
	}

	t.Run("Success", func(t *testing.T) {
		answareRepo.On("DeleteAnsware", mock.Anything, mock.Anything).Return(nil).Once()

		answareRepo := domains.AnswareRepository(answareRepo)
		err := answareRepo.DeleteAnsware(int(answareData.ID), answareData)

		assert.NoError(t, err)
	})

	t.Run("Failed", func(t *testing.T) {
		answareRepo.On("DeleteAnsware", mock.Anything, mock.Anything).Return(errors.New("Error To Update The Answare")).Once()

		answareRepo := domains.AnswareRepository(answareRepo)
		err := answareRepo.DeleteAnsware(int(answareData.ID), answareData)

		assert.Error(t, err)
	})
}
