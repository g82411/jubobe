package businessLogic

import (
	"JuboTest/internal/models"
	"context"
	"gorm.io/gorm"
)

func ListPatient(ctx context.Context, pageSize, page int) ([]models.Patient, error) {
	db := ctx.Value("db").(*gorm.DB)
	var patients []models.Patient
	tx := db.Select("id", "name")
	tx = tx.Table("patients")
	tx = tx.Limit(pageSize)
	tx = tx.Offset(page * pageSize)
	tx = tx.Order("created_at desc")
	tx = tx.Find(&patients)
	return patients, nil
}
