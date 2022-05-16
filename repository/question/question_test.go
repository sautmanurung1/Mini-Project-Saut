package question

import (
	"Tugas-Mini-Project/domains"
	"Tugas-Mini-Project/domains/mocks"
	"Tugas-Mini-Project/entities"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestCreateQuestion(t *testing.T) {
	questionRepository := new(mocks.QuestionRepository)
	questionData := entities.Question{
		AssignmentId:    1,
		AssignmentTitle: "Assignment Title Testing",
		UserId:          1,
		QuestionUser:    "User Testing",
		Name:            "Name Testing",
	}

	t.Run("Success", func(t *testing.T) {
		questionRepository.On("CreateQuestion", mock.Anything).Return(nil).Once()

		questionRepository := domains.QuestionRepository(questionRepository)
		err := questionRepository.CreateQuestion(questionData)

		assert.NoError(t, err)
	})

	t.Run("Failed", func(t *testing.T) {
		questionRepository.On("CreateQuestion", mock.Anything).Return(errors.New("Error To Create Question")).Once()

		questionRepository := domains.QuestionRepository(questionRepository)
		err := questionRepository.CreateQuestion(questionData)

		assert.Error(t, err)
	})
}

func TestGetAllQuestion(t *testing.T) {
	questionRepo := new(mocks.QuestionRepository)
	questionData := []entities.Question{
		{
			AssignmentId:    1,
			AssignmentTitle: "Assignment Title Testing",
			UserId:          1,
			QuestionUser:    "User Testing",
			Name:            "Name Testing",
		},
		{
			AssignmentId:    2,
			AssignmentTitle: "Assignment Title Testing",
			UserId:          2,
			QuestionUser:    "User Testing",
			Name:            "Name Testing",
		},
	}

	t.Run("Success", func(t *testing.T) {
		questionRepo.On("GetAllQuestion").Return(questionData, nil).Once()

		questionRepo := domains.QuestionRepository(questionRepo)
		getAllQuestion, err := questionRepo.GetAllQuestion()

		assert.Equal(t, getAllQuestion, questionData)
		assert.NoError(t, err)
	})

	t.Run("Failed", func(t *testing.T) {
		questionRepo.On("GetAllQuestion").Return(questionData, errors.New("Error To Get All Question")).Once()

		questionRepo := domains.QuestionRepository(questionRepo)
		getAllQuestion, err := questionRepo.GetAllQuestion()

		assert.Equal(t, getAllQuestion, questionData)
		assert.Error(t, err)
	})
}

func TestGetQuestionByID(t *testing.T) {
	questionRepository := new(mocks.QuestionRepository)
	questionData := entities.Question{
		AssignmentId:    1,
		AssignmentTitle: "Assignment Title Testing",
		UserId:          1,
		QuestionUser:    "User Testing",
		Name:            "Name Testing",
	}

	t.Run("Success", func(t *testing.T) {
		questionRepository.On("GetQuestionByID", mock.Anything, mock.Anything).Return(nil).Once()

		questionRepository := domains.QuestionRepository(questionRepository)
		err := questionRepository.GetQuestionByID(int(questionData.ID), questionData)

		assert.NoError(t, err)
	})

	t.Run("Failed", func(t *testing.T) {
		questionRepository.On("GetQuestionByID", mock.Anything, mock.Anything).Return(errors.New("Error To Get Question By Id")).Once()

		questionRepository := domains.QuestionRepository(questionRepository)
		err := questionRepository.GetQuestionByID(int(questionData.ID), questionData)

		assert.Error(t, err)
	})
}

func TestUpdateQuestion(t *testing.T) {
	questionRepository := new(mocks.QuestionRepository)
	questionData := entities.Question{
		AssignmentId:    1,
		AssignmentTitle: "Assignment Title Testing",
		UserId:          1,
		QuestionUser:    "User Testing",
		Name:            "Name Testing",
	}

	t.Run("Success", func(t *testing.T) {
		questionRepository.On("UpdateQuestion", mock.Anything, mock.Anything).Return(nil).Once()

		questionRepository := domains.QuestionRepository(questionRepository)
		err := questionRepository.UpdateQuestion(int(questionData.ID), questionData)

		assert.NoError(t, err)
	})

	t.Run("Failed", func(t *testing.T) {
		questionRepository.On("UpdateQuestion", mock.Anything, mock.Anything).Return(errors.New("Error To Update The Question")).Once()

		questionRepository := domains.QuestionRepository(questionRepository)
		err := questionRepository.UpdateQuestion(int(questionData.ID), questionData)

		assert.Error(t, err)
	})
}

func TestDeleteQuestion(t *testing.T) {
	questionRepository := new(mocks.QuestionRepository)
	questionData := entities.Question{
		AssignmentId:    1,
		AssignmentTitle: "Assignment Title Testing",
		UserId:          1,
		QuestionUser:    "User Testing",
		Name:            "Name Testing",
	}

	t.Run("Success", func(t *testing.T) {
		questionRepository.On("DeleteQuestion", mock.Anything, mock.Anything).Return(nil).Once()

		questionRepository := domains.QuestionRepository(questionRepository)
		err := questionRepository.DeleteQuestion(int(questionData.ID), questionData)

		assert.NoError(t, err)
	})

	t.Run("Failed", func(t *testing.T) {
		questionRepository.On("DeleteQuestion", mock.Anything, mock.Anything).Return(errors.New("Error To Delete Question")).Once()

		questionRepository := domains.QuestionRepository(questionRepository)
		err := questionRepository.DeleteQuestion(int(questionData.ID), questionData)

		assert.Error(t, err)
	})
}
