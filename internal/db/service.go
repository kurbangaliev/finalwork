package db

import (
	"finalwork/internal/models"
	"fmt"
	"log"
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
	for _, manager := range managers {
		fmt.Println(manager.Name)
	}
	return managers
}

func SelectAllNews() []models.News {
	var news []models.News
	db, err := DbConnection()
	if err != nil {
		log.Fatal(err)
	}
	result := db.Find(&news)
	if result.Error != nil {
		log.Fatal(result.Error)
	}
	for _, newsItem := range news {
		log.Println(newsItem.Title)
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

func SaveNews(title string, content string, date string, imagePath string) error {
	db, err := DbConnection()
	if err != nil {
		return err
	}
	news := models.News{
		Title:     title,
		Content:   content,
		Date:      date,
		ImagePath: imagePath,
	}
	result := db.Create(&news)
	if result.Error != nil {
		log.Println(result.Error)
		return result.Error
	}
	return nil
}
