package db

import (
	"finalwork/internal/models"
	"log"
	"time"

	"gorm.io/gorm"
)

func SelectAllManagers() []models.Manager {
	var managers []models.Manager

	db, err := DbConnection()
	if err != nil {
		log.Fatal(err)
	}

	result := db.Find(&managers)
	if result.Error != nil {
		log.Fatal(result.Error)
	}

	return managers
}

func SelectAllNews() []models.News {
	var news []models.News
	db, err := DbConnection()
	if err != nil {
		log.Fatal(err)
	}
	result := db.Order("created_at ASC").Find(&news)
	if result.Error != nil {
		log.Fatal(result.Error)
	}

	return news
}

func SaveImage(data []byte, filename string, serverPath string, folderPath string) error {
	db, err := DbConnection()
	if err != nil {
		return err
	}
	img := models.Image{
		Name:       filename,
		Data:       data,
		ServerPath: serverPath,
		FolderPath: folderPath,
	}
	result := db.Create(&img)
	if result.Error != nil {
		log.Println(result.Error)
		return result.Error
	}
	return nil
}

func SaveNews(title string, content string, date string, image string) error {
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

func UpdateNews(id uint, title string, content string, date string, image string) error {
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
