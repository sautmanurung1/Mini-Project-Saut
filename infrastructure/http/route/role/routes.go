package role

import (
	ur "Tugas-Mini-Project/handler/role"
	"Tugas-Mini-Project/infrastructure/database"
	"Tugas-Mini-Project/repository/role"
	"github.com/labstack/echo/v4"
)

func Routes(echo *echo.Echo) {
	db := database.InitDB()

	reporsitory := role.NewRoleRepository(db)
	handler := ur.NewRoleHandler(reporsitory)

	echo.GET("/role/:id", handler.GetRole)
	echo.POST("/role", handler.CreateRole)
}
