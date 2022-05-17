package question

import (
	"Tugas-Mini-Project/domains"
	"Tugas-Mini-Project/entities"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestCreateQuestionService(t *testing.T) {
	questionService := new(QuestionService)
	questionData := entities.Question{
		AssignmentId:    1,
		AssignmentTitle: "Assignment Title Testing",
		UserId:          1,
		QuestionUser:    "User Testing",
		Name:            "Name Testing",
	}

	t.Run("Success", func(t *testing.T) {
		questionService.On("CreateQuestionService", mock.Anything).Return("Success Create Question", nil).Once()

		questionService := domains.QuestionService(questionService)
		question, err := questionService.CreateQuestionService(questionData)

		assert.Equal(t, question, "Success Create Question")
		assert.NoError(t, err)
	})

	t.Run("Failed", func(t *testing.T) {
		questionService.On("CreateQuestionService", mock.Anything).Return("Failed Create Question", errors.New("Error To Make Unit Testing")).Once()

		questionService := domains.QuestionService(questionService)
		question, err := questionService.CreateQuestionService(questionData)

		assert.Equal(t, question, "Failed Create Question")
		assert.Error(t, err)
	})
}

func TestGetQuestionByIdService(t *testing.T) {
	questionService := new(QuestionService)
	questionData := entities.Question{
		AssignmentId:    1,
		AssignmentTitle: "Assignment Title Testing",
		UserId:          1,
		QuestionUser:    "User Testing",
		Name:            "Name Testing",
	}

	t.Run("Success", func(t *testing.T) {
		questionService.On("GetQuestionByIDService", mock.Anything, mock.Anything).Return(questionData, nil).Once()

		questionService := domains.QuestionService(questionService)
		question, err := questionService.GetQuestionByIDService(int(questionData.ID))

		assert.Equal(t, question, questionData)
		assert.NoError(t, err)
	})

	t.Run("Failed", func(t *testing.T) {
		questionService.On("GetQuestionByIDService", mock.Anything, mock.Anything).Return(questionData, errors.New("Error To Make Unit Testing")).Once()

		questionService := domains.QuestionService(questionService)
		question, err := questionService.GetQuestionByIDService(int(questionData.ID))

		assert.Equal(t, question, questionData)
		assert.Error(t, err)
	})
}

func TestGetAllQuestionService(t *testing.T) {
	questionService := new(QuestionService)
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
		questionService.On("GetAllQuestionService").Return(questionData, nil).Once()

		questionService := domains.QuestionService(questionService)
		getAllQuestion, err := questionService.GetAllQuestionService()

		assert.Equal(t, getAllQuestion, questionData)
		assert.NoError(t, err)
	})

	t.Run("Failed", func(t *testing.T) {
		questionService.On("GetAllQuestionService").Return(questionData, errors.New("Error To Unit Testing")).Once()

		questionService := domains.QuestionService(questionService)
		getAllQuestion, err := questionService.GetAllQuestionService()

		assert.Equal(t, getAllQuestion, questionData)
		assert.Error(t, err)
	})
}

func TestUpdateQuestionService(t *testing.T) {
	questionService := new(QuestionService)
	questionData := entities.Question{
		AssignmentId:    1,
		AssignmentTitle: "Assignment Title Testing",
		UserId:          1,
		QuestionUser:    "User Testing",
		Name:            "Name Testing",
	}

	t.Run("Success", func(t *testing.T) {
		questionService.On("UpdateQuestionService", mock.Anything, mock.Anything).Return(questionData, nil).Once()

		questionService := domains.QuestionService(questionService)
		question, err := questionService.UpdateQuestionService(int(questionData.ID), questionData)

		assert.Equal(t, question, questionData)
		assert.NoError(t, err)
	})

	t.Run("Failed", func(t *testing.T) {
		questionService.On("UpdateQuestionService", mock.Anything, mock.Anything).Return(questionData, errors.New("Error To Make Unit Testing")).Once()

		questionService := domains.QuestionService(questionService)
		question, err := questionService.UpdateQuestionService(int(questionData.ID), questionData)

		assert.Equal(t, question, questionData)
		assert.Error(t, err)
	})
}

func TestDeleteQuestionService(t *testing.T) {
	questionService := new(QuestionService)
	questionData := entities.Question{
		AssignmentId:    1,
		AssignmentTitle: "Assignment Title Testing",
		UserId:          1,
		QuestionUser:    "User Testing",
		Name:            "Name Testing",
	}

	t.Run("Success", func(t *testing.T) {
		questionService.On("DeleteQuestionService", mock.Anything, mock.Anything).Return("Success Delete Question", nil).Once()

		questionService := domains.QuestionService(questionService)
		question, err := questionService.DeleteQuestionService(int(questionData.ID), questionData)

		assert.Equal(t, question, "Success Delete Question")
		assert.NoError(t, err)
	})

	t.Run("Failed", func(t *testing.T) {
		questionService.On("DeleteQuestionService", mock.Anything, mock.Anything).Return("Failed Delete Question", errors.New("Error To Make Unit Testing")).Once()

		questionService := domains.QuestionService(questionService)
		question, err := questionService.DeleteQuestionService(int(questionData.ID), questionData)

		assert.Equal(t, question, "Failed Delete Question")
		assert.Error(t, err)
	})
}
