package discussions

import (
	"Tugas-Mini-Project/domains"
	"Tugas-Mini-Project/entities"
	"github.com/labstack/echo/v4"
	"net/http"
)

type DiscussionHandler struct {
	Svc domains.DiscussionsService
}

func (h *DiscussionHandler) CreateDiscsussions(c echo.Context) error {
	discuss := entities.Discussions{}

	e := c.Bind(&discuss)

	if e != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"Error": "Error To Create Assignment",
			"Data":  e.Error(),
		})
	}

	discussions, err := h.Svc.CreateDiscussionsService(discuss)

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
		"Data":    discussions,
	})
}

func (h *DiscussionHandler) GetAllDiscussions(c echo.Context) error {
	discuss, err := h.Svc.GetAllDiscussionsService()
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"Status":  http.StatusBadRequest,
			"Message": "Error TO Get All Discussions",
			"Data":    err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"Status":  http.StatusOK,
		"Message": "Success Get All Discussions",
		"Data":    discuss,
	})
}
