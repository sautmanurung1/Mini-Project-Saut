package question_test

import (
	"Tugas-Mini-Project/domains"
	"Tugas-Mini-Project/entities"
	"Tugas-Mini-Project/testing/repo_test/question"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestCreateQuestionRepository(t *testing.T) {
	questionRepository := new(question.QuestionRepository)
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

func TestGetAllQuestionRepository(t *testing.T) {
	questionRepo := new(question.QuestionRepository)
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

func TestGetQuestionByIDRepository(t *testing.T) {
	questionRepository := new(question.QuestionRepository)
	questionData := entities.Question{
		AssignmentId:    1,
		AssignmentTitle: "Assignment Title Testing",
		UserId:          1,
		QuestionUser:    "User Testing",
		Name:            "Name Testing",
	}

	t.Run("Success", func(t *testing.T) {
		questionRepository.On("GetQuestionByID", mock.Anything).Return(questionData, nil).Once()

		questionRepository := domains.QuestionRepository(questionRepository)
		result, err := questionRepository.GetQuestionByID(int(questionData.ID))

		assert.Equal(t, result, questionData)
		assert.NoError(t, err)
	})

	t.Run("Failed", func(t *testing.T) {
		questionRepository.On("GetQuestionByID", mock.Anything).Return(questionData, errors.New("Error To Get Question By Id")).Once()

		questionRepository := domains.QuestionRepository(questionRepository)
		result, err := questionRepository.GetQuestionByID(int(questionData.ID))

		assert.Equal(t, result, questionData)
		assert.Error(t, err)
	})
}

func TestUpdateQuestionRepository(t *testing.T) {
	questionRepository := new(question.QuestionRepository)
	questionData := entities.Question{
		AssignmentId:    1,
		AssignmentTitle: "Assignment Title Testing",
		UserId:          1,
		QuestionUser:    "User Testing",
		Name:            "Name Testing",
	}

	t.Run("Success", func(t *testing.T) {
		questionRepository.On("UpdateQuestion", mock.Anything, mock.Anything).Return(questionData, nil).Once()

		questionRepository := domains.QuestionRepository(questionRepository)
		result, err := questionRepository.UpdateQuestion(int(questionData.ID), questionData)

		assert.Equal(t, result, questionData)
		assert.NoError(t, err)
	})

	t.Run("Failed", func(t *testing.T) {
		questionRepository.On("UpdateQuestion", mock.Anything, mock.Anything).Return(questionData, errors.New("Error To Update The Question")).Once()

		questionRepository := domains.QuestionRepository(questionRepository)
		result, err := questionRepository.UpdateQuestion(int(questionData.ID), questionData)

		assert.Equal(t, result, questionData)
		assert.Error(t, err)
	})
}

func TestDeleteQuestionRepository(t *testing.T) {
	questionRepository := new(question.QuestionRepository)
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
