package answare

import (
	"Tugas-Mini-Project/domains"
	"Tugas-Mini-Project/entities"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type AnswareHandler struct {
	svc domains.AnswareService
}

func (h *AnswareHandler) CreateAnswareHandler(c echo.Context) error {
	ans := entities.Answare{}

	e := c.Bind(&ans)

	if e != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"Error": "Error To Create Assignment",
			"Data":  e.Error(),
		})
	}

	answered, err := h.svc.CreateAnswareService(ans)

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

func (h *AnswareHandler) GetAnswareByIdHandler(c echo.Context) error {
	answerers := entities.Answare{}
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"Message": "Error to Get the Assignment",
			"Error":   err.Error(),
		})
	}

	ans, e := h.svc.GetAnswareByIdService(id, answerers)

	if e != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"Message": "Error to get the Assignment By ID",
			"Error":   e.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"Message": "Success Get Assignment By ID",
		"Data":    ans,
	})
}

func (h *AnswareHandler) GetAllAnswareHandler(c echo.Context) error {
	ans, err := h.svc.GetAllAnswareService()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"Status":  "error",
			"Message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"Message": "Success Get All Assignment",
		"Data":    ans,
	})
}

func (h *AnswareHandler) UpdateAnswareHandler(c echo.Context) error {
	var ans entities.Answare
	id, _ := strconv.Atoi(c.Param("id"))

	if err := c.Bind(&ans); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"Message": "Error to Update the Assignment",
			"Error":   err.Error(),
		})
	}

	result, e := h.svc.UpdateAnswareService(id, ans)

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

func (h *AnswareHandler) DeleteAnswareHandler(c echo.Context) error {
	var ans entities.Answare
	id, _ := strconv.Atoi(c.Param("id"))

	if err := c.Bind(&ans); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"Status": http.StatusInternalServerError,
			"Error":  "Error To Delete the Assignment",
			"Data":   err.Error(),
		})
	}

	result, er := h.svc.DeleteAnswareService(id, ans)

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
		"Data":    result,
	})
}
