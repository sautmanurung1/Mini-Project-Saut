package server

import (
	"Tugas-Mini-Project/infrastructure/database"
	"Tugas-Mini-Project/infrastructure/http/routes/answare"
	"Tugas-Mini-Project/infrastructure/http/routes/assignment"
	"Tugas-Mini-Project/infrastructure/http/routes/auth"
	"Tugas-Mini-Project/infrastructure/http/routes/discussions"
	"Tugas-Mini-Project/infrastructure/http/routes/question"
	"Tugas-Mini-Project/infrastructure/http/routes/role"
	"github.com/labstack/echo/v4"
)

func Server() *echo.Echo {
	app := echo.New()
	conf := database.Config{}

	auth.Routes(app, conf)
	role.Routes(app, conf)
	assignment.Routes(app, conf)
	question.Routes(app, conf)
	answare.Routes(app, conf)
	discussions.Routes(app, conf)
	return app
}
