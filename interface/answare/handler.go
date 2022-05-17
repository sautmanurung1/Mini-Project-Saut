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
			"Error": "Error To Create Assignment",
			"Data":  e.Error(),
		})
	}

	answered, err := h.Svc.CreateAnswareService(ans)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"Status": http.StatusInternalServerError,
			"Error":  "Error",
			"Data":   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"Status":  http.StatusOK,
		"Message": "Success Create Answare",
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
			"Message": "Error to get the Answare By ID",
			"Error":   e.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"Message": "Success Get Answare By ID",
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
			"Status":  "error",
			"Message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"Message": "Success Get All Answare",
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
			"Message": "Error to Update the Answare",
			"Error":   err.Error(),
		})
	}

	result, e := h.Svc.UpdateAnswareService(id, ans)

	if e != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"Message": "Error to Update the Answare",
			"Error":   e.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"Status":  http.StatusOK,
		"Message": "Success Update Answare",
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
// @param data body entities.AnswareResponse true "required"
// @Success 200 {object} entities.Answare
// @Failure 400 {object} entities.Answare
// @Failure 500 {object} entities.Answare
// @Security JWT
func (h *AnswareHandler) DeleteAnsware(c echo.Context) error {
	var ans entities.Answare
	id, _ := strconv.Atoi(c.Param("id"))

	if err := c.Bind(&ans); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"Status": http.StatusBadRequest,
			"Error":  "Error To Delete the Answare",
			"Data":   err.Error(),
		})
	}

	result, er := h.Svc.DeleteAnswareService(id, ans)

	if er != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"Status": http.StatusInternalServerError,
			"Error":  "Error To Delete the Answare",
			"Data":   er.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"Status":  http.StatusOK,
		"Message": "Success To Delete The Answare",
		"Data":    result,
	})
}
