package db

import (
	"errors"
	"finalwork/internal/models"
	"log"
	"time"

	"gorm.io/gorm"
)

// SelectAllUsers - загрузка всех пользователей
func SelectAllUsers() []models.User {
	var users []models.User

	db, err := DbConnection()
	if err != nil {
		log.Fatal(err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}
	defer sqlDB.Close()

	result := db.Find(&users)
	if result.Error != nil {
		log.Fatal(result.Error)
	}

	return users
}

// SaveUser - Сохранение пользователя в базу данных
func SaveUser(user models.User) error {
	db, err := DbConnection()
	if err != nil {
		return err
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}
	defer sqlDB.Close()

	result := db.Create(&user)
	if result.Error != nil {
		log.Println(result.Error)
		return result.Error
	}
	return nil
}

// UpdateUser - Обновление пользователя в базе данных
func UpdateUser(user models.User) error {
	db, err := DbConnection()
	if err != nil {
		return err
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}
	defer sqlDB.Close()

	user.UpdatedAt = time.Now()
	result := db.Updates(&user)
	if result.Error != nil {
		log.Println(result.Error)
		return result.Error
	}
	return nil
}

// DeleteUser - Удаление пользователя из базы данных
func DeleteUser(id uint) error {
	db, err := DbConnection()
	if err != nil {
		return err
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}
	defer sqlDB.Close()

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

// CreateDefaultUser - создание пользователя по умолчанию в базе данных при первом запуске системы
func CreateDefaultUser() error {
	db, err := DbConnection()
	if err != nil {
		return err
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}
	defer sqlDB.Close()

	var users []models.User
	result := db.Find(&users, "login = ?", "admin")
	if result.Error != nil {
		log.Fatal(result.Error)
	}
	if len(users) == 0 {
		adminUser := models.User{
			Login:    "admin",
			Password: "ISMvKXpXpadDiUoOSoAfww==",
			Role:     "admin",
		}
		result = db.Create(&adminUser)
		if result.Error != nil {
			log.Println(result.Error)
		}
	}
	return nil
}

// FindUserByLogin - поиск пользователя по имени пользователя
func FindUserByLogin(login string) (*models.User, error) {
	db, err := DbConnection()
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}
	defer sqlDB.Close()

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
