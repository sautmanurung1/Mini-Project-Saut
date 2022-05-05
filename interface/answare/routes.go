package answare

import (
	"Tugas-Mini-Project/infrastructure/database"
	m "Tugas-Mini-Project/infrastructure/http/middleware"
	"Tugas-Mini-Project/repository/answare"
	"github.com/labstack/echo/v4"
)

func Routes(echo *echo.Echo, conf database.Config) {
	db := database.InitDB(conf)

	repo := answare.NewAnswareRepository(db)
	handler := NewAnswareHandler(repo)

	student := echo.Group("/student")

	student.POST("/answare", handler.CreateAnswareHandler, m.JWTStudentMiddleware())
	echo.GET("/answare", handler.GetAllAnswareHandler)
	echo.GET("/answare/:id", handler.GetAnswareByIdHandler)
	student.PUT("/answare/:id", handler.UpdateAnswareHandler, m.JWTStudentMiddleware())
	student.DELETE("/answare/:id", handler.DeleteAnswareHandler, m.JWTStudentMiddleware())
}
