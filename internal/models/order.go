package models

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	Text      string
	PatientID int
}
