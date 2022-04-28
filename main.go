package main

import "Tugas-Mini-Project/infrastructure/http/server"

func main() {
	app := server.Server()
	err := app.Start(":8080")
	if err != nil {
		return
	}
}
