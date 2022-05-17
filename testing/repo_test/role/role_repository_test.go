package role_test

import (
	"Tugas-Mini-Project/domains"
	"Tugas-Mini-Project/entities"
	"Tugas-Mini-Project/testing/repo_test/role"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestCreateRoleRepository(t *testing.T) {
	roleRepository := new(role.RoleRepository)
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
