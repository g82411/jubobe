package models

import "gorm.io/gorm"

type OrderLog struct {
	gorm.Model
	OrderID int
	Order   Order `gorm:"foreignKey:OrderID"`
	Text    string
}
