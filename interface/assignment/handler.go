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
// @Router /teacher/assignment [post]
// @param data body entities.AssignmentResponse true "required"
// @Success 200 {object} entities.Assignment
// @Failure 400 {object} entities.Assignment
// @Failure 500 {object} entities.Answare
// @Security JWT
func (h *AssignmentHandler) CreateAssignment(c echo.Context) error {
	assign := entities.Assignment{}

	e := c.Bind(&assign)

	if e != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"Error": "error to create assignment",
			"Data":  e.Error(),
		})
	}

	result, err := h.Svc.CreateAssignmentService(assign)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"Error": "error to create assignment",
			"Data":  err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"Message": "success to create assignment",
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
func (h *AssignmentHandler) GetAllAssignment(c echo.Context) error {
	assignments, err := h.Svc.GetAllAssignmentService()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"Status":  "error to get all assignment",
			"Message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"Message": "success get all assignment",
		"Data":    assignments,
	})
}

// GetAssignmentById godoc
// @Summary Get Assignment By Id
// @Description User Can Get Answare By ID
// @Tags Assignment
// @accept json
// @Produce json
// @Router /assignment/{id} [get]
// @param id path int true "id"
// @Success 200 {object} entities.Assignment
// @Failure 400 {object} entities.Assignment
func (h *AssignmentHandler) GetAssignmentById(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	assign, e := h.Svc.GetAssignmentByIdService(id)

	if e != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"Message": "error to get the Assignment by id",
			"Error":   e.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"Message": "success get assignment by id",
		"Data":    assign,
	})
}

// UpdateAssignment godoc
// @Summary Update Assignment Teacher
// @Description Teacher Can Update Assignment
// @Tags Assignment
// @accept json
// @Produce json
// @Router /teacher/assignment/{id} [put]
// @param id path int true "id"
// @param data body entities.AssignmentResponse true "required"
// @Success 200 {object} entities.Assignment
// @Failure 400 {object} entities.Assignment
// @Failure 500 {object} entities.Answare
// @Security JWT
func (h *AssignmentHandler) UpdateAssignment(c echo.Context) error {
	var assignments entities.Assignment
	id, _ := strconv.Atoi(c.Param("id"))

	if err := c.Bind(&assignments); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"Message": "error to update the assignment",
			"Error":   err.Error(),
		})
	}

	result, e := h.Svc.UpdateAssignmentService(id, assignments)

	if e != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"Message": "error to update the assignment",
			"Error":   e.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"Message": "success to update assignment",
		"Data":    result,
	})
}

// DeleteAssignment godoc
// @Summary Delete Assignment Teacher
// @Description Teacher Can Delete Assignment
// @Tags Assignment
// @accept json
// @Produce json
// @Router /teacher/assignment/{id} [delete]
// @param id path int true "id"
// @Success 200 {object} entities.Assignment
// @Failure 400 {object} entities.Assignment
// @Failure 500 {object} entities.Answare
// @Security JWT
func (h *AssignmentHandler) DeleteAssignment(c echo.Context) error {
	var assign entities.Assignment
	id, _ := strconv.Atoi(c.Param("id"))

	if err := c.Bind(&assign); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"Error": "error to delete the assignment",
			"Data":  err.Error(),
		})
	}

	assignments, er := h.Svc.DeleteAssignmentService(id, assign)

	if er != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"Error": "error To delete the assignment",
			"Data":  er.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"Message": "success to delete the assignment",
		"Data":    assignments,
	})
}
