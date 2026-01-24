package db

import (
	"errors"
	"finalwork/internal/models"
	"log"
	"time"

	"gorm.io/gorm"
)

func SelectAllUsers() []models.User {
	var users []models.User

	db, err := DbConnection()
	if err != nil {
		log.Fatal(err)
	}

	result := db.Find(&users)
	if result.Error != nil {
		log.Fatal(result.Error)
	}

	return users
}

func SaveUser(user models.User) error {
	db, err := DbConnection()
	if err != nil {
		return err
	}
	result := db.Create(&user)
	if result.Error != nil {
		log.Println(result.Error)
		return result.Error
	}
	return nil
}

func UpdateUser(user models.User) error {
	db, err := DbConnection()
	if err != nil {
		return err
	}
	user.UpdatedAt = time.Now()
	result := db.Updates(&user)
	if result.Error != nil {
		log.Println(result.Error)
		return result.Error
	}
	return nil
}

func DeleteUser(id uint) error {
	db, err := DbConnection()
	if err != nil {
		return err
	}
	user := models.User{
		Model: gorm.Model{
			ID: id,
		},
	}
	result := db.Delete(&user)
	if result.Error != nil {
		log.Println(result.Error)
		return result.Error
	}
	return nil
}

func CreateDefaultUser() error {
	db, err := DbConnection()
	if err != nil {
		return err
	}

	var users []models.User
	result := db.Find(&users, "login = ?", "admin")
	if result.Error != nil {
		log.Fatal(result.Error)
	}
	if len(users) == 0 {
		adminUser := models.User{
			Login:    "admin",
			Password: "hqBWNunDWATHLyNavE+CEQ==",
		}
		result = db.Create(&adminUser)
		if result.Error != nil {
			log.Println(result.Error)
		}
	}
	return nil
}

func FindUserByLogin(login string) (*models.User, error) {
	db, err := DbConnection()
	if err != nil {
		return nil, err
	}
	var users []models.User
	result := db.Find(&users)
	if result.Error != nil {
		log.Fatal(result.Error)
	}
	for _, user := range users {
		if user.Login == login {
			return &user, nil
		}
	}
	return nil, errors.New("user not found")
}
