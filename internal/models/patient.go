package models

import "gorm.io/gorm"

type Patient struct {
	gorm.Model
	Name string
}
