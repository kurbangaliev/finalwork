package db

import (
	"finalwork/internal/models"
	"log"
	"time"

	"gorm.io/gorm"
)

// Deprecated: SaveNews - Сохранение новостей в базу данных
func SaveNews(newsItem models.News) error {
	db, err := DbConnection()
	if err != nil {
		return err
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}
	defer sqlDB.Close()

	result := db.Create(&newsItem)
	if result.Error != nil {
		log.Println(result.Error)
		return result.Error
	}
	return nil
}

// Deprecated: UpdateNews - Обновление новостей из базы данных
func UpdateNews(newsItem models.News) error {
	db, err := DbConnection()
	if err != nil {
		return err
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}
	defer sqlDB.Close()

	newsItem.UpdatedAt = time.Now()
	result := db.Updates(&newsItem)
	if result.Error != nil {
		log.Println(result.Error)
		return result.Error
	}
	return nil
}

// Deprecated: DeleteNews - Удаление новостей из базы данных
func DeleteNews(id uint) error {
	db, err := DbConnection()
	if err != nil {
		return err
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}
	defer sqlDB.Close()

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
