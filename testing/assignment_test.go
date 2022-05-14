package testing

import (
	"Tugas-Mini-Project/domains"
	"Tugas-Mini-Project/domains/mocks"
	"Tugas-Mini-Project/entities"
	"Tugas-Mini-Project/interface/assignment"
	"errors"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"io"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestCreateAssignment(t *testing.T) {
	assignmentRepository := new(mocks.AssignmentRepository)
	assignmentData := entities.Assignment{
		UserId:      1,
		Description: "Description Testing",
		Questions:   "Question Testing",
		Name:        "Name Testing",
		Title:       "Title Testing",
	}

	t.Run("Success", func(t *testing.T) {
		assignmentRepository.On("CreateAssignment", mock.Anything).Return(nil).Once()

		assignmentRepository := domains.AssignmentRepository(assignmentRepository)
		err := assignmentRepository.CreateAssignment(assignmentData)

		assert.NoError(t, err)
	})

	t.Run("Failed", func(t *testing.T) {
		assignmentRepository.On("CreateAssignment", mock.Anything).Return(errors.New("Error To Create Assignment")).Once()

		assignmentRepository := domains.AssignmentRepository(assignmentRepository)
		err := assignmentRepository.CreateAssignment(assignmentData)

		assert.Error(t, err)
	})
}

func TestCreateAssignmentService(t *testing.T) {
	svc := mocks.AssignmentService{}

	assignmentController := assignment.AssignmentHandler{
		Svc: &svc,
	}

	t.Run("Success", func(t *testing.T) {
		svc.On("CreateAssignmentService", mock.Anything).Return("Success To Create Assignment", nil).Once()

		e := echo.New()
		r := httptest.NewRequest("POST", "/teacher/assignment", nil)
		w := httptest.NewRecorder()

		echoContext := e.NewContext(r, w)
		err := assignmentController.CreateAssignmentHandler(echoContext)

		if err != nil {
			return
		}

		assert.Equal(t, 200, w.Result().StatusCode)
	})

	t.Run("BadRequest", func(t *testing.T) {
		svc.On("CreateAssignmentService", mock.Anything).Return("Server Have Bad Request", errors.New("Error To Create Assignment")).Once()

		e := echo.New()
		r := httptest.NewRequest("POST", "/teacher/assignment", io.Reader(strings.NewReader(`{"Status" : "Bad Request"}`)))
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		err := assignmentController.CreateAssignmentHandler(echoContext)

		if err != nil {
			return
		}

		assert.Equal(t, 400, w.Result().StatusCode)
	})

	t.Run("ServerError", func(t *testing.T) {
		svc.On("CreateAssignmentService", mock.Anything).Return("Internal Server Error", errors.New("Error To Create Assignment")).Once()

		e := echo.New()
		r := httptest.NewRequest("POST", "/teacher/Assignment", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		err := assignmentController.CreateAssignmentHandler(echoContext)

		if err != nil {
			return
		}

		assert.Equal(t, 500, w.Result().StatusCode)
	})
}

func TestGetAllAssignment(t *testing.T) {
	assignmentRepo := new(mocks.AssignmentRepository)
	assignmentData := []entities.Assignment{
		{
			UserId:      1,
			Description: "Description Testing",
			Questions:   "Question Testing",
			Name:        "Name Testing",
			Title:       "Title Testing",
		},
		{
			UserId:      2,
			Description: "Description Testing 1",
			Questions:   "Question Testing 1",
			Name:        "Name Testing 1",
			Title:       "Title Testing 1",
		},
	}

	t.Run("Success", func(t *testing.T) {
		assignmentRepo.On("GetAllAssignment").Return(assignmentData, nil).Once()

		assignmentRepo := domains.AssignmentRepository(assignmentRepo)
		getAllAssignment, err := assignmentRepo.GetAllAssignment()

		assert.Equal(t, getAllAssignment, assignmentData)
		assert.NoError(t, err)
	})

	t.Run("Failed", func(t *testing.T) {
		assignmentRepo.On("GetAllAssignment").Return(assignmentData, errors.New("Error To Get All Assginment")).Once()

		assignmentRepo := domains.AssignmentRepository(assignmentRepo)
		getAllAssignment, err := assignmentRepo.GetAllAssignment()

		assert.Equal(t, getAllAssignment, assignmentData)
		assert.Error(t, err)
	})
}

func TestGetAllAssignmentService(t *testing.T) {
	svc := mocks.AssignmentService{}

	assignmentService := new(mocks.AssignmentService)
	assignmentData := []entities.Assignment{
		{
			UserId:      1,
			Description: "Description Testing",
			Questions:   "Question Testing",
			Name:        "Name Testing",
			Title:       "Title Testing",
		},
		{
			UserId:      2,
			Description: "Description Testing 1",
			Questions:   "Question Testing 1",
			Name:        "Name Testing 1",
			Title:       "Title Testing 1",
		},
	}

	assignmentController := assignment.AssignmentHandler{
		Svc: &svc,
	}

	t.Run("Success", func(t *testing.T) {
		assignmentService.On("GetAllAssignmentService").Return(assignmentData, nil).Once()
		service, err := assignmentService.GetAllAssignmentService()

		svc.On("GetAllAssignmentService").Return(assignmentData, nil).Once()
		e := echo.New()
		r := httptest.NewRequest("GET", "/assignment", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		er := assignmentController.GetAllAssignmentHandler(echoContext)

		if er != nil {
			return
		}

		assert.Equal(t, service, assignmentData)
		assert.Equal(t, 200, w.Result().StatusCode)
		assert.NoError(t, err)
	})

	t.Run("ServerError", func(t *testing.T) {
		assignmentService.On("GetAllAssignmentService").Return(assignmentData, errors.New("Error To Get All Data Assignment")).Once()
		service, err := assignmentService.GetAllAssignmentService()

		svc.On("GetAllAssignmentService").Return(assignmentData, errors.New("Error To Get All Data Assignment")).Once()
		e := echo.New()
		r := httptest.NewRequest("GET", "/assignment", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		er := assignmentController.GetAllAssignmentHandler(echoContext)

		if er != nil {
			return
		}

		assert.Equal(t, service, assignmentData)
		assert.Equal(t, 500, w.Result().StatusCode)
		assert.Error(t, err)
	})
}

func TestGetAssignmentById(t *testing.T) {
	assignmentRepository := new(mocks.AssignmentRepository)
	assignmentData := entities.Assignment{
		UserId:      1,
		Description: "Description Testing",
		Questions:   "Question Testing",
		Name:        "Name Testing",
		Title:       "Title Testing",
	}

	t.Run("Success", func(t *testing.T) {
		assignmentRepository.On("GetAssignmentById", mock.Anything, mock.Anything).Return(nil).Once()

		assignmentRepository := domains.AssignmentRepository(assignmentRepository)
		err := assignmentRepository.GetAssignmentById(int(assignmentData.ID), assignmentData)

		assert.NoError(t, err)
	})

	t.Run("Failed", func(t *testing.T) {
		assignmentRepository.On("GetAssignmentById", mock.Anything, mock.Anything).Return(errors.New("Error To Get Assignment By Id")).Once()

		assignmentRepository := domains.AssignmentRepository(assignmentRepository)
		err := assignmentRepository.GetAssignmentById(int(assignmentData.ID), assignmentData)

		assert.Error(t, err)
	})
}

func TestGetAssignmentByIdService(t *testing.T) {
	svc := mocks.AssignmentService{}

	assignmentService := new(mocks.AssignmentService)

	assignmentData := entities.Assignment{
		UserId:      1,
		Description: "Description Testing",
		Questions:   "Question Testing",
		Name:        "Name Testing",
		Title:       "Title Testing",
	}

	assignmentController := assignment.AssignmentHandler{
		Svc: &svc,
	}

	t.Run("Success", func(t *testing.T) {
		assignmentService.On("GetAssignmentByIdService", mock.Anything, mock.Anything).Return("Success To Get Assignment By Id", nil).Once()
		service, err := assignmentService.GetAssignmentByIdService(int(assignmentData.ID), assignmentData)

		svc.On("GetAssignmentByIdService", mock.Anything, mock.Anything).Return("Success To Get Assignment By Id", nil).Once()
		e := echo.New()
		r := httptest.NewRequest("GET", "/assignment/:id", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		er := assignmentController.GetAssignmentByIdHandler(echoContext)

		if er != nil {
			return
		}

		assert.Equal(t, service, "Success To Get Assignment By Id")
		assert.Equal(t, 200, w.Result().StatusCode)
		assert.NoError(t, err)
	})

	t.Run("BadRequest", func(t *testing.T) {
		assignmentService.On("GetAssignmentByIdService", mock.Anything, mock.Anything).Return("Failed To Get Assignment By Id", errors.New("Error To Get Assignment By Id")).Once()
		service, err := assignmentService.GetAssignmentByIdService(int(assignmentData.ID), assignmentData)

		svc.On("GetAssignmentByIdService", mock.Anything, mock.Anything).Return("Failed To Get Assignment By Id", errors.New("Error To Get Assignment By Id")).Once()
		e := echo.New()
		r := httptest.NewRequest("GET", "/assignment/:id", io.Reader(strings.NewReader(`{"Status" : "Bad Request"}`)))
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		er := assignmentController.GetAssignmentByIdHandler(echoContext)

		if er != nil {
			return
		}

		assert.Equal(t, service, "Failed To Get Assignment By Id")
		assert.Equal(t, 400, w.Result().StatusCode)
		assert.Error(t, err)
	})
}

func TestUpdateAssignment(t *testing.T) {
	assignmentRepository := new(mocks.AssignmentRepository)
	assignmentData := entities.Assignment{
		UserId:      1,
		Description: "Description Testing",
		Questions:   "Question Testing",
		Name:        "Name Testing",
		Title:       "Title Testing",
	}

	t.Run("Success", func(t *testing.T) {
		assignmentRepository.On("UpdateAssignment", mock.Anything, mock.Anything).Return(nil).Once()

		assignmentRepository := domains.AssignmentRepository(assignmentRepository)
		err := assignmentRepository.UpdateAssignment(int(assignmentData.ID), assignmentData)

		assert.NoError(t, err)
	})

	t.Run("Failed", func(t *testing.T) {
		assignmentRepository.On("UpdateAssignment", mock.Anything, mock.Anything).Return(errors.New("Error To Update The Assignment")).Once()

		assignmentRepository := domains.AssignmentRepository(assignmentRepository)
		err := assignmentRepository.UpdateAssignment(int(assignmentData.ID), assignmentData)

		assert.Error(t, err)
	})
}

func TestUpdateAssignmentService(t *testing.T) {
	svc := mocks.AssignmentService{}

	assignmentService := new(mocks.AssignmentService)
	assignmentData := entities.Assignment{
		UserId:      1,
		Description: "Description Testing",
		Questions:   "Question Testing",
		Name:        "Name Testing",
		Title:       "Title Testing",
	}

	assignmentController := assignment.AssignmentHandler{
		Svc: &svc,
	}

	t.Run("Success", func(t *testing.T) {
		assignmentService.On("UpdateAssignmentService", mock.Anything, mock.Anything).Return("Success To Update Assignment By Id", nil).Once()
		service, err := assignmentService.UpdateAssignmentService(int(assignmentData.ID), assignmentData)

		svc.On("UpdateAssignmentService", mock.Anything, mock.Anything).Return("Success To Update Assignment By Id", nil).Once()
		e := echo.New()
		r := httptest.NewRequest("PUT", "/teacher/assignment/:id", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		er := assignmentController.UpdateAssignmentHandler(echoContext)

		if er != nil {
			return
		}

		assert.Equal(t, service, "Success To Update Assignment By Id")
		assert.Equal(t, 200, w.Result().StatusCode)
		assert.NoError(t, err)
	})

	t.Run("BadRequest", func(t *testing.T) {
		assignmentService.On("UpdateAssignmentService", mock.Anything, mock.Anything).Return("Error To Update Assignment By Id", errors.New("Error To Update Assignment By Id")).Once()
		service, err := assignmentService.UpdateAssignmentService(int(assignmentData.ID), assignmentData)

		svc.On("UpdateAssignmentService", mock.Anything, mock.Anything).Return("Error To Update Assignment By Id", errors.New("Error To Update Assignment By Id")).Once()
		e := echo.New()
		r := httptest.NewRequest("PUT", "/teacher/assignment/:id", io.Reader(strings.NewReader(`{"Status" : "Bad Request"}`)))
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		er := assignmentController.UpdateAssignmentHandler(echoContext)

		if er != nil {
			return
		}

		assert.Equal(t, service, "Error To Update Assignment By Id")
		assert.Equal(t, 400, w.Result().StatusCode)
		assert.Error(t, err)
	})

	t.Run("ServerError", func(t *testing.T) {
		assignmentService.On("UpdateAssignmentService", mock.Anything, mock.Anything).Return("Error To Update Assignment By Id", errors.New("Error To Update Assignment By Id")).Once()
		service, err := assignmentService.UpdateAssignmentService(int(assignmentData.ID), assignmentData)

		svc.On("UpdateAssignmentService", mock.Anything, mock.Anything).Return("Error To Update Assignment By Id", errors.New("Error To Update Assignment By Id")).Once()
		e := echo.New()
		r := httptest.NewRequest("PUT", "/teacher/assignment/:id", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		er := assignmentController.UpdateAssignmentHandler(echoContext)

		if er != nil {
			return
		}

		assert.Equal(t, service, "Error To Update Assignment By Id")
		assert.Equal(t, 500, w.Result().StatusCode)
		assert.Error(t, err)
	})
}

func TestDeleteAssignment(t *testing.T) {
	assignmentRepository := new(mocks.AssignmentRepository)
	assignmentData := entities.Assignment{
		UserId:      1,
		Description: "Description Testing",
		Questions:   "Question Testing",
		Name:        "Name Testing",
		Title:       "Title Testing",
	}

	t.Run("Success", func(t *testing.T) {
		assignmentRepository.On("DeleteAssignment", mock.Anything, mock.Anything).Return(nil).Once()

		assignmentRepository := domains.AssignmentRepository(assignmentRepository)
		err := assignmentRepository.DeleteAssignment(int(assignmentData.ID), assignmentData)

		assert.NoError(t, err)
	})

	t.Run("Failed", func(t *testing.T) {
		assignmentRepository.On("DeleteAssignment", mock.Anything, mock.Anything).Return(errors.New("Error To Delete Assignment")).Once()

		assignmentRepository := domains.AssignmentRepository(assignmentRepository)
		err := assignmentRepository.DeleteAssignment(int(assignmentData.ID), assignmentData)

		assert.Error(t, err)
	})
}

func TestDeleteAssignmentService(t *testing.T) {
	svc := mocks.AssignmentService{}

	assignmentService := new(mocks.AssignmentService)
	assignmentData := entities.Assignment{
		UserId:      1,
		Description: "Description Testing",
		Questions:   "Question Testing",
		Name:        "Name Testing",
		Title:       "Title Testing",
	}

	assignmentController := assignment.AssignmentHandler{
		Svc: &svc,
	}

	t.Run("Success", func(t *testing.T) {
		assignmentService.On("DeleteAssignmentService", mock.Anything, mock.Anything).Return("Success To Delete Assignment By Id", nil).Once()
		service, err := assignmentService.DeleteAssignmentService(int(assignmentData.ID), assignmentData)

		svc.On("DeleteAssignmentService", mock.Anything, mock.Anything).Return("Success To Delete Assignment By Id", nil).Once()
		e := echo.New()
		r := httptest.NewRequest("DELETE", "/teacher/assignment/:id", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		er := assignmentController.DeleteAssignmentHandler(echoContext)

		if er != nil {
			return
		}

		assert.Equal(t, service, "Success To Delete Assignment By Id")
		assert.Equal(t, 200, w.Result().StatusCode)
		assert.NoError(t, err)
	})

	t.Run("BadRequest", func(t *testing.T) {
		assignmentService.On("DeleteAssignmentService", mock.Anything, mock.Anything).Return("Error To Delete Assignment By Id", errors.New("Error To Delete Assignment By Id")).Once()
		service, err := assignmentService.DeleteAssignmentService(int(assignmentData.ID), assignmentData)

		svc.On("DeleteAnswareService", mock.Anything, mock.Anything).Return("Error To Delete Assignment By Id", errors.New("Error To Delete Assignment By Id")).Once()
		e := echo.New()
		r := httptest.NewRequest("DELETE", "/teacher/assignment/:id", io.Reader(strings.NewReader(`{"Status" : "Bad Request"}`)))
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		er := assignmentController.DeleteAssignmentHandler(echoContext)

		if er != nil {
			return
		}

		assert.Equal(t, service, "Error To Delete Assignment By Id")
		assert.Equal(t, 400, w.Result().StatusCode)
		assert.Error(t, err)
	})

	t.Run("ServerError", func(t *testing.T) {
		assignmentService.On("DeleteAssignmentService", mock.Anything, mock.Anything).Return("Error To Delete Assignment By Id", errors.New("Error To Delete Assignment By Id")).Once()
		service, err := assignmentService.DeleteAssignmentService(int(assignmentData.ID), assignmentData)

		svc.On("DeleteAssignmentService", mock.Anything, mock.Anything).Return("Error To Delete Assignment By Id", errors.New("Error To Delete Assignment By Id")).Once()
		e := echo.New()
		r := httptest.NewRequest("DELETE", "/teacher/assignment/:id", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		er := assignmentController.DeleteAssignmentHandler(echoContext)

		if er != nil {
			return
		}

		assert.Equal(t, service, "Error To Delete Assignment By Id")
		assert.Equal(t, 500, w.Result().StatusCode)
		assert.Error(t, err)
	})
}
