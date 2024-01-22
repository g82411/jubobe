package httpServer

import (
	"JuboTest/internal/server/route"
	"github.com/gofiber/fiber/v2"
)

func Handler(c *fiber.Ctx) error {
	return c.SendString("Hello, World 👋!")
}

func BindingRouter(app *fiber.App) {
	app.Get("/patients", route.PatientList)
	app.Get("/patient/:patientID", Handler)
	app.Post("/order", Handler)
	app.Patch("/order/:orderID", Handler)
}
