package main

import (
	"Tugas-Mini-Project/infrastructure/http/server"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title School Discussions Room API Documentation
// @description This is School Discussions Room API
// @version 2.0
// @host localhost:8080
// @BasePath /api/v1
// @schemes http https
// @securityDefinitions.apiKey JWT
// @in header
// @name Authorization
func main() {
	// e := echo.New()
	app := server.Server()
	app.GET("/swagger/*", echoSwagger.WrapHandler)
	err := app.Start(":8080")
	if err != nil {
		return
	}
}
