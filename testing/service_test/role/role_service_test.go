package role

import (
	"Tugas-Mini-Project/domains"
	"Tugas-Mini-Project/entities"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestCreateRoleService(t *testing.T) {
	roleService := new(RoleService)
	roleData := entities.Role{
		ID:   1,
		Name: "Name Testing",
	}

	t.Run("Success", func(t *testing.T) {
		roleService.On("CreateRoleService", mock.Anything).Return(nil).Once()

		roleService := domains.RoleService(roleService)
		err := roleService.CreateRoleService(roleData)
		assert.NoError(t, err)
	})

	t.Run("Failed", func(t *testing.T) {
		roleService.On("CreateRoleService", mock.Anything).Return(errors.New("Error to Make Unit Testing")).Once()

		roleService := domains.RoleService(roleService)
		err := roleService.CreateRoleService(roleData)
		assert.Error(t, err)
	})
}
