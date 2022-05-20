package answare

import (
	"Tugas-Mini-Project/domains"
	"Tugas-Mini-Project/entities"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type AnswareHandler struct {
	Svc domains.AnswareService
}

// CreateAnsware godoc
// @Summary Create Answare Student
// @Description Student Can Create Answare
// @Tags Answare
// @accept json
// @Produce json
// @Router /student/answare [post]
// @param data body entities.AnswareResponse true "required"
// @Success 200 {object} entities.Answare
// @Failure 400 {object} entities.Answare
// @Failure 500 {object} entities.Answare
// @Security JWT
func (h *AnswareHandler) CreateAnsware(c echo.Context) error {
	ans := entities.Answare{}

	e := c.Bind(&ans)

	if e != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"Error": "error to create answare",
			"Data":  e.Error(),
		})
	}

	answered, err := h.Svc.CreateAnswareService(ans)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"Error": "error to create answare",
			"Data":  err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"Message": "success to create answare",
		"Data":    answered,
	})
}

// GetAnswareById godoc
// @Summary Get Answare Student By Id
// @Description Student Can Get Answare By Id
// @Tags Answare
// @accept json
// @Produce json
// @Router /answare/{id} [get]
// @param id path int true "id"
// @Success 200 {object} entities.Answare
// @Failure 400 {object} entities.Answare
func (h *AnswareHandler) GetAnswareById(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	ans, e := h.Svc.GetAnswareByIdService(id)

	if e != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"Message": "error to get answare by id",
			"Error":   e.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"Message": "success to get answare by id",
		"Data":    ans,
	})
}

// GetAllAnsware godoc
// @Summary Get All Answare Student
// @Description Student Can Get Answare
// @Tags Answare
// @accept json
// @Produce json
// @Router /answare [get]
// @Success 200 {object} entities.Answare
// @Failure 500 {object} entities.Answare
func (h *AnswareHandler) GetAllAnsware(c echo.Context) error {
	ans, err := h.Svc.GetAllAnswareService()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"Status":  "error to get all answare",
			"Message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"Message": "success to get all answare",
		"Data":    ans,
	})
}

// UpdateAnsware godoc
// @Summary Update Answare Student
// @Description Student Can Update Them Answare
// @Tags Answare
// @accept json
// @Produce json
// @Router /student/answare/{id} [put]
// @param id path int true "id"
// @param data body entities.AnswareResponse true "required"
// @Success 200 {object} entities.Answare
// @Failure 400 {object} entities.Answare
// @Failure 500 {object} entities.Answare
// @Security JWT
func (h *AnswareHandler) UpdateAnsware(c echo.Context) error {
	var ans entities.Answare
	id, _ := strconv.Atoi(c.Param("id"))

	if err := c.Bind(&ans); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"Message": "error to update the answare",
			"Error":   err.Error(),
		})
	}

	result, e := h.Svc.UpdateAnswareService(id, ans)

	if e != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"Message": "error to update the answare",
			"Error":   e.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"Message": "success update answare",
		"Data":    result,
	})
}

// DeleteAnsware godoc
// @Summary Delete Answare Student
// @Description Student Can Delete Them Answare
// @Tags Answare
// @accept json
// @Produce json
// @Router /student/answare/{id} [delete]
// @param id path int true "id"
// @Success 200 {object} entities.Answare
// @Failure 400 {object} entities.Answare
// @Failure 500 {object} entities.Answare
// @Security JWT
func (h *AnswareHandler) DeleteAnsware(c echo.Context) error {
	var ans entities.Answare
	id, _ := strconv.Atoi(c.Param("id"))

	if err := c.Bind(&ans); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"Error": "error to delete the answare",
			"Data":  err.Error(),
		})
	}

	result, er := h.Svc.DeleteAnswareService(id, ans)

	if er != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"Error": "error to delete the answare",
			"Data":  er.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"Message": "success to delete the answare",
		"Data":    result,
	})
}
