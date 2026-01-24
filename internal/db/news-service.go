package db

import (
	"finalwork/internal/models"
	"log"
	"time"

	"gorm.io/gorm"
)

func SaveNews(newsItem models.News) error {
	db, err := DbConnection()
	if err != nil {
		return err
	}
	result := db.Create(&newsItem)
	if result.Error != nil {
		log.Println(result.Error)
		return result.Error
	}
	return nil
}

func UpdateNews(newsItem models.News) error {
	db, err := DbConnection()
	if err != nil {
		return err
	}
	newsItem.UpdatedAt = time.Now()
	result := db.Updates(&newsItem)
	if result.Error != nil {
		log.Println(result.Error)
		return result.Error
	}
	return nil
}

func DeleteNews(id uint) error {
	db, err := DbConnection()
	if err != nil {
		return err
	}
	news := models.News{
		Model: gorm.Model{
			ID: id,
		},
	}
	result := db.Delete(&news)
	if result.Error != nil {
		log.Println(result.Error)
		return result.Error
	}
	return nil
}
