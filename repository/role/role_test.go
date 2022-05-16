package role_test

import (
	"Tugas-Mini-Project/domains"
	"Tugas-Mini-Project/domains/mocks"
	"Tugas-Mini-Project/entities"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestCreateRole(t *testing.T) {
	roleRepository := new(mocks.RoleRepository)
	roleData := entities.Role{
		ID:   1,
		Name: "Name Testing",
	}

	t.Run("Success", func(t *testing.T) {
		roleRepository.On("CreateRole", mock.Anything).Return(nil).Once()

		roleRepository := domains.RoleRepository(roleRepository)
		err := roleRepository.CreateRole(roleData)
		assert.NoError(t, err)
	})

	t.Run("Failed", func(t *testing.T) {
		roleRepository.On("CreateRole", mock.Anything).Return(errors.New("Error to Make Unit Testing")).Once()

		roleRepository := domains.RoleRepository(roleRepository)
		err := roleRepository.CreateRole(roleData)
		assert.Error(t, err)
	})
}
