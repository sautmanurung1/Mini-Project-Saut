package mocks

import (
	"Tugas-Mini-Project/domains"
	"Tugas-Mini-Project/entities"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"testing"
)

func TestRegisterRepository(t *testing.T) {
	userRegister := new(AuthRepository)
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

func TestLoginRepository(t *testing.T) {
	userLogin := new(AuthRepository)
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
	userRegister := new(AuthService)
	userRegisterData := entities.User{
		Name:     "Name Testing",
		Username: "Username Testing",
		Password: "Password Testing",
		RoleId:   1,
	}

	t.Run("Success", func(t *testing.T) {
		userRegister.On("RegisterService", mock.Anything).Return(nil).Once()

		userRegister := domains.AuthService(userRegister)
		err := userRegister.RegisterService(userRegisterData)
		assert.NoError(t, err)
	})

	t.Run("Failed", func(t *testing.T) {
		userRegister.On("RegisterService", mock.Anything).Return(errors.New("Error to Make Unit Testing")).Once()

		userRegister := domains.AuthService(userRegister)
		err := userRegister.RegisterService(userRegisterData)
		assert.Error(t, err)
	})
}

func TestLoginService(t *testing.T) {
	userLogin := new(AuthService)
	userLoginData := entities.User{
		Username: "Username Testing",
		Password: "Password Testing",
		RoleId:   1,
	}

	t.Run("Success", func(t *testing.T) {
		userLogin.On("LoginService", mock.Anything, mock.Anything, mock.Anything).Return("Success Login", http.StatusOK).Once()
		userLogin := domains.AuthService(userLogin)
		login, _ := userLogin.LoginService(userLoginData.Username, userLoginData.Password, userLoginData.RoleId)
		assert.Equal(t, login, "Success Login")
	})

	t.Run("Failed", func(t *testing.T) {
		userLogin.On("LoginService", mock.Anything, mock.Anything, mock.Anything).Return("Bad Request", http.StatusBadRequest).Once()
		userLogin := domains.AuthService(userLogin)
		login, _ := userLogin.LoginService(userLoginData.Username, userLoginData.Password, userLoginData.RoleId)
		assert.Equal(t, login, "Bad Request")
	})
}
