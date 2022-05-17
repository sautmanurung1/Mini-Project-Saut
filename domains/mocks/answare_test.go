package mocks

import (
	"Tugas-Mini-Project/domains"
	"Tugas-Mini-Project/entities"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestCreateAnswareRepository(t *testing.T) {
	answareRepository := new(AnswareRepository)
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

func TestGetAllAnswareRepository(t *testing.T) {
	answareRepo := new(AnswareRepository)
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

func TestGetAnswareByIdRepository(t *testing.T) {
	answareRepo := new(AnswareRepository)
	answareData := entities.Answare{
		QuestionId:  1,
		UserId:      2,
		Questions:   "Testing Question",
		AnswareUser: "Testing Answare",
		Name:        "Testing Name",
	}

	t.Run("Success", func(t *testing.T) {
		answareRepo.On("GetAnswareById", mock.Anything).Return(answareData, nil).Once()

		answareRepo := domains.AnswareRepository(answareRepo)
		answare, err := answareRepo.GetAnswareById(int(answareData.ID))

		assert.Equal(t, answare, answareData)
		assert.NoError(t, err)
	})

	t.Run("Failed", func(t *testing.T) {
		answareRepo.On("GetAnswareById", mock.Anything, mock.Anything).Return(answareData, errors.New("Error To Get Answare By Id")).Once()

		answareRepo := domains.AnswareRepository(answareRepo)
		answare, err := answareRepo.GetAnswareById(int(answareData.ID))

		assert.Equal(t, answare, answareData)
		assert.Error(t, err)
	})
}

func TestUpdateAnswareRepository(t *testing.T) {
	answareRepo := new(AnswareRepository)
	answareData := entities.Answare{
		QuestionId:  1,
		UserId:      2,
		Questions:   "Testing Question",
		AnswareUser: "Testing Answare",
		Name:        "Testing Name",
	}

	t.Run("Success", func(t *testing.T) {
		answareRepo.On("UpdateAnsware", mock.Anything, mock.Anything).Return(answareData, nil).Once()

		answareRepo := domains.AnswareRepository(answareRepo)
		answare, err := answareRepo.UpdateAnsware(int(answareData.ID), answareData)

		assert.Equal(t, answare, answareData)
		assert.NoError(t, err)
	})

	t.Run("Failed", func(t *testing.T) {
		answareRepo.On("UpdateAnsware", mock.Anything, mock.Anything).Return(answareData, errors.New("Error To Update The Answare")).Once()

		answareRepo := domains.AnswareRepository(answareRepo)
		answare, err := answareRepo.UpdateAnsware(int(answareData.ID), answareData)

		assert.Equal(t, answare, answareData)
		assert.Error(t, err)
	})
}

func TestDeleteAnswareRepository(t *testing.T) {
	answareRepo := new(AnswareRepository)
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

func TestCreateAnswareService(t *testing.T) {
	answareService := new(AnswareService)
	answareData := entities.Answare{
		QuestionId:  1,
		UserId:      2,
		Questions:   "Testing Question",
		AnswareUser: "Testing Answare",
		Name:        "Testing Name",
	}

	t.Run("Success", func(t *testing.T) {
		answareService.On("CreateAnswareService", mock.Anything).Return("Success To Create Answare", nil).Once()

		answareService := domains.AnswareService(answareService)
		answare, err := answareService.CreateAnswareService(answareData)

		assert.Equal(t, answare, "Success To Create Answare")
		assert.NoError(t, err)
	})

	t.Run("Failed", func(t *testing.T) {
		answareService.On("CreateAnswareService", mock.Anything).Return("Error To Create Answare", errors.New("Error To Create Answare")).Once()

		answareService := domains.AnswareService(answareService)
		answare, err := answareService.CreateAnswareService(answareData)

		assert.Equal(t, answare, "Error To Create Answare")
		assert.Error(t, err)
	})
}

func TestGetAllAnswareService(t *testing.T) {
	answareService := new(AnswareService)
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
		answareService.On("GetAllAnswareService").Return(answareData, nil).Once()

		answareService := domains.AnswareService(answareService)
		getAllAnsware, err := answareService.GetAllAnswareService()

		assert.Equal(t, getAllAnsware, answareData)
		assert.NoError(t, err)
	})

	t.Run("Failed", func(t *testing.T) {
		answareService.On("GetAllAnswareService").Return(answareData, errors.New("Error To Get All Answare")).Once()

		answareService := domains.AnswareService(answareService)
		getAllAnsware, err := answareService.GetAllAnswareService()

		assert.Equal(t, getAllAnsware, answareData)
		assert.Error(t, err)
	})
}

func TestGetAnswareByIdService(t *testing.T) {
	answareService := new(AnswareService)
	answareData := entities.Answare{
		QuestionId:  1,
		UserId:      2,
		Questions:   "Testing Question",
		AnswareUser: "Testing Answare",
		Name:        "Testing Name",
	}

	t.Run("Success", func(t *testing.T) {
		answareService.On("GetAnswareByIdService", mock.Anything, mock.Anything).Return(answareData, nil).Once()

		answareService := domains.AnswareService(answareService)
		answare, err := answareService.GetAnswareByIdService(int(answareData.ID))

		assert.Equal(t, answare, answareData)
		assert.NoError(t, err)
	})

	t.Run("Failed", func(t *testing.T) {
		answareService.On("GetAnswareByIdService", mock.Anything, mock.Anything).Return(answareData, errors.New("Error to Unit Testing")).Once()

		answareService := domains.AnswareService(answareService)
		answare, err := answareService.GetAnswareByIdService(int(answareData.ID))

		assert.Equal(t, answare, answareData)
		assert.Error(t, err)
	})
}

func TestUpdateAnswareService(t *testing.T) {
	answareService := new(AnswareService)
	answareData := entities.Answare{
		QuestionId:  1,
		UserId:      2,
		Questions:   "Testing Question",
		AnswareUser: "Testing Answare",
		Name:        "Testing Name",
	}

	t.Run("Success", func(t *testing.T) {
		answareService.On("UpdateAnswareService", mock.Anything, mock.Anything).Return(answareData, nil).Once()

		answareService := domains.AnswareService(answareService)
		answare, err := answareService.UpdateAnswareService(int(answareData.ID), answareData)

		assert.Equal(t, answare, answareData)
		assert.NoError(t, err)
	})

	t.Run("Failed", func(t *testing.T) {
		answareService.On("UpdateAnswareService", mock.Anything, mock.Anything).Return(answareData, errors.New("Error to Unit Testing")).Once()

		answareService := domains.AnswareService(answareService)
		answare, err := answareService.UpdateAnswareService(int(answareData.ID), answareData)

		assert.Equal(t, answare, answareData)
		assert.Error(t, err)
	})
}

func TestDeleteAnswareService(t *testing.T) {
	answareService := new(AnswareService)
	answareData := entities.Answare{
		QuestionId:  1,
		UserId:      2,
		Questions:   "Testing Question",
		AnswareUser: "Testing Answare",
		Name:        "Testing Name",
	}

	t.Run("Success", func(t *testing.T) {
		answareService.On("DeleteAnswareService", mock.Anything, mock.Anything).Return("Succes Delete Answare", nil).Once()

		answareService := domains.AnswareService(answareService)
		answare, err := answareService.DeleteAnswareService(int(answareData.ID), answareData)

		assert.Equal(t, answare, "Succes Delete Answare")
		assert.NoError(t, err)
	})

	t.Run("Failed", func(t *testing.T) {
		answareService.On("DeleteAnswareService", mock.Anything, mock.Anything).Return("Failed To Delete Answare", errors.New("Error to Unit Testing")).Once()

		answareService := domains.AnswareService(answareService)
		answare, err := answareService.DeleteAnswareService(int(answareData.ID), answareData)

		assert.Equal(t, answare, "Failed To Delete Answare")
		assert.Error(t, err)
	})
}
