package discussions

import (
	"Tugas-Mini-Project/domains"
	"github.com/labstack/echo/v4"
	"net/http"
)

type DiscussionHandler struct {
	svc domains.DiscussionsService
}

func (h *DiscussionHandler) GetAllDiscussions(c echo.Context) error {
	discussions, err := h.svc.GetAllDiscussionsService()
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
		"Data":    discussions,
	})
}
