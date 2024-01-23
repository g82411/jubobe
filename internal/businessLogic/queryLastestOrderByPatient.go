package businessLogic

import (
	"JuboTest/internal/models"
	"context"
	"gorm.io/gorm"
)

func QueryLastestOrderByPatientId(ctx context.Context, patientId int) (*models.Order, error) {
	db := ctx.Value("db").(*gorm.DB)
	var order models.Order
	db.Model(&models.Order{}).Where("patient_id = ?", patientId).Order("id desc").Limit(1).Find(&order)
	return &order, db.Error
}
