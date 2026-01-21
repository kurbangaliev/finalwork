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
