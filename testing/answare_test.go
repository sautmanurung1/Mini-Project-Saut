package testing

import (
	"Tugas-Mini-Project/domains"
	"Tugas-Mini-Project/domains/mocks"
	"Tugas-Mini-Project/entities"
	answares "Tugas-Mini-Project/interface/answare"
	"errors"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"io"
	"net/http/httptest"
	"strings"
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

func TestCreateAnswareService(t *testing.T) {
	svc := mocks.AnswareService{}

	answareController := answares.AnswareHandler{
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

func TestGetAllAnswareService(t *testing.T) {
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

	answareController := answares.AnswareHandler{
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

func TestGetAnswareByIdService(t *testing.T) {
	svc := mocks.AnswareService{}

	answareService := new(mocks.AnswareService)

	answareData := entities.Answare{
		QuestionId:  1,
		UserId:      2,
		Questions:   "Testing Question",
		AnswareUser: "Testing Answare",
		Name:        "Testing Name",
	}

	answareController := answares.AnswareHandler{
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

func TestUpdateAnswareService(t *testing.T) {
	svc := mocks.AnswareService{}

	answareService := new(mocks.AnswareService)

	answareData := entities.Answare{
		QuestionId:  1,
		UserId:      2,
		Questions:   "Testing Question",
		AnswareUser: "Testing Answare",
		Name:        "Testing Name",
	}

	answareController := answares.AnswareHandler{
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

func TestDeleteAnswareService(t *testing.T) {
	svc := mocks.AnswareService{}

	answareService := new(mocks.AnswareService)

	answareData := entities.Answare{
		QuestionId:  1,
		UserId:      2,
		Questions:   "Testing Question",
		AnswareUser: "Testing Answare",
		Name:        "Testing Name",
	}

	answareController := answares.AnswareHandler{
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
