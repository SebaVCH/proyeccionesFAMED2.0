package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"proyeccionesFAMED/models"
)

var DB *gorm.DB

func StartDB() error {
	var err error
	DB, err = gorm.Open(sqlite.Open("pfamed.DB"), &gorm.Config{})

	if err != nil {
		return err
	}

	if err := DB.AutoMigrate(&models.Grade{}, &models.Subject{}, &models.Student{}); err != nil {
		return err
	}

	return nil

}
