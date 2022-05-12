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

func (h *QuestionHandler) CreateQuestionHandler(c echo.Context) error {
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

func (h *QuestionHandler) GetQuestionByIdHandler(c echo.Context) error {
	question := entities.Question{}
	id, _ := strconv.Atoi(c.Param("id"))

	questions, er := h.Svc.GetQuestionByIDService(id, question)

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

func (h *QuestionHandler) GetAllQuestionsHandler(c echo.Context) error {
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

func (h *QuestionHandler) UpdateQuestionHandler(c echo.Context) error {
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

func (h *QuestionHandler) DeleteQuestionHandler(c echo.Context) error {
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
