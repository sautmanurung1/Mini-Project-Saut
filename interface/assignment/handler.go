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

// CreateAssignment godoc
// @Summary Create Assignment Teacher
// @Description Teacher Can Create Assignment
// @Tags Assignment
// @accept json
// @Produce json
// @Router /teacher/answare [post]
// @param data body entities.AssignmentResponse true "required"
// @Success 200 {object} entities.Assignment
// @Failure 400 {object} entities.Assignment
// @Failute 500 {object} entities.Answare
// @Security JWT
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

// GetAllAssignment godoc
// @Summary Get All Assignment
// @Description User Can Get Assignment
// @Tags Assignment
// @accept json
// @Produce json
// @Router /assignment [get]
// @Success 200 {object} entities.Assignment
// @Failure 500 {object} entities.Assignment
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

// GetAssignmentByID godoc
// @Summary Get Assignment By Id
// @Description User Can Get Answare By ID
// @Tags Assignment
// @accept json
// @Produce json
// @Router /answare/{id} [get]
// @param id path entities.Assignment true "id"
// @Success 200 {object} entities.Assignment
// @Failure 400 {object} entities.Assignment
// @Security JWT
func (h *AssignmentHandler) GetAssignmentByIdHandler(c echo.Context) error {
	var assignment entities.Assignment
	id, _ := strconv.Atoi(c.Param("id"))

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

// UpdateAssignment godoc
// @Summary Update Assignment Teacher
// @Description Teacher Can Update Assignment
// @Tags Assignment
// @accept json
// @Produce json
// @Router /teacher/answare/{id} [put]
// @param id path entities.Assignment true "id"
// @Success 200 {object} entities.Assignment
// @Failure 400 {object} entities.Assignment
// @Failute 500 {object} entities.Answare
// @Security JWT
func (h *AssignmentHandler) UpdateAssignmentHandler(c echo.Context) error {
	var assignments entities.Assignment
	id, _ := strconv.Atoi(c.Param("id"))

	if err := c.Bind(&assignments); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
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

// DeleteAssignment godoc
// @Summary Delete Assignment Teacher
// @Description Teacher Can Delete Assignment
// @Tags Assignment
// @accept json
// @Produce json
// @Router /teacher/answare/{id} [delete]
// @param id path entities.Assignment true "id"
// @Success 200 {object} entities.Assignment
// @Failure 400 {object} entities.Assignment
// @Failute 500 {object} entities.Answare
// @Security JWT
func (h *AssignmentHandler) DeleteAssignmentHandler(c echo.Context) error {
	var assign entities.Assignment
	id, _ := strconv.Atoi(c.Param("id"))

	if err := c.Bind(&assign); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"Status": http.StatusBadRequest,
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
