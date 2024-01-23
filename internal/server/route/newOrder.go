package route

import (
	"JuboTest/internal/businessLogic"
	"JuboTest/internal/utils"
	"context"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

func NewOrder(c *fiber.Ctx) error {
	type orderBody struct {
		PatientID   int    `json:"patient_id"`
		LastOrderID int    `json:"last_order_id"`
		Text        string `json:"text"`
	}
	var body orderBody
	if err := c.BodyParser(&body); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.Send([]byte("invalid body"))
	}
	ctx := context.Context(context.Background())
	db, err := utils.InitDBConnection()
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		log.Fatal(fmt.Errorf("init connection fail %v", err))
	}
	connCtx := context.WithValue(ctx, "db", db)
	patient, err := businessLogic.QueryPatientById(connCtx, body.PatientID)
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.Send([]byte("internal error"))
	}
	if patient == nil {
		c.Status(fiber.StatusNotFound)
		return c.Send([]byte("invalid patient id"))
	}
	lastOrderId := -1
	lastOrder, err := businessLogic.QueryLastestOrderByPatientId(connCtx, body.PatientID)
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.Send([]byte("internal error"))
	}
	if lastOrder != nil {
		lastOrderId = int(lastOrder.ID)
	}
	if lastOrderId != body.LastOrderID {
		c.Status(fiber.StatusConflict)
		return c.Send([]byte("data not updated"))
	}
	createdOrder, err := businessLogic.NewOrder(connCtx, body.PatientID, body.Text)
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.Send([]byte("internal error"))
	}
	c.Status(fiber.StatusCreated)
	return c.JSON(fiber.Map{
		"message": map[string]interface{}{
			"id":   createdOrder.ID,
			"text": createdOrder.Text,
		},
	})
}
