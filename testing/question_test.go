package testing

import (
	"Tugas-Mini-Project/domains"
	"Tugas-Mini-Project/domains/mocks"
	"Tugas-Mini-Project/entities"
	"Tugas-Mini-Project/interface/question"
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

func TestCreateQuestionService(t *testing.T) {
	svc := mocks.QuestionService{}

	questionController := question.QuestionHandler{
		Svc: &svc,
	}

	t.Run("Success", func(t *testing.T) {
		svc.On("CreateQuestionService", mock.Anything).Return("Success To Create Question", nil).Once()

		e := echo.New()
		r := httptest.NewRequest("POST", "/teacher/question", nil)
		w := httptest.NewRecorder()

		echoContext := e.NewContext(r, w)
		err := questionController.CreateQuestionHandler(echoContext)

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

		err := questionController.CreateQuestionHandler(echoContext)

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

		err := questionController.CreateQuestionHandler(echoContext)

		if err != nil {
			return
		}

		assert.Equal(t, 500, w.Result().StatusCode)
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

func TestGetAllQuestionService(t *testing.T) {
	svc := mocks.QuestionService{}

	questionService := new(mocks.QuestionService)
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

		er := questionController.GetAllQuestionsHandler(echoContext)

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

		er := questionController.GetAllQuestionsHandler(echoContext)

		if er != nil {
			return
		}

		assert.Equal(t, service, questionData)
		assert.Equal(t, 400, w.Result().StatusCode)
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

func TestGetQuestionByIdService(t *testing.T) {
	svc := mocks.QuestionService{}

	questionService := new(mocks.QuestionService)

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
		questionService.On("GetQuestionByIDService", mock.Anything, mock.Anything).Return("Success To Get Question By Id", nil).Once()
		service, err := questionService.GetQuestionByIDService(int(questionData.ID), questionData)

		svc.On("GetQuestionByIDService", mock.Anything, mock.Anything).Return("Success To Get Question By Id", nil).Once()
		e := echo.New()
		r := httptest.NewRequest("GET", "/question/:id", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		er := questionController.GetQuestionByIdHandler(echoContext)

		if er != nil {
			return
		}

		assert.Equal(t, service, "Success To Get Question By Id")
		assert.Equal(t, 200, w.Result().StatusCode)
		assert.NoError(t, err)
	})

	t.Run("BadRequest", func(t *testing.T) {
		questionService.On("GetQuestionByIDService", mock.Anything, mock.Anything).Return("Failed To Get Question By Id", errors.New("Error To Get Question By Id")).Once()
		service, err := questionService.GetQuestionByIDService(int(questionData.ID), questionData)

		svc.On("GetQuestionByIDService", mock.Anything, mock.Anything).Return("Failed To Get Assignment By Id", errors.New("Error To Get Question By Id")).Once()
		e := echo.New()
		r := httptest.NewRequest("GET", "/question/:id", io.Reader(strings.NewReader(`{"Status" : "Bad Request"}`)))
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		er := questionController.GetQuestionByIdHandler(echoContext)

		if er != nil {
			return
		}

		assert.Equal(t, service, "Failed To Get Question By Id")
		assert.Equal(t, 400, w.Result().StatusCode)
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

func TestUpdateQuestionService(t *testing.T) {
	svc := mocks.QuestionService{}

	questionService := new(mocks.QuestionService)
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
		questionService.On("UpdateQuestionService", mock.Anything, mock.Anything).Return("Success To Update Question By Id", nil).Once()
		service, err := questionService.UpdateQuestionService(int(questionData.ID), questionData)

		svc.On("UpdateQuestionService", mock.Anything, mock.Anything).Return("Success To Update Assignment By Id", nil).Once()
		e := echo.New()
		r := httptest.NewRequest("PUT", "/teacher/question/:id", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		er := questionController.UpdateQuestionHandler(echoContext)

		if er != nil {
			return
		}

		assert.Equal(t, service, "Success To Update Question By Id")
		assert.Equal(t, 200, w.Result().StatusCode)
		assert.NoError(t, err)
	})

	t.Run("BadRequest", func(t *testing.T) {
		questionService.On("UpdateQuestionService", mock.Anything, mock.Anything).Return("Error To Update Question By Id", errors.New("Error To Update Question By Id")).Once()
		service, err := questionService.UpdateQuestionService(int(questionData.ID), questionData)

		svc.On("UpdateQuestionService", mock.Anything, mock.Anything).Return("Error To Update Question By Id", errors.New("Error To Update Question By Id")).Once()
		e := echo.New()
		r := httptest.NewRequest("PUT", "/teacher/question/:id", io.Reader(strings.NewReader(`{"Status" : "Bad Request"}`)))
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		er := questionController.UpdateQuestionHandler(echoContext)

		if er != nil {
			return
		}

		assert.Equal(t, service, "Error To Update Question By Id")
		assert.Equal(t, 400, w.Result().StatusCode)
		assert.Error(t, err)
	})

	t.Run("ServerError", func(t *testing.T) {
		questionService.On("UpdateQuestionService", mock.Anything, mock.Anything).Return("Error To Update Question By Id", errors.New("Error To Update Question By Id")).Once()
		service, err := questionService.UpdateQuestionService(int(questionData.ID), questionData)

		svc.On("UpdateQuestionService", mock.Anything, mock.Anything).Return("Error To Update Question By Id", errors.New("Error To Update Question By Id")).Once()
		e := echo.New()
		r := httptest.NewRequest("PUT", "/teacher/question/:id", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		er := questionController.UpdateQuestionHandler(echoContext)

		if er != nil {
			return
		}

		assert.Equal(t, service, "Error To Update Question By Id")
		assert.Equal(t, 500, w.Result().StatusCode)
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

func TestDeleteQuestionService(t *testing.T) {
	svc := mocks.QuestionService{}

	questionService := new(mocks.QuestionService)
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

		er := questionController.DeleteQuestionHandler(echoContext)

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

		er := questionController.DeleteQuestionHandler(echoContext)

		if er != nil {
			return
		}

		assert.Equal(t, service, "Error To Delete Question By Id")
		assert.Equal(t, 500, w.Result().StatusCode)
		assert.Error(t, err)
	})
}
