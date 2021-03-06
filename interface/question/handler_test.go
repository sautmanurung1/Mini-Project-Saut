package question_test

import (
	"Tugas-Mini-Project/entities"
	"Tugas-Mini-Project/interface/question"
	question2 "Tugas-Mini-Project/testing/service_test/question"
	"errors"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"io"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestCreateQuestion(t *testing.T) {
	svc := question2.QuestionService{}

	questionController := question.QuestionHandler{
		Svc: &svc,
	}

	t.Run("Success", func(t *testing.T) {
		svc.On("CreateQuestionService", mock.Anything).Return("Success To Create Question", nil).Once()

		e := echo.New()
		r := httptest.NewRequest("POST", "/teacher/question", nil)
		w := httptest.NewRecorder()

		echoContext := e.NewContext(r, w)
		err := questionController.CreateQuestion(echoContext)

		if err != nil {
			return
		}

		assert.Equal(t, 201, w.Result().StatusCode)
	})

	t.Run("BadRequest", func(t *testing.T) {
		svc.On("CreateQuestionService", mock.Anything).Return("Server Have Bad Request", errors.New("Error To Create Question")).Once()

		e := echo.New()
		r := httptest.NewRequest("POST", "/teacher/assignment", io.Reader(strings.NewReader(`{"Status" : "Bad Request"}`)))
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		err := questionController.CreateQuestion(echoContext)

		if err != nil {
			return
		}

		assert.Equal(t, 400, w.Result().StatusCode)
	})

	t.Run("ServerError", func(t *testing.T) {
		svc.On("CreateQuestionService", mock.Anything).Return("Internal Server Error", errors.New("Error To Create Question")).Once()

		e := echo.New()
		r := httptest.NewRequest("POST", "/teacher/question", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		err := questionController.CreateQuestion(echoContext)

		if err != nil {
			return
		}

		assert.Equal(t, 500, w.Result().StatusCode)
	})
}

func TestGetAllQuestion(t *testing.T) {
	svc := question2.QuestionService{}

	questionService := new(question2.QuestionService)
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

	questionController := question.QuestionHandler{
		Svc: &svc,
	}

	t.Run("Success", func(t *testing.T) {
		questionService.On("GetAllQuestionService").Return(questionData, nil).Once()
		service, err := questionService.GetAllQuestionService()

		svc.On("GetAllQuestionService").Return(questionData, nil).Once()
		e := echo.New()
		r := httptest.NewRequest("GET", "/question", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		er := questionController.GetAllQuestions(echoContext)

		if er != nil {
			return
		}

		assert.Equal(t, service, questionData)
		assert.Equal(t, 200, w.Result().StatusCode)
		assert.NoError(t, err)
	})

	t.Run("BadRequest", func(t *testing.T) {
		questionService.On("GetAllQuestionService").Return(questionData, errors.New("Error To Get All Data Question")).Once()
		service, err := questionService.GetAllQuestionService()

		svc.On("GetAllQuestionService").Return(questionData, errors.New("Error To Get All Data Question")).Once()
		e := echo.New()
		r := httptest.NewRequest("GET", "/question", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		er := questionController.GetAllQuestions(echoContext)

		if er != nil {
			return
		}

		assert.Equal(t, service, questionData)
		assert.Equal(t, 400, w.Result().StatusCode)
		assert.Error(t, err)
	})
}

func TestGetQuestionById(t *testing.T) {
	svc := question2.QuestionService{}

	questionService := new(question2.QuestionService)

	questionData := entities.Question{
		AssignmentId:    1,
		AssignmentTitle: "Assignment Title Testing",
		UserId:          1,
		QuestionUser:    "User Testing",
		Name:            "Name Testing",
	}

	questionController := question.QuestionHandler{
		Svc: &svc,
	}

	t.Run("Success", func(t *testing.T) {
		questionService.On("GetQuestionByIDService", mock.Anything, mock.Anything).Return(questionData, nil).Once()
		service, err := questionService.GetQuestionByIDService(int(questionData.ID))

		svc.On("GetQuestionByIDService", mock.Anything, mock.Anything).Return(questionData, nil).Once()
		e := echo.New()
		r := httptest.NewRequest("GET", "/question/:id", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		er := questionController.GetQuestionById(echoContext)

		if er != nil {
			return
		}

		assert.Equal(t, service, questionData)
		assert.Equal(t, 200, w.Result().StatusCode)
		assert.NoError(t, err)
	})

	t.Run("BadRequest", func(t *testing.T) {
		questionService.On("GetQuestionByIDService", mock.Anything, mock.Anything).Return(questionData, errors.New("Error To Get Question By Id")).Once()
		service, err := questionService.GetQuestionByIDService(int(questionData.ID))

		svc.On("GetQuestionByIDService", mock.Anything, mock.Anything).Return(questionData, errors.New("Error To Get Question By Id")).Once()
		e := echo.New()
		r := httptest.NewRequest("GET", "/question/:id", io.Reader(strings.NewReader(`{"Status" : "Bad Request"}`)))
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		er := questionController.GetQuestionById(echoContext)

		if er != nil {
			return
		}

		assert.Equal(t, service, questionData)
		assert.Equal(t, 400, w.Result().StatusCode)
		assert.Error(t, err)
	})
}

func TestUpdateQuestion(t *testing.T) {
	svc := question2.QuestionService{}

	questionService := new(question2.QuestionService)
	questionData := entities.Question{
		AssignmentId:    1,
		AssignmentTitle: "Assignment Title Testing",
		UserId:          1,
		QuestionUser:    "User Testing",
		Name:            "Name Testing",
	}

	questionController := question.QuestionHandler{
		Svc: &svc,
	}

	t.Run("Success", func(t *testing.T) {
		questionService.On("UpdateQuestionService", mock.Anything, mock.Anything).Return(questionData, nil).Once()
		service, err := questionService.UpdateQuestionService(int(questionData.ID), questionData)

		svc.On("UpdateQuestionService", mock.Anything, mock.Anything).Return(questionData, nil).Once()
		e := echo.New()
		r := httptest.NewRequest("PUT", "/teacher/question/:id", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		er := questionController.UpdateQuestion(echoContext)

		if er != nil {
			return
		}

		assert.Equal(t, service, questionData)
		assert.Equal(t, 200, w.Result().StatusCode)
		assert.NoError(t, err)
	})

	t.Run("BadRequest", func(t *testing.T) {
		questionService.On("UpdateQuestionService", mock.Anything, mock.Anything).Return(questionData, errors.New("Error To Update Question By Id")).Once()
		service, err := questionService.UpdateQuestionService(int(questionData.ID), questionData)

		svc.On("UpdateQuestionService", mock.Anything, mock.Anything).Return(questionData, errors.New("Error To Update Question By Id")).Once()
		e := echo.New()
		r := httptest.NewRequest("PUT", "/teacher/question/:id", io.Reader(strings.NewReader(`{"Status" : "Bad Request"}`)))
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		er := questionController.UpdateQuestion(echoContext)

		if er != nil {
			return
		}

		assert.Equal(t, service, questionData)
		assert.Equal(t, 400, w.Result().StatusCode)
		assert.Error(t, err)
	})

	t.Run("ServerError", func(t *testing.T) {
		questionService.On("UpdateQuestionService", mock.Anything, mock.Anything).Return(questionData, errors.New("Error To Update Question By Id")).Once()
		service, err := questionService.UpdateQuestionService(int(questionData.ID), questionData)

		svc.On("UpdateQuestionService", mock.Anything, mock.Anything).Return(questionData, errors.New("Error To Update Question By Id")).Once()
		e := echo.New()
		r := httptest.NewRequest("PUT", "/teacher/question/:id", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		er := questionController.UpdateQuestion(echoContext)

		if er != nil {
			return
		}

		assert.Equal(t, service, questionData)
		assert.Equal(t, 500, w.Result().StatusCode)
		assert.Error(t, err)
	})
}

func TestDeleteQuestion(t *testing.T) {
	svc := question2.QuestionService{}

	questionService := new(question2.QuestionService)
	questionData := entities.Question{
		AssignmentId:    1,
		AssignmentTitle: "Assignment Title Testing",
		UserId:          1,
		QuestionUser:    "User Testing",
		Name:            "Name Testing",
	}

	questionController := question.QuestionHandler{
		Svc: &svc,
	}

	t.Run("Success", func(t *testing.T) {
		questionService.On("DeleteQuestionService", mock.Anything, mock.Anything).Return("Success To Delete Question By Id", nil).Once()
		service, err := questionService.DeleteQuestionService(int(questionData.ID), questionData)

		svc.On("DeleteQuestionService", mock.Anything, mock.Anything).Return("Success To Delete Question By Id", nil).Once()
		e := echo.New()
		r := httptest.NewRequest("DELETE", "/teacher/question/:id", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		er := questionController.DeleteQuestion(echoContext)

		if er != nil {
			return
		}

		assert.Equal(t, service, "Success To Delete Question By Id")
		assert.Equal(t, 200, w.Result().StatusCode)
		assert.NoError(t, err)
	})

	t.Run("ServerError", func(t *testing.T) {
		questionService.On("DeleteQuestionService", mock.Anything, mock.Anything).Return("Error To Delete Question By Id", errors.New("Error To Delete Question By Id")).Once()
		service, err := questionService.DeleteQuestionService(int(questionData.ID), questionData)

		svc.On("DeleteQuestionService", mock.Anything, mock.Anything).Return("Error To Delete Question By Id", errors.New("Error To Delete Question By Id")).Once()
		e := echo.New()
		r := httptest.NewRequest("DELETE", "/teacher/question/:id", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		er := questionController.DeleteQuestion(echoContext)

		if er != nil {
			return
		}

		assert.Equal(t, service, "Error To Delete Question By Id")
		assert.Equal(t, 500, w.Result().StatusCode)
		assert.Error(t, err)
	})
}
