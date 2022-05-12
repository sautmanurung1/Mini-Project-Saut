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
		err := answareRepo.GetAnswareById(1, answareData)

		assert.NoError(t, err)
	})

	t.Run("Failed", func(t *testing.T) {
		answareRepo.On("GetAnswareById", mock.Anything, mock.Anything).Return(errors.New("Error To Get Answare By Id")).Once()

		answareRepo := domains.AnswareRepository(answareRepo)
		err := answareRepo.GetAnswareById(1, answareData)

		assert.Error(t, err)
	})
}
