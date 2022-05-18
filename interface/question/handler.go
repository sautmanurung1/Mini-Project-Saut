package question

import (
	"Tugas-Mini-Project/domains"
	"Tugas-Mini-Project/entities"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type QuestionHandler struct {
	Svc domains.QuestionService
}

// CreateQuestion godoc
// @Summary Teacher Create Question
// @Description Teacher Can Create Question
// @Tags Question
// @accept json
// @Produce json
// @Router /teacher/questions [post]
// @param data body entities.QuestionResponse true "required"
// @Success 200 {object} entities.Question
// @Failure 400 {object} entities.Question
// @Failure 500 {object} entities.Question
// @Security JWT
func (h *QuestionHandler) CreateQuestion(c echo.Context) error {
	questions := entities.Question{}

	e := c.Bind(&questions)

	if e != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  http.StatusBadRequest,
			"message": e.Error(),
		})
	}

	question, err := h.Svc.CreateQuestionService(questions)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  http.StatusInternalServerError,
			"Error":   "Error to Create Question",
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"status":  http.StatusCreated,
		"message": "Success Create Question",
		"Data":    question,
	})
}

// GetQuestionById godoc
// @Summary Get Question By Id
// @Description User Can Get Question By Id
// @Tags Question
// @accept json
// @Produce json
// @Router /questions/{id} [get]
// @param id path int true "id"
// @Success 200 {object} entities.Question
// @Failure 400 {object} entities.Question
func (h *QuestionHandler) GetQuestionById(c echo.Context) error {
	// question := entities.Question{}
	id, _ := strconv.Atoi(c.Param("id"))

	questions, er := h.Svc.GetQuestionByIDService(id)
	if er != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"Status":  http.StatusBadRequest,
			"Message": "Error To Get The Questions",
			"Error":   er.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"Status":  http.StatusOK,
		"Message": "Success Get Question By ID",
		"Data":    questions,
	})
}

// GetAllQuestions godoc
// @Summary Get All Question
// @Description Get ALl Question
// @Tags Question
// @accept json
// @Produce json
// @Router /questions [get]
// @Success 200 {object} entities.Question
// Failure 400 {object} entities.Question
func (h *QuestionHandler) GetAllQuestions(c echo.Context) error {
	questions, err := h.Svc.GetAllQuestionService()
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"Status":  http.StatusBadRequest,
			"Message": "Error TO Get All Questions",
			"Data":    err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"Status":  http.StatusOK,
		"Message": "Success Get All Questions",
		"Data":    questions,
	})
}

// UpdateQuestion godoc
// @Summary Update Question Teacher
// @Description Teacher Can Update Them Question
// @Tags Question
// @accept json
// @Produce json
// @Router /teacher/question/{id} [put]
// @param id path int true "id"
// @param data body entities.QuestionResponse true "required"
// @Success 200 {object} entities.Question
// @Failure 400 {object} entities.Question
// @Failure 500 {object} entities.Question
// @Security JWT
func (h *QuestionHandler) UpdateQuestion(c echo.Context) error {
	var questions entities.Question
	id, _ := strconv.Atoi(c.Param("id"))

	if err := c.Bind(&questions); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"Status":  http.StatusBadRequest,
			"Message": "Error",
			"Error":   err.Error(),
		})
	}

	result, er := h.Svc.UpdateQuestionService(id, questions)

	if er != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"Status":  http.StatusInternalServerError,
			"Message": "Error To Update The Questions",
			"Data":    er.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"Status":  http.StatusOK,
		"Message": "Success Update Questions",
		"Data":    result,
	})
}

// DeleteQuestion godoc
// @Summary Delete Question Teacher
// @Description Teacher Can delete Them Question
// @Tags Question
// @accept json
// @Produce json
// @Router /teacher/question/{id} [delete]
// @param id path int true "id"
// @Success 200 {object} entities.Question
// @Failure 400 {object} entities.Question
// @Failure 500 {object} entities.Question
// @Security JWT
func (h *QuestionHandler) DeleteQuestion(c echo.Context) error {
	var questions entities.Question
	id, _ := strconv.Atoi(c.Param("id"))

	result, er := h.Svc.DeleteQuestionService(id, questions)

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
