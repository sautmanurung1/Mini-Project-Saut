package testing

import (
	"Tugas-Mini-Project/domains"
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

func TestRegister(t *testing.T) {
	userRegister := new(mocks.AuthRepository)
	userRegisterData := entities.User{
		Name:     "Name Testing",
		Username: "Username Testing",
		Password: "Password Testing",
		RoleId:   1,
	}

	t.Run("Success", func(t *testing.T) {
		userRegister.On("Register", mock.Anything).Return(nil).Once()

		userRegister := domains.AuthRepository(userRegister)
		err := userRegister.Register(userRegisterData)
		assert.NoError(t, err)
	})

	t.Run("Failed", func(t *testing.T) {
		userRegister.On("Register", mock.Anything).Return(errors.New("Error to Make Unit Testing")).Once()

		userRegister := domains.AuthRepository(userRegister)
		err := userRegister.Register(userRegisterData)
		assert.Error(t, err)
	})
}

func TestLogin(t *testing.T) {
	userLogin := new(mocks.AuthRepository)
	userLoginData := entities.User{
		Username: "Username Testing",
		Password: "Password Testing",
		RoleId:   1,
	}

	t.Run("Success", func(t *testing.T) {
		userLogin.On("Login", mock.Anything, mock.Anything, mock.Anything).Return(userLoginData, nil).Once()
		userLogin := domains.AuthRepository(userLogin)
		login, err := userLogin.Login(userLoginData.Username, userLoginData.Password, userLoginData.RoleId)
		assert.Equal(t, login, userLoginData)
		assert.NoError(t, err)
	})

	t.Run("Failed", func(t *testing.T) {
		userLogin.On("Login", mock.Anything, mock.Anything, mock.Anything).Return(userLoginData, errors.New("Error to Make Unit Testing")).Once()
		userLogin := domains.AuthRepository(userLogin)
		login, err := userLogin.Login(userLoginData.Username, userLoginData.Password, userLoginData.RoleId)
		assert.Equal(t, login, userLoginData)
		assert.Error(t, err)
	})
}

func TestRegisterService(t *testing.T) {
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

		err := authController.RegisterHandler(echoContext)
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
		err := authController.RegisterHandler(echoContext)
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
		err := authController.RegisterHandler(echoContext)
		if err != nil {
			return
		}

		assert.Equal(t, 401, w.Result().StatusCode)
	})
}

func TestLoginService(t *testing.T) {
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

		err := authController.LoginHandler(echoContext)
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

		err := authController.LoginHandler(echoContext)
		if err != nil {
			return
		}

		assert.Equal(t, 400, w.Result().StatusCode)
	})
}
