package db

import (
	"finalwork/internal/models"
	"log"
	"time"

	"gorm.io/gorm"
)

// Deprecated: SaveManager
func SaveManager(item models.Manager) error {
	db, err := DbConnection()
	if err != nil {
		return err
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}
	defer sqlDB.Close()

	result := db.Create(&item)
	if result.Error != nil {
		log.Println(result.Error)
		return result.Error
	}
	return nil
}

// Deprecated: UpdateManager
func UpdateManager(item models.Manager) error {
	db, err := DbConnection()
	if err != nil {
		return err
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}
	defer sqlDB.Close()

	item.UpdatedAt = time.Now()
	result := db.Updates(&item)
	if result.Error != nil {
		log.Println(result.Error)
		return result.Error
	}
	return nil
}

// Deprecated: DeleteManager
func DeleteManager(id uint) error {
	db, err := DbConnection()
	if err != nil {
		return err
	}
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}
	defer sqlDB.Close()

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
