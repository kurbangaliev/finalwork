package db

import (
	"finalwork/internal/models"
	"log"
	"time"

	"gorm.io/gorm"
)

func saveManager(name string, jobTitle string, address string, phone string, email string, schedule string, image string) error {
	db, err := DbConnection()
	if err != nil {
		return err
	}
	manager := models.Manager{
		Name:     name,
		JobTitle: jobTitle,
		Address:  address,
		Phone:    phone,
		Email:    email,
		Schedule: schedule,
		Image:    image,
	}
	result := db.Create(&manager)
	if result.Error != nil {
		log.Println(result.Error)
		return result.Error
	}
	return nil
}

func SaveManager(item models.ManagerItem) error {
	return saveManager(item.Name, item.JobTitle, item.Address, item.Phone, item.Email, item.Schedule, item.Image)
}

func updateManager(id uint, name string, jobTitle string, address string, phone string, email string, schedule string, image string) error {
	db, err := DbConnection()
	if err != nil {
		return err
	}
	manager := models.Manager{
		Model: gorm.Model{
			ID:        id,
			UpdatedAt: time.Time{},
		},
		Name:     name,
		JobTitle: jobTitle,
		Address:  address,
		Phone:    phone,
		Email:    email,
		Schedule: schedule,
		Image:    image,
	}
	result := db.Updates(&manager)
	if result.Error != nil {
		log.Println(result.Error)
		return result.Error
	}
	return nil
}

func UpdateManager(item models.ManagerItem) error {
	return updateManager(item.Id, item.Name, item.JobTitle, item.Address, item.Phone, item.Email, item.Schedule, item.Image)
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
