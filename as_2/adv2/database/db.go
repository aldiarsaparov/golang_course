package database

import (
	"fmt"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"github.com/aldik/crud_go/models"
)

func ConnectGORM() (*gorm.DB, error) {
	dsn := "host=localhost user=postgres password=28081932 dbname=user_service_db port=5432 sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}
	sqlDB.SetMaxOpenConns(10)
	sqlDB.SetMaxIdleConns(5)
	sqlDB.SetConnMaxLifetime(time.Minute * 30)

	fmt.Println("Successfully connected to PostgreSQL using GORM!")
	return db, nil
}

func AutoMigrate(db *gorm.DB) error {
	err := db.AutoMigrate(&models.User{}, &models.Profile{})
	if err != nil {
		return fmt.Errorf("failed to migrate models: %w", err)
	}
	fmt.Println("Tables created successfully with associations!")
	return nil
}
