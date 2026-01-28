package db

import (
	"log"
)

func SaveObject[T comparable](item T) error {
	db, err := DbConnection()
	if err != nil {
		return err
	}
	result := db.Create(&item)
	if result.Error != nil {
		log.Println(result.Error)
		return result.Error
	}
	return nil
}

func DeleteObject[T comparable](item T) error {
	db, err := DbConnection()
	if err != nil {
		return err
	}

	result := db.Delete(&item)
	if result.Error != nil {
		log.Println(result.Error)
		return result.Error
	}
	return nil
}

func UpdateObject[T comparable](item T) error {
	db, err := DbConnection()
	if err != nil {
		return err
	}

	result := db.Updates(&item)
	if result.Error != nil {
		log.Println(result.Error)
		return result.Error
	}
	return nil
}
