package httpServer

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func Init() *fiber.App {
	app := fiber.New(fiber.Config{
		AppName: "JuboOfflineTest",
	})
	app.Use(logger.New())
	app.Use(cors.New())
	BindingRouter(app)
	return app
}
