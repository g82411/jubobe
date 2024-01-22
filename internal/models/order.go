package models

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	ID        int
	Text      string
	IsNewest  bool
	PatientID uint
}
