package answare

import (
	"Tugas-Mini-Project/infrastructure/database"
	m "Tugas-Mini-Project/infrastructure/http/middleware"
	answares "Tugas-Mini-Project/interface/answare"
	"Tugas-Mini-Project/repository/answare"
	service "Tugas-Mini-Project/service/answare"
	"github.com/labstack/echo/v4"
)

func Routes(echo *echo.Echo, conf database.Config) {
	db := database.InitDB(conf)

	repo := answare.NewAnswareRepository(db)
	svc := service.NewServiceAnsware(repo, conf)

	controller := answares.AnswareHandler{
		Svc: svc,
	}

	student := echo.Group("/student")

	student.POST("/answare", controller.CreateAnsware, m.JWTTokenMiddleware())
	echo.GET("/answare", controller.GetAllAnsware)
	echo.GET("/answare/:id", controller.GetAnswareById)
	student.PUT("/answare/:id", controller.UpdateAnsware, m.JWTTokenMiddleware())
	student.DELETE("/answare/:id", controller.DeleteAnsware, m.JWTTokenMiddleware())
}
