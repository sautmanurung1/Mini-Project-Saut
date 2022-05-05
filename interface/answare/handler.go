package answare

import (
	"Tugas-Mini-Project/domains"
	"Tugas-Mini-Project/entities"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type handler struct {
	repository domains.AnswareRepository
}

func NewAnswareHandler(repository domains.AnswareRepository) domains.AnswareHandler {
	return &handler{
		repository: repository,
	}
}

func (h handler) CreateAnswareHandler(c echo.Context) error {
	ans := entities.Answare{}

	e := c.Bind(&ans)

	if e != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"Error": "Error To Create Assignment",
			"Data":  e.Error(),
		})
	}

	err := h.repository.CreateAnsware(ans)

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
		"Data":    ans,
	})
}

func (h handler) GetAnswareByIdHandler(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"Message": "Error to Get the Assignment",
			"Error":   err.Error(),
		})
	}

	ans, e := h.repository.GetAnswareById(id)

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

func (h handler) GetAllAnswareHandler(c echo.Context) error {
	ans, err := h.repository.GetAllAnsware()
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

func (h handler) UpdateAnswareHandler(c echo.Context) error {
	var ans entities.Answare
	id, _ := strconv.Atoi(c.Param("id"))

	if err := c.Bind(&ans); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"Message": "Error to Update the Assignment",
			"Error":   err.Error(),
		})
	}

	result, e := h.repository.UpdateAnsware(id, ans)

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

func (h handler) DeleteAnswareHandler(c echo.Context) error {
	var ans entities.Answare
	id, _ := strconv.Atoi(c.Param("id"))

	if err := c.Bind(&ans); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"Status": http.StatusInternalServerError,
			"Error":  "Error To Delete the Assignment",
			"Data":   err.Error(),
		})
	}

	result, er := h.repository.DeleteAnsware(id)

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
