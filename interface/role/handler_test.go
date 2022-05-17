package role_test

import (
	"Tugas-Mini-Project/interface/role"
	role2 "Tugas-Mini-Project/testing/service_test/role"
	"errors"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"io"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestCreateRoleHandler(t *testing.T) {
	svc := role2.RoleService{}

	roleController := role.RoleHandler{
		Svc: &svc,
	}

	t.Run("Success", func(t *testing.T) {
		svc.On("CreateRoleService", mock.Anything).Return(nil).Once()
		e := echo.New()
		r := httptest.NewRequest("POST", "/role", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		err := roleController.CreateRole(echoContext)
		if err != nil {
			return
		}

		assert.Equal(t, 200, w.Result().StatusCode)
	})

	t.Run("NotFound", func(t *testing.T) {
		svc.On("CreateRoleService", mock.Anything).Return(errors.New("Error to Make Unit Testing")).Once()
		e := echo.New()
		r := httptest.NewRequest("POST", "/role", io.Reader(strings.NewReader(`{"Status" : "Not Found"}`)))
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		err := roleController.CreateRole(echoContext)
		if err != nil {
			return
		}

		assert.Equal(t, 404, w.Result().StatusCode)
	})

	t.Run("BadRequest", func(t *testing.T) {
		svc.On("CreateRoleService", mock.Anything).Return(errors.New("Error To Make Unit Testing")).Once()
		e := echo.New()
		r := httptest.NewRequest("POST", "/role", nil)
		w := httptest.NewRecorder()

		echoContext := e.NewContext(r, w)
		err := roleController.CreateRole(echoContext)
		if err != nil {
			return
		}
		assert.Equal(t, 400, w.Result().StatusCode)
	})

}
