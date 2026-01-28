package db

import (
	"log"
)

func SelectAll[T comparable]() ([]T, error) {
	var items []T

	db, err := DbConnection()
	if err != nil {
		log.Fatal(err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}
	defer sqlDB.Close()

	result := db.Find(&items)
	if result.Error != nil {
		log.Fatal(result.Error)
	}

	return items, nil
}
