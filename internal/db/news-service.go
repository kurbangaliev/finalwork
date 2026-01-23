package db

import (
	"finalwork/internal/models"
	"log"
	"time"

	"gorm.io/gorm"
)

func saveNews(title string, content string, date string, image string) error {
	db, err := DbConnection()
	if err != nil {
		return err
	}
	news := models.News{
		Title:   title,
		Content: content,
		Date:    date,
		Image:   image,
	}
	result := db.Create(&news)
	if result.Error != nil {
		log.Println(result.Error)
		return result.Error
	}
	return nil
}

func SaveNews(newsItem models.NewsItem) error {
	return saveNews(newsItem.Title, newsItem.Content, newsItem.Date, newsItem.Image)
}

func updateNews(id uint, title string, content string, date string, image string) error {
	db, err := DbConnection()
	if err != nil {
		return err
	}
	news := models.News{
		Model: gorm.Model{
			ID:        id,
			UpdatedAt: time.Time{},
		},
		Title:   title,
		Content: content,
		Date:    date,
		Image:   image,
	}
	result := db.Updates(&news)
	if result.Error != nil {
		log.Println(result.Error)
		return result.Error
	}
	return nil
}

func UpdateNews(newsItem models.NewsItem) error {
	return updateNews(newsItem.Id, newsItem.Title, newsItem.Content, newsItem.Date, newsItem.Image)
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
