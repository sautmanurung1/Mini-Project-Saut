package answare

import (
	"Tugas-Mini-Project/domains/mocks"
	"Tugas-Mini-Project/entities"
	"errors"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"io"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestCreateAnswareHandler(t *testing.T) {
	svc := mocks.AnswareService{}

	answareController := AnswareHandler{
		Svc: &svc,
	}

	t.Run("Success", func(t *testing.T) {
		svc.On("CreateAnswareService", mock.Anything).Return("Success To Create Answare", nil).Once()

		e := echo.New()
		r := httptest.NewRequest("POST", "/student/answare", nil)
		w := httptest.NewRecorder()

		echoContext := e.NewContext(r, w)
		err := answareController.CreateAnswareHandler(echoContext)

		if err != nil {
			return
		}

		assert.Equal(t, 200, w.Result().StatusCode)
	})

	t.Run("BadRequest", func(t *testing.T) {
		svc.On("CreateAnswareService", mock.Anything).Return("Server Have Bad Request", errors.New("Error To Create Discussions")).Once()

		e := echo.New()
		r := httptest.NewRequest("POST", "/student/answare", io.Reader(strings.NewReader(`{"Status" : "Bad Request"}`)))
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		err := answareController.CreateAnswareHandler(echoContext)

		if err != nil {
			return
		}

		assert.Equal(t, 400, w.Result().StatusCode)
	})

	t.Run("ServerError", func(t *testing.T) {
		svc.On("CreateAnswareService", mock.Anything).Return("Internal Server Error", errors.New("Error To Create Discussions")).Once()

		e := echo.New()
		r := httptest.NewRequest("POST", "/student/answare", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		err := answareController.CreateAnswareHandler(echoContext)

		if err != nil {
			return
		}

		assert.Equal(t, 500, w.Result().StatusCode)
	})
}

func TestGetAllAnswareHandler(t *testing.T) {
	svc := mocks.AnswareService{}

	answareService := new(mocks.AnswareService)
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

	answareController := AnswareHandler{
		Svc: &svc,
	}

	t.Run("Success", func(t *testing.T) {
		answareService.On("GetAllAnswareService").Return(answareData, nil).Once()
		service, err := answareService.GetAllAnswareService()

		svc.On("GetAllAnswareService").Return(answareData, nil).Once()
		e := echo.New()
		r := httptest.NewRequest("GET", "/answare", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		er := answareController.GetAllAnswareHandler(echoContext)

		if er != nil {
			return
		}

		assert.Equal(t, service, answareData)
		assert.Equal(t, 200, w.Result().StatusCode)
		assert.NoError(t, err)
	})

	t.Run("ServerError", func(t *testing.T) {
		answareService.On("GetAllAnswareService").Return(answareData, errors.New("Error To Get All Data Discussions")).Once()
		service, err := answareService.GetAllAnswareService()

		svc.On("GetAllAnswareService").Return(answareData, errors.New("Error To Get All Data Discussions")).Once()
		e := echo.New()
		r := httptest.NewRequest("GET", "/answare", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		er := answareController.GetAllAnswareHandler(echoContext)

		if er != nil {
			return
		}

		assert.Equal(t, service, answareData)
		assert.Equal(t, 500, w.Result().StatusCode)
		assert.Error(t, err)
	})
}

func TestGetAnswareByIdHandler(t *testing.T) {
	svc := mocks.AnswareService{}

	answareService := new(mocks.AnswareService)

	answareData := entities.Answare{
		QuestionId:  1,
		UserId:      2,
		Questions:   "Testing Question",
		AnswareUser: "Testing Answare",
		Name:        "Testing Name",
	}

	answareController := AnswareHandler{
		Svc: &svc,
	}

	t.Run("Success", func(t *testing.T) {
		answareService.On("GetAnswareByIdService", mock.Anything, mock.Anything).Return("Success To Get Answare By Id", nil).Once()
		service, err := answareService.GetAnswareByIdService(int(answareData.ID), answareData)

		svc.On("GetAnswareByIdService", mock.Anything, mock.Anything).Return("Success To Get Answare By Id", nil).Once()
		e := echo.New()
		r := httptest.NewRequest("GET", "/answare/:id", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		er := answareController.GetAnswareByIdHandler(echoContext)

		if er != nil {
			return
		}

		assert.Equal(t, service, "Success To Get Answare By Id")
		assert.Equal(t, 200, w.Result().StatusCode)
		assert.NoError(t, err)
	})

	t.Run("BadRequest", func(t *testing.T) {
		answareService.On("GetAnswareByIdService", mock.Anything, mock.Anything).Return("Failed To Get Answare By Id", errors.New("Error To Get Answare By Id")).Once()
		service, err := answareService.GetAnswareByIdService(int(answareData.ID), answareData)

		svc.On("GetAnswareByIdService", mock.Anything, mock.Anything).Return("Failed To Get Answare By Id", errors.New("Error To Get Answare By Id")).Once()
		e := echo.New()
		r := httptest.NewRequest("GET", "/answare/:id", io.Reader(strings.NewReader(`{"Status" : "Bad Request"}`)))
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		er := answareController.GetAnswareByIdHandler(echoContext)

		if er != nil {
			return
		}

		assert.Equal(t, service, "Failed To Get Answare By Id")
		assert.Equal(t, 400, w.Result().StatusCode)
		assert.Error(t, err)
	})
}

func TestUpdateAnswareHandler(t *testing.T) {
	svc := mocks.AnswareService{}

	answareService := new(mocks.AnswareService)

	answareData := entities.Answare{
		QuestionId:  1,
		UserId:      2,
		Questions:   "Testing Question",
		AnswareUser: "Testing Answare",
		Name:        "Testing Name",
	}

	answareController := AnswareHandler{
		Svc: &svc,
	}

	t.Run("Success", func(t *testing.T) {
		answareService.On("UpdateAnswareService", mock.Anything, mock.Anything).Return("Success To Update Answare By Id", nil).Once()
		service, err := answareService.UpdateAnswareService(int(answareData.ID), answareData)

		svc.On("UpdateAnswareService", mock.Anything, mock.Anything).Return("Success To Update Answare By Id", nil).Once()
		e := echo.New()
		r := httptest.NewRequest("PUT", "/student/answare/:id", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		er := answareController.UpdateAnswareHandler(echoContext)

		if er != nil {
			return
		}

		assert.Equal(t, service, "Success To Update Answare By Id")
		assert.Equal(t, 200, w.Result().StatusCode)
		assert.NoError(t, err)
	})

	t.Run("BadRequest", func(t *testing.T) {
		answareService.On("UpdateAnswareService", mock.Anything, mock.Anything).Return("Error To Update Answare By Id", errors.New("Error To Update Answare By Id")).Once()
		service, err := answareService.UpdateAnswareService(int(answareData.ID), answareData)

		svc.On("UpdateAnswareService", mock.Anything, mock.Anything).Return("Error To Update Answare By Id", errors.New("Error To Update Answare By Id")).Once()
		e := echo.New()
		r := httptest.NewRequest("PUT", "/student/answare/:id", io.Reader(strings.NewReader(`{"Status" : "Bad Request"}`)))
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		er := answareController.UpdateAnswareHandler(echoContext)

		if er != nil {
			return
		}

		assert.Equal(t, service, "Error To Update Answare By Id")
		assert.Equal(t, 400, w.Result().StatusCode)
		assert.Error(t, err)
	})

	t.Run("ServerError", func(t *testing.T) {
		answareService.On("UpdateAnswareService", mock.Anything, mock.Anything).Return("Error To Update Answare By Id", errors.New("Error To Update Answare By Id")).Once()
		service, err := answareService.UpdateAnswareService(int(answareData.ID), answareData)

		svc.On("UpdateAnswareService", mock.Anything, mock.Anything).Return("Error To Update Answare By Id", errors.New("Error To Update Answare By Id")).Once()
		e := echo.New()
		r := httptest.NewRequest("PUT", "/student/answare/:id", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		er := answareController.UpdateAnswareHandler(echoContext)

		if er != nil {
			return
		}

		assert.Equal(t, service, "Error To Update Answare By Id")
		assert.Equal(t, 500, w.Result().StatusCode)
		assert.Error(t, err)
	})
}

func TestDeleteAnswareHandler(t *testing.T) {
	svc := mocks.AnswareService{}

	answareService := new(mocks.AnswareService)

	answareData := entities.Answare{
		QuestionId:  1,
		UserId:      2,
		Questions:   "Testing Question",
		AnswareUser: "Testing Answare",
		Name:        "Testing Name",
	}

	answareController := AnswareHandler{
		Svc: &svc,
	}

	t.Run("Success", func(t *testing.T) {
		answareService.On("DeleteAnswareService", mock.Anything, mock.Anything).Return("Success To Update Answare By Id", nil).Once()
		service, err := answareService.DeleteAnswareService(int(answareData.ID), answareData)

		svc.On("DeleteAnswareService", mock.Anything, mock.Anything).Return("Success To Update Answare By Id", nil).Once()
		e := echo.New()
		r := httptest.NewRequest("DELETE", "/student/answare/:id", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		er := answareController.DeleteAnswareHandler(echoContext)

		if er != nil {
			return
		}

		assert.Equal(t, service, "Success To Update Answare By Id")
		assert.Equal(t, 200, w.Result().StatusCode)
		assert.NoError(t, err)
	})

	t.Run("BadRequest", func(t *testing.T) {
		answareService.On("DeleteAnswareService", mock.Anything, mock.Anything).Return("Error To Update Answare By Id", errors.New("Error To Update Answare By Id")).Once()
		service, err := answareService.DeleteAnswareService(int(answareData.ID), answareData)

		svc.On("DeleteAnswareService", mock.Anything, mock.Anything).Return("Error To Update Answare By Id", errors.New("Error To Update Answare By Id")).Once()
		e := echo.New()
		r := httptest.NewRequest("DELETE", "/student/answare/:id", io.Reader(strings.NewReader(`{"Status" : "Bad Request"}`)))
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		er := answareController.DeleteAnswareHandler(echoContext)

		if er != nil {
			return
		}

		assert.Equal(t, service, "Error To Update Answare By Id")
		assert.Equal(t, 400, w.Result().StatusCode)
		assert.Error(t, err)
	})

	t.Run("ServerError", func(t *testing.T) {
		answareService.On("DeleteAnswareService", mock.Anything, mock.Anything).Return("Error To Update Answare By Id", errors.New("Error To Update Answare By Id")).Once()
		service, err := answareService.DeleteAnswareService(int(answareData.ID), answareData)

		svc.On("DeleteAnswareService", mock.Anything, mock.Anything).Return("Error To Update Answare By Id", errors.New("Error To Update Answare By Id")).Once()
		e := echo.New()
		r := httptest.NewRequest("DELETE", "/student/answare/:id", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		er := answareController.DeleteAnswareHandler(echoContext)

		if er != nil {
			return
		}

		assert.Equal(t, service, "Error To Update Answare By Id")
		assert.Equal(t, 500, w.Result().StatusCode)
		assert.Error(t, err)
	})
}
