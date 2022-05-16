package auth

import (
	"Tugas-Mini-Project/domains"
	"Tugas-Mini-Project/domains/mocks"
	"Tugas-Mini-Project/entities"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
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
