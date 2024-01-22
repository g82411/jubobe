package main

import (
	"JuboTest/internal/bootstrap/httpServer"
)

func main() {
	app := httpServer.Init()
	app.Listen(":8080")
}
