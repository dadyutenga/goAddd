package db

import (
	"log"

	"goAddd/internal/config"
	"goAddd/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// InitDB initializes the database connection
func InitDB(cfg *config.Config) error {
	var err error
	DB, err = gorm.Open(postgres.Open(cfg.DatabaseURL), &gorm.Config{})
	if err != nil {
		return err
	}

	// Auto-migrate models
	if err := DB.AutoMigrate(
		&models.User{},
		&models.Organization{},
		&models.Member{},
		&models.Board{},
		&models.Column{},
		&models.Card{},
		&models.Message{},
	); err != nil {
		return err
	}

	log.Println("Database connection established and models migrated")
	return nil
}

// CloseDB closes the database connection
func CloseDB() {
	if sqlDB, err := DB.DB(); err == nil {
		sqlDB.Close()
	}
}
