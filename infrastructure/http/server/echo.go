package server

import (
	"Tugas-Mini-Project/infrastructure/http/route/auth"
	"Tugas-Mini-Project/infrastructure/http/route/role"
	"github.com/labstack/echo/v4"
)

func Server() *echo.Echo {
	app := echo.New()
	auth.Routes(app)
	role.Routes(app)
	return app
}
