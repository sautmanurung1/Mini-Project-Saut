package assignment

import (
	"Tugas-Mini-Project/domains"
	"Tugas-Mini-Project/entities"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type AssignmentHandler struct {
	Svc domains.AssignmentService
}

func (h *AssignmentHandler) CreateAssignmentHandler(c echo.Context) error {
	assign := entities.Assignment{}

	e := c.Bind(&assign)

	if e != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"Error": "Error To Create Assignment",
			"Data":  e.Error(),
		})
	}

	result, err := h.Svc.CreateAssignmentService(assign)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"Status": http.StatusInternalServerError,
			"Error":  "Error",
			"Data":   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"Status":  http.StatusOK,
		"Message": "Success Create Assignment",
		"Data":    result,
	})
}

func (h *AssignmentHandler) GetAllAssignmentHandler(c echo.Context) error {
	assignments, err := h.Svc.GetAllAssignmentService()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"Status":  "error",
			"Message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"Message": "Success Get All Assignment",
		"Data":    assignments,
	})
}

func (h *AssignmentHandler) GetAssignmentByIdHandler(c echo.Context) error {
	var assignment entities.Assignment
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"Message": "Error to Get the Assignment",
			"Error":   err.Error(),
		})
	}

	assign, e := h.Svc.GetAssignmentByIdService(id, assignment)

	if e != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"Message": "Error to get the Assignment By ID",
			"Error":   e.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"Message": "Success Get Assignment By ID",
		"Data":    assign,
	})
}

func (h *AssignmentHandler) UpdateAssignmentHandler(c echo.Context) error {
	var assignments entities.Assignment
	id, _ := strconv.Atoi(c.Param("id"))

	if err := c.Bind(&assignments); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"Message": "Error to Update the Assignment",
			"Error":   err.Error(),
		})
	}

	result, e := h.Svc.UpdateAssignmentService(id, assignments)

	if e != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"Message": "Error to Update the Assignment",
			"Error":   e.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"Status":  http.StatusOK,
		"Message": "Success Update Assignment",
		"Data":    result,
	})
}

func (h *AssignmentHandler) DeleteAssignmentHandler(c echo.Context) error {
	var assign entities.Assignment
	id, _ := strconv.Atoi(c.Param("id"))

	if err := c.Bind(&assign); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"Status": http.StatusInternalServerError,
			"Error":  "Error To Delete the Assignment",
			"Data":   err.Error(),
		})
	}

	assignments, er := h.Svc.DeleteAssignmentService(id, assign)

	if er != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"Status": http.StatusInternalServerError,
			"Error":  "Error To Delete the Assignment",
			"Data":   er.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"Status":  http.StatusOK,
		"Message": "Success To Delete The Assignment",
		"Data":    assignments,
	})
}
