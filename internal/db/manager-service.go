package db

import (
	"finalwork/internal/models"
	"log"
	"time"

	"gorm.io/gorm"
)

func SaveManager(item models.Manager) error {
	db, err := DbConnection()
	if err != nil {
		return err
	}
	result := db.Create(&item)
	if result.Error != nil {
		log.Println(result.Error)
		return result.Error
	}
	return nil
}

func UpdateManager(item models.Manager) error {
	db, err := DbConnection()
	if err != nil {
		return err
	}
	item.UpdatedAt = time.Now()
	result := db.Updates(&item)
	if result.Error != nil {
		log.Println(result.Error)
		return result.Error
	}
	return nil
}

func DeleteManager(id uint) error {
	db, err := DbConnection()
	if err != nil {
		return err
	}
	manager := models.Manager{
		Model: gorm.Model{
			ID: id,
		},
	}
	result := db.Delete(&manager)
	if result.Error != nil {
		log.Println(result.Error)
		return result.Error
	}
	return nil
}
