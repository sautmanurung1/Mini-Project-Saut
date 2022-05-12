package testing_test

import (
	"Tugas-Mini-Project/domains"
	"Tugas-Mini-Project/domains/mocks"
	"Tugas-Mini-Project/entities"
	"Tugas-Mini-Project/interface/discussions"
	"errors"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"io"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestCreateDiscussions(t *testing.T) {
	discussionsRepo := new(mocks.DiscussionsRepository)
	discussionsData := entities.Discussions{
		UserId:     1,
		QuestionId: 1,
		AnswareId:  1,
	}

	t.Run("Success", func(t *testing.T) {
		discussionsRepo.On("CreateDiscussions", mock.Anything).Return(nil).Once()

		discussionsRepo := domains.DiscussionsRepository(discussionsRepo)
		err := discussionsRepo.CreateDiscussions(discussionsData)
		assert.NoError(t, err)
	})

	t.Run("Failed", func(t *testing.T) {
		discussionsRepo.On("CreateDiscussions", mock.Anything).Return(errors.New("Error To Create Discussions")).Once()

		discussionsRepo := domains.DiscussionsRepository(discussionsRepo)
		err := discussionsRepo.CreateDiscussions(discussionsData)
		assert.Error(t, err)
	})
}

func TestGetAllDiscussions(t *testing.T) {
	discussionsRepo := new(mocks.DiscussionsRepository)
	discussionsData := []entities.Discussions{
		{
			UserId:       1,
			Name:         "User Testing",
			QuestionId:   1,
			QuestionUser: "Question Testing",
			AnswareId:    1,
			AnswereUser:  "Answare Testing",
		},
	}

	t.Run("Success", func(t *testing.T) {
		discussionsRepo.On("GetAllDiscussions").Return(discussionsData, nil).Once()

		discussRepo := domains.DiscussionsRepository(discussionsRepo)
		getAllDiscuss, err := discussRepo.GetAllDiscussions()

		assert.Equal(t, getAllDiscuss, discussionsData)
		assert.NoError(t, err)
	})

	t.Run("Failed", func(t *testing.T) {
		discussionsRepo.On("GetAllDiscussions").Return(discussionsData, errors.New("Error To Create Discussions")).Once()

		discussRepo := domains.DiscussionsRepository(discussionsRepo)
		getAllDiscuss, err := discussRepo.GetAllDiscussions()

		assert.Equal(t, getAllDiscuss, discussionsData)
		assert.Error(t, err)
	})
}

func TestCreateDiscussionsService(t *testing.T) {
	svc := mocks.DiscussionsService{}

	discussController := discussions.DiscussionHandler{
		Svc: &svc,
	}

	t.Run("Success", func(t *testing.T) {
		svc.On("CreateDiscussionsService", mock.Anything).Return("Success Create Discussions", nil).Once()

		e := echo.New()
		r := httptest.NewRequest("POST", "/discussions", nil)
		w := httptest.NewRecorder()

		echoContext := e.NewContext(r, w)
		err := discussController.CreateDiscsussions(echoContext)

		if err != nil {
			return
		}

		assert.Equal(t, 200, w.Result().StatusCode)
	})

	t.Run("BadRequest", func(t *testing.T) {
		svc.On("CreateDiscussionsService", mock.Anything).Return("Server Have Bad Request", errors.New("Error To Create Discussions")).Once()

		e := echo.New()
		r := httptest.NewRequest("POST", "/discussions", io.Reader(strings.NewReader(`{"Status" : "Bad Request"}`)))
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		err := discussController.CreateDiscsussions(echoContext)

		if err != nil {
			return
		}

		assert.Equal(t, 400, w.Result().StatusCode)
	})

	t.Run("ServerError", func(t *testing.T) {
		svc.On("CreateDiscussionsService", mock.Anything).Return("Internal Server Error", errors.New("Error To Create Discussions")).Once()

		e := echo.New()
		r := httptest.NewRequest("POST", "/discussions", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		err := discussController.CreateDiscsussions(echoContext)

		if err != nil {
			return
		}

		assert.Equal(t, 500, w.Result().StatusCode)
	})
}

func TestGetAllDiscussionsService(t *testing.T) {
	svc := mocks.DiscussionsService{}

	discussionService := new(mocks.DiscussionsService)
	discussionsData := []entities.Discussions{
		{
			UserId:       1,
			Name:         "User Testing",
			QuestionId:   1,
			QuestionUser: "Question Testing",
			AnswareId:    1,
			AnswereUser:  "Answare Testing",
		},
	}

	discussController := discussions.DiscussionHandler{
		Svc: &svc,
	}

	t.Run("Success", func(t *testing.T) {
		discussionService.On("GetAllDiscussionsService").Return(discussionsData, nil).Once()
		service, err := discussionService.GetAllDiscussionsService()

		svc.On("GetAllDiscussionsService").Return(discussionsData, nil).Once()
		e := echo.New()
		r := httptest.NewRequest("GET", "/discussions", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		er := discussController.GetAllDiscussions(echoContext)

		if er != nil {
			return
		}

		assert.Equal(t, service, discussionsData)
		assert.Equal(t, 200, w.Result().StatusCode)
		assert.NoError(t, err)
	})

	t.Run("BadRequest", func(t *testing.T) {
		discussionService.On("GetAllDiscussionsService").Return(discussionsData, errors.New("Error To Get All Data Discussions")).Once()
		service, err := discussionService.GetAllDiscussionsService()

		svc.On("GetAllDiscussionsService").Return(discussionsData, errors.New("Error To Get All Data Discussions")).Once()
		e := echo.New()
		r := httptest.NewRequest("GET", "/discussions", io.Reader(strings.NewReader(`{"Status" : "Bad Request"}`)))
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		er := discussController.GetAllDiscussions(echoContext)

		if er != nil {
			return
		}

		assert.Equal(t, service, discussionsData)
		assert.Equal(t, 400, w.Result().StatusCode)
		assert.Error(t, err)
	})
}
