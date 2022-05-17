package assignment_test

import (
	"Tugas-Mini-Project/entities"
	"Tugas-Mini-Project/interface/assignment"
	assignment2 "Tugas-Mini-Project/testing/service_test/assignment"
	"errors"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"io"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestCreateAssignmentHandler(t *testing.T) {
	svc := assignment2.AssignmentService{}

	assignmentController := assignment.AssignmentHandler{
		Svc: &svc,
	}

	t.Run("Success", func(t *testing.T) {
		svc.On("CreateAssignmentService", mock.Anything).Return("Success To Create Assignment", nil).Once()

		e := echo.New()
		r := httptest.NewRequest("POST", "/teacher/assignment", nil)
		w := httptest.NewRecorder()

		echoContext := e.NewContext(r, w)
		err := assignmentController.CreateAssignment(echoContext)

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

		err := assignmentController.CreateAssignment(echoContext)

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

		err := assignmentController.CreateAssignment(echoContext)

		if err != nil {
			return
		}

		assert.Equal(t, 500, w.Result().StatusCode)
	})
}

func TestGetAllAssignmentHandler(t *testing.T) {
	svc := assignment2.AssignmentService{}

	assignmentService := new(assignment2.AssignmentService)
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

		er := assignmentController.GetAllAssignment(echoContext)

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

		er := assignmentController.GetAllAssignment(echoContext)

		if er != nil {
			return
		}

		assert.Equal(t, service, assignmentData)
		assert.Equal(t, 500, w.Result().StatusCode)
		assert.Error(t, err)
	})
}

func TestGetAssignmentByIdHandler(t *testing.T) {
	svc := assignment2.AssignmentService{}

	assignmentService := new(assignment2.AssignmentService)

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
		assignmentService.On("GetAssignmentByIdService", mock.Anything, mock.Anything).Return(assignmentData, nil).Once()
		service, err := assignmentService.GetAssignmentByIdService(int(assignmentData.ID))

		svc.On("GetAssignmentByIdService", mock.Anything, mock.Anything).Return(assignmentData, nil).Once()
		e := echo.New()
		r := httptest.NewRequest("GET", "/assignment/:id", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		er := assignmentController.GetAssignmentById(echoContext)

		if er != nil {
			return
		}

		assert.Equal(t, service, assignmentData)
		assert.Equal(t, 200, w.Result().StatusCode)
		assert.NoError(t, err)
	})

	t.Run("BadRequest", func(t *testing.T) {
		assignmentService.On("GetAssignmentByIdService", mock.Anything, mock.Anything).Return(assignmentData, errors.New("Error To Get Assignment By Id")).Once()
		service, err := assignmentService.GetAssignmentByIdService(int(assignmentData.ID))

		svc.On("GetAssignmentByIdService", mock.Anything, mock.Anything).Return(assignmentData, errors.New("Error To Get Assignment By Id")).Once()
		e := echo.New()
		r := httptest.NewRequest("GET", "/assignment/:id", io.Reader(strings.NewReader(`{"Status" : "Bad Request"}`)))
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		er := assignmentController.GetAssignmentById(echoContext)

		if er != nil {
			return
		}

		assert.Equal(t, service, assignmentData)
		assert.Equal(t, 400, w.Result().StatusCode)
		assert.Error(t, err)
	})
}

func TestUpdateAssignmentHandler(t *testing.T) {
	svc := assignment2.AssignmentService{}

	assignmentService := new(assignment2.AssignmentService)
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
		assignmentService.On("UpdateAssignmentService", mock.Anything, mock.Anything).Return(assignmentData, nil).Once()
		service, err := assignmentService.UpdateAssignmentService(int(assignmentData.ID), assignmentData)

		svc.On("UpdateAssignmentService", mock.Anything, mock.Anything).Return(assignmentData, nil).Once()
		e := echo.New()
		r := httptest.NewRequest("PUT", "/teacher/assignment/:id", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		er := assignmentController.UpdateAssignment(echoContext)

		if er != nil {
			return
		}

		assert.Equal(t, service, assignmentData)
		assert.Equal(t, 200, w.Result().StatusCode)
		assert.NoError(t, err)
	})

	t.Run("BadRequest", func(t *testing.T) {
		assignmentService.On("UpdateAssignmentService", mock.Anything, mock.Anything).Return(assignmentData, errors.New("Error To Update Assignment By Id")).Once()
		service, err := assignmentService.UpdateAssignmentService(int(assignmentData.ID), assignmentData)

		svc.On("UpdateAssignmentService", mock.Anything, mock.Anything).Return(assignmentData, errors.New("Error To Update Assignment By Id")).Once()
		e := echo.New()
		r := httptest.NewRequest("PUT", "/teacher/assignment/:id", io.Reader(strings.NewReader(`{"Status" : "Bad Request"}`)))
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		er := assignmentController.UpdateAssignment(echoContext)

		if er != nil {
			return
		}

		assert.Equal(t, service, assignmentData)
		assert.Equal(t, 400, w.Result().StatusCode)
		assert.Error(t, err)
	})

	t.Run("ServerError", func(t *testing.T) {
		assignmentService.On("UpdateAssignmentService", mock.Anything, mock.Anything).Return(assignmentData, errors.New("Error To Update Assignment By Id")).Once()
		service, err := assignmentService.UpdateAssignmentService(int(assignmentData.ID), assignmentData)

		svc.On("UpdateAssignmentService", mock.Anything, mock.Anything).Return(assignmentData, errors.New("Error To Update Assignment By Id")).Once()
		e := echo.New()
		r := httptest.NewRequest("PUT", "/teacher/assignment/:id", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		er := assignmentController.UpdateAssignment(echoContext)

		if er != nil {
			return
		}

		assert.Equal(t, service, assignmentData)
		assert.Equal(t, 500, w.Result().StatusCode)
		assert.Error(t, err)
	})
}

func TestDeleteAssignmentHandler(t *testing.T) {
	svc := assignment2.AssignmentService{}

	assignmentService := new(assignment2.AssignmentService)
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

		er := assignmentController.DeleteAssignment(echoContext)

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

		er := assignmentController.DeleteAssignment(echoContext)

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

		er := assignmentController.DeleteAssignment(echoContext)

		if er != nil {
			return
		}

		assert.Equal(t, service, "Error To Delete Assignment By Id")
		assert.Equal(t, 500, w.Result().StatusCode)
		assert.Error(t, err)
	})
}
