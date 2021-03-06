package discussions

import (
	"Tugas-Mini-Project/infrastructure/database"
	discussion "Tugas-Mini-Project/interface/discussions"
	"Tugas-Mini-Project/repository/discussions"
	service "Tugas-Mini-Project/service/discussions"
	"github.com/labstack/echo/v4"
)

func Routes(echo *echo.Echo, conf database.Config) {
	db := database.InitDB(conf)

	repo := discussions.NewDiscussionsRepository(db)
	svc := service.NewDiscussionsService(repo, conf)
	controller := discussion.DiscussionHandler{
		Svc: svc,
	}
	echo.GET("/discussions", controller.GetAllDiscussions)
	echo.POST("/discussions", controller.CreateDiscussions)
}
