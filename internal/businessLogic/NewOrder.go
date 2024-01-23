package businessLogic

import (
	"JuboTest/internal/models"
	"context"
	"fmt"

	"gorm.io/gorm"
)

func NewOrder(ctx context.Context, patientId int, text string) (*models.Order, error) {
	db := ctx.Value("db").(*gorm.DB)
	order := models.Order{
		PatientID: patientId,
		Text:      text,
	}
	db.Create(&order)
	if db.Error != nil {
		return nil, fmt.Errorf("Error when create order %v", db.Error)
	}
	return &order, nil
}
