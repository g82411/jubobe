package route

import (
	"JuboTest/internal/businessLogic"
	"JuboTest/internal/utils"
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log"
	"strconv"
)

func PatientList(c *fiber.Ctx) error {
	// FIXME: now time page size is fixed
	const PageSize = 10
	page := c.Query("page")
	if page == "" {
		page = "0"
	}
	intPage, err := strconv.Atoi(page)
	if err != nil {
		intPage = 0
	}
	ctx := context.Context(context.Background())
	db, err := utils.InitDBConnection()
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		log.Fatal(fmt.Errorf("init connection fail %v", err))
	}
	connCtx := context.WithValue(ctx, "db", db)
	patitntEntities, err := businessLogic.ListPatient(connCtx, PageSize, intPage)
	var patients []map[string]interface{}
	for _, patient := range patitntEntities {
		jsonPat := map[string]interface{}{
			"id":   patient.ID,
			"name": patient.Name,
		}
		patients = append(patients, jsonPat)
	}

	return c.JSON(map[string]interface{}{
		"patients": patients,
	})

}
