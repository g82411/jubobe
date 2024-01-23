package businessLogic

import (
	"JuboTest/internal/models"
	"context"

	"gorm.io/gorm"
)

func ListOrderByPatient(ctx context.Context, patientId, pageSize, page int) ([]models.Order, error) {
	db := ctx.Value("db").(*gorm.DB)
	var orders []models.Order
	tx := db.Select("id", "text")
	tx = tx.Where("patient_id = ?", patientId)
	tx = tx.Order("id desc")
	tx = tx.Limit(pageSize)
	tx = tx.Offset(page * pageSize)
	tx = tx.Find(&orders)
	return orders, tx.Error
}
