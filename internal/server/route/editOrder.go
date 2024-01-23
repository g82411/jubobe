package route

import (
	"JuboTest/internal/businessLogic"
	"JuboTest/internal/utils"
	"context"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

func EditOrder(c *fiber.Ctx) error {
	type orderBody struct {
		OrderId   int    `json:"order_id"`
		PrevOrder string `json:"prev_order"`
		Text      string `json:"text"`
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
	order, err := businessLogic.QueryOrderById(connCtx, body.OrderId)
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.Send([]byte("internal error"))
	}
	if order == nil {
		c.Status(fiber.StatusNotFound)
		return c.Send([]byte("invalid order id"))
	}
	if order.Text != body.PrevOrder {
		c.Status(fiber.StatusConflict)
		return c.Send([]byte("order has been modified"))
	}
	err = businessLogic.EditOrder(connCtx, body.OrderId, body.PrevOrder, body.Text)
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.Send([]byte("internal error"))
	}
	c.Status(fiber.StatusCreated)
	return c.JSON(fiber.Map{
		"message": "success",
	})
}
