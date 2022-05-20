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

// CreateDiscussions godoc
// @Summary Create Discussions
// @Description User can Create Discussions
// @Tags Discussions
// @accept json
// @Produce json
// @Router /discussions [post]
// @param data body entities.DiscussionsResponse true "required"
// @Success 200 {object} entities.Discussions
// @Failure 400 {object} entities.Discussions
// @Failure 500 {object} entities.Discussions
func (h *DiscussionHandler) CreateDiscussions(c echo.Context) error {
	discuss := entities.Discussions{}

	e := c.Bind(&discuss)

	if e != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"Error": "error to create assignment",
			"Data":  e.Error(),
		})
	}

	discussions := h.Svc.CreateDiscussionsService(discuss)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"Message": "success create discussions",
		"Data":    discussions,
	})
}

// GetAllDiscussions godoc
// @Summary Get All Discussions
// @Description User can Get All Discussions
// @Tags Discussions
// @accept json
// @Produce json
// @Router /discussions [get]
// @Success 200 {object} entities.Discussions
// @Failure 400 {object} entities.Discussions
func (h *DiscussionHandler) GetAllDiscussions(c echo.Context) error {
	discuss, err := h.Svc.GetAllDiscussionsService()
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"Message": "error to get all discussions",
			"Data":    err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"Message": "success get all discussions",
		"Data":    discuss,
	})
}
