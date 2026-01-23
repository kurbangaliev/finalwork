package db

import (
	"finalwork/internal/models"
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
