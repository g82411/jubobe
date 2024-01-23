package route

import (
	"JuboTest/internal/businessLogic"
	"JuboTest/internal/utils"
	"context"
	"fmt"
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func ListOrder(c *fiber.Ctx) error {
	const PageSize = 10
	// FIXME: now time page size is fixed
	patientIdStr := c.Params("patientID")
	patientId, err := strconv.Atoi(patientIdStr)
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		c.Send([]byte("invalid patient id"))
	}
	pageStr := c.Query("page")
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		page = 0
	}

	ctx := context.Context(context.Background())
	db, err := utils.InitDBConnection()
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		log.Fatal(fmt.Errorf("init connection fail %v", err))
	}
	connCtx := context.WithValue(ctx, "db", db)
	orderEntities, err := businessLogic.ListOrderByPatient(connCtx, patientId, PageSize, page)
	var orders []map[string]interface{}
	for _, order := range orderEntities {
		jsonPat := map[string]interface{}{
			"id":   order.ID,
			"text": order.Text,
		}
		orders = append(orders, jsonPat)
	}

	return c.JSON(map[string]interface{}{
		"orders": orders,
	})

}
