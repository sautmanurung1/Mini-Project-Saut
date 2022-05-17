package auth_test

import (
	"Tugas-Mini-Project/domains/mocks"
	"Tugas-Mini-Project/entities"
	"Tugas-Mini-Project/interface/auth"
	"errors"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestRegisterHandler(t *testing.T) {
	svc := mocks.AuthService{}

	authController := auth.AuthHandler{
		Svc: &svc,
	}

	t.Run("Success", func(t *testing.T) {
		svc.On("RegisterService", mock.Anything).Return(nil).Once()
		e := echo.New()
		r := httptest.NewRequest("POST", "/register", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		err := authController.Register(echoContext)
		if err != nil {
			return
		}

		assert.Equal(t, 200, w.Result().StatusCode)
	})

	t.Run("ServerError", func(t *testing.T) {
		svc.On("RegisterService", mock.Anything).Return(errors.New("Error to Make Unit Testing")).Once()
		e := echo.New()
		r := httptest.NewRequest("POST", "/register", io.Reader(strings.NewReader(`{"Status" : "Internal Server Error"}`)))
		w := httptest.NewRecorder()

		echoContext := e.NewContext(r, w)
		err := authController.Register(echoContext)
		if err != nil {
			return
		}

		assert.Equal(t, 500, w.Result().StatusCode)
	})

	t.Run("StatusUnauthorized", func(t *testing.T) {
		svc.On("RegisterService", mock.Anything).Return(errors.New("Error to Make Unit Testing")).Once()
		e := echo.New()
		r := httptest.NewRequest("POST", "/register", nil)
		w := httptest.NewRecorder()

		echoContext := e.NewContext(r, w)
		err := authController.Register(echoContext)
		if err != nil {
			return
		}

		assert.Equal(t, 401, w.Result().StatusCode)
	})
}

func TestLoginHandler(t *testing.T) {
	svc := mocks.AuthService{}
	userLogin := new(mocks.AuthService)
	userLoginData := entities.User{
		Username: "Username Testing",
		Password: "Password Testing",
		RoleId:   1,
	}
	authController := auth.AuthHandler{
		Svc: &svc,
	}

	t.Run("Success", func(t *testing.T) {
		userLogin.On("LoginService", mock.Anything, mock.Anything, mock.Anything).Return("Success Login", http.StatusOK).Once()
		userLogin.LoginService(userLoginData.Username, userLoginData.Password, userLoginData.RoleId)
		svc.On("LoginService", mock.Anything, mock.Anything, mock.Anything).Return("Success Login", http.StatusOK).Once()
		e := echo.New()
		r := httptest.NewRequest("POST", "/login", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		err := authController.Login(echoContext)
		if err != nil {
			return
		}

		assert.Equal(t, 200, w.Result().StatusCode)
	})

	t.Run("BadRequest", func(t *testing.T) {
		userLogin.On("LoginService", mock.Anything, mock.Anything, mock.Anything).Return("Bad Request", http.StatusBadRequest).Once()
		userLogin.LoginService(userLoginData.Username, userLoginData.Password, userLoginData.RoleId)
		svc.On("LoginService", mock.Anything, mock.Anything, mock.Anything).Return("Bad Request", http.StatusBadRequest).Once()
		e := echo.New()
		r := httptest.NewRequest("POST", "/login", io.Reader(strings.NewReader(`{"Status" : "Bad Request"}`)))
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		err := authController.Login(echoContext)
		if err != nil {
			return
		}

		assert.Equal(t, 400, w.Result().StatusCode)
	})
}
