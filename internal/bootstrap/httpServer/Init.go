package httpServer

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func Init() *fiber.App {
	app := fiber.New(fiber.Config{
		AppName: "JuboOfflineTest",
	})
	app.Use(logger.New())
	BindingRouter(app)
	return app
}
