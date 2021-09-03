package config

import (
	"fmt"
	"log"
	"os"

	"github.com/subosito/gotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Config struct {
	AppName     string
	AppPort     int
	LogLevel    string
	Environment string
	JWTSecret   string
}

func Init() *Config {
	defaultEnv := ".env"

	if err := gotenv.Load(defaultEnv); err != nil {
		log.Fatal("failed load .env")
	}

	log.SetOutput(os.Stdout)

	appConfig := &Config{
		AppName:     GetString("APP_NAME"),
		AppPort:     GetInt("APP_PORT"),
		LogLevel:    GetString("LOG_LEVEL"),
		Environment: GetString("ENVIRONMENT"),
		JWTSecret:   GetString("JWT_SECRET"),
	}

	return appConfig
}

func Database() (*gorm.DB, error) {
	dsn := "root:rakamin@tcp(127.21.0.1:3306)/earth?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Errorf("Cannot connect db", err)
		return nil, err
	}

	return db, nil
}
