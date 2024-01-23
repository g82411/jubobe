package businessLogic

import (
	"JuboTest/internal/models"
	"context"

	"gorm.io/gorm"
)

func QueryOrderById(ctx context.Context, orderId int) (*models.Order, error) {
	db := ctx.Value("db").(*gorm.DB)
	var order models.Order
	db.Model(&models.Order{}).Where("id = ?", orderId).Find(&order)
	return &order, db.Error
}
