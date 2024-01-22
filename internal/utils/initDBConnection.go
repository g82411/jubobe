package utils

import (
	"JuboTest/internal/config"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func resolveDSN() string {
	c := config.GetConfig()
	host := c.DBHost
	user := c.DBUser
	pass := c.DBPass
	name := c.DBName
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=5432 sslmode=disable TimeZone=Asia/Taipei", host, user, pass, name)
	return dsn
}

func InitDBConnection() (*gorm.DB, error) {
	dsn := resolveDSN()
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
