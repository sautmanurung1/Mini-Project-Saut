package role

import (
	"Tugas-Mini-Project/infrastructure/database"
	"Tugas-Mini-Project/repository/role"
	service "Tugas-Mini-Project/service/role"
	"github.com/labstack/echo/v4"
)

func Routes(echo *echo.Echo, conf database.Config) {
	db := database.InitDB(conf)

	reporsitory := role.NewRoleRepository(db)
	svc := service.NewRoleService(reporsitory, conf)

	controller := RoleHandler{
		Svc: svc,
	}

	echo.POST("/role", controller.CreateRole)
}
