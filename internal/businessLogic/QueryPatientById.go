package businessLogic

import (
	"JuboTest/internal/models"
	"context"

	"gorm.io/gorm"
)

func QueryPatientById(ctx context.Context, patientId int) (*models.Patient, error) {
	db := ctx.Value("db").(*gorm.DB)
	var patient models.Patient
	db.Model(&models.Patient{}).Where("id = ?", patientId).Find(&patient)
	return &patient, db.Error
}
