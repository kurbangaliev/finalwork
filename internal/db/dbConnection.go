package db

import (
	"finalwork/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func DbConnection() (*gorm.DB, error) {
	log.Println("Postgres Connect")
	pgConnect := "host=localhost dbname=finalwork user=postgres password=postgres port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(pgConnect), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return db, nil
}

func AutoMigrate() error {
	db, err := DbConnection()
	if err != nil {
		log.Fatal(err)
		return err
	}
	log.Println(db)
	err = db.AutoMigrate(&models.Manager{})
	if err != nil {
		log.Fatal(err)
		return &MigrateError{
			Model:   "Manager",
			Message: err.Error(),
			Err:     ErrAutoMigrate, // Wrap the sentinel error
		}
	}
	err = db.AutoMigrate(&models.News{})
	if err != nil {
		log.Fatal(err)
		return &MigrateError{
			Model:   "News",
			Message: err.Error(),
			Err:     ErrAutoMigrate, // Wrap the sentinel error
		}
	}
	err = db.AutoMigrate(&models.Partner{})
	if err != nil {
		log.Fatal(err)
		return &MigrateError{
			Model:   "Partner",
			Message: err.Error(),
			Err:     ErrAutoMigrate, // Wrap the sentinel error
		}
	}
	err = db.AutoMigrate(&models.Image{})
	if err != nil {
		log.Fatal(err)
		return &MigrateError{
			Model:   "Image",
			Message: err.Error(),
			Err:     ErrAutoMigrate, // Wrap the sentinel error
		}
	}
	return nil
}
