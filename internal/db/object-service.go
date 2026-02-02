package db

import (
	"log"
)

// SelectAll - Загрузка списка объектов по передаваемому типу
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

func Select[T comparable](item T) (T, error) {
	db, err := DbConnection()
	if err != nil {
		log.Fatal(err)
	}

	// Get the underlying *sql.DB connection pool
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}

	// Defer the closing of the underlying connection pool
	defer sqlDB.Close()

	result := db.Find(&item)
	if result.Error != nil {
		log.Fatal(result.Error)
	}

	return item, nil
}

// SaveObject - Сохранение объекта в базу данных
func SaveObject[T comparable](item T) error {
	db, err := DbConnection()
	if err != nil {
		return err
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}
	defer sqlDB.Close()

	result := db.Create(&item)
	if result.Error != nil {
		log.Println(result.Error)
		return result.Error
	}
	return nil
}

// DeleteObject - Удаление объекта из базы данных
func DeleteObject[T comparable](item T) error {
	db, err := DbConnection()
	if err != nil {
		return err
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}
	defer sqlDB.Close()

	result := db.Delete(&item)
	if result.Error != nil {
		log.Println(result.Error)
		return result.Error
	}
	return nil
}

// UpdateObject - Обновление объекта в базе данных
func UpdateObject[T comparable](item T) error {
	db, err := DbConnection()
	if err != nil {
		return err
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}
	defer sqlDB.Close()

	result := db.Updates(&item)
	if result.Error != nil {
		log.Println(result.Error)
		return result.Error
	}
	return nil
}
