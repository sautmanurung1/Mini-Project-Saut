package role

import (
	"Tugas-Mini-Project/infrastructure/database"
	ur "Tugas-Mini-Project/internal/handler/role"
	"Tugas-Mini-Project/internal/repository/role"
	"github.com/labstack/echo/v4"
)

func Routes(echo *echo.Echo) {
	db := database.InitDB()

	reporsitory := role.NewRoleRepository(db)
	handler := ur.NewRoleHandler(reporsitory)

	echo.GET("/role/:id", handler.GetRole)
	echo.POST("/role", handler.CreateRole)
}
