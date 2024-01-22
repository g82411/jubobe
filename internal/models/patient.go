package models

import "gorm.io/gorm"

type Patient struct {
	gorm.Model
	ID            uint `gorm:"primarykey"`
	Name          string
	NewestOrderId int
	NewestOrder   Order `gorm:"foreignKey:PatientID"`
}
