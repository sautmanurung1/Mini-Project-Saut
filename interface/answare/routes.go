package answare

import (
	"Tugas-Mini-Project/infrastructure/database"
	m "Tugas-Mini-Project/infrastructure/http/middleware"
	"Tugas-Mini-Project/repository/answare"
	service "Tugas-Mini-Project/service/answare"
	"github.com/labstack/echo/v4"
)

func Routes(echo *echo.Echo, conf database.Config) {
	db := database.InitDB(conf)

	repo := answare.NewAnswareRepository(db)
	svc := service.NewServiceAnsware(repo, conf)

	controller := AnswareHandler{
		svc: svc,
	}

	student := echo.Group("/student")

	student.POST("/answare", controller.CreateAnswareHandler, m.JWTStudentMiddleware())
	echo.GET("/answare", controller.GetAllAnswareHandler)
	echo.GET("/answare/:id", controller.GetAnswareByIdHandler)
	student.PUT("/answare/:id", controller.UpdateAnswareHandler, m.JWTStudentMiddleware())
	student.DELETE("/answare/:id", controller.DeleteAnswareHandler, m.JWTStudentMiddleware())
}
