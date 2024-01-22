package config

import "os"

type Config struct {
	DBHost string
	DBUser string
	DBPass string
	DBName string
}

func GetConfig() Config {
	return Config{
		DBHost: os.Getenv("DB_HOST"),
		DBUser: os.Getenv("DB_USER"),
		DBPass: os.Getenv("DB_PASS"),
		DBName: os.Getenv("DB_NAME"),
	}
}
