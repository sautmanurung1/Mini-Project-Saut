package question

import (
	"Tugas-Mini-Project/domains"
	"Tugas-Mini-Project/entities"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type handler struct {
	repository domains.QuestionRepository
}

func NewQuestionHandler(repository domains.QuestionRepository) domains.QuestionHandler {
	return &handler{
		repository: repository,
	}
}

func (h *handler) CreateQuestionHandler(c echo.Context) error {
	questions := entities.Question{}

	e := c.Bind(&questions)

	if e != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  http.StatusBadRequest,
			"message": e.Error(),
		})
	}

	err := h.repository.CreateQuestion(questions)

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
		"Data":    questions,
	})
}

func (h *handler) GetQuestionByIdHandler(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"Status":  http.StatusNotFound,
			"Message": "Nothing to Get",
			"Error":   err.Error(),
		})
	}

	questions, er := h.repository.GetQuestionByID(id)

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

func (h *handler) GetAllQuestionsHandler(c echo.Context) error {
	questions, err := h.repository.GetAllQuestion()
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

func (h *handler) UpdateQuestionHandler(c echo.Context) error {
	var questions entities.Question
	id, _ := strconv.Atoi(c.Param("id"))

	if err := c.Bind(&questions); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"Status":  http.StatusBadRequest,
			"Message": "Error",
			"Error":   err.Error(),
		})
	}

	result, er := h.repository.UpdateQuestion(id, questions)

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

func (h *handler) DeleteQuestionHandler(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"Status":  http.StatusBadRequest,
			"Message": "Error",
			"Data":    err.Error(),
		})
	}

	questions, er := h.repository.DeleteQuestion(id)

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
		"Data":    questions,
	})
}
