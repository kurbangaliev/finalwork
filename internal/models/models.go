package models

import "gorm.io/gorm"

type Manager struct {
	gorm.Model
	Name     string `gorm:"size:255" json:"name"`
	JobTitle string `gorm:"size:255" json:"jobTitle"`
	Address  string `gorm:"size:255" json:"address"`
	Phone    string `gorm:"size:255" json:"phone"`
	Email    string `gorm:"size:255" json:"email"`
	Schedule string `gorm:"size:255" json:"schedule"`
	Image    string `gorm:"size:255" json:"image"`
}

type News struct {
	gorm.Model
	Title   string `gorm:"type:text" json:"title"`
	Content string `gorm:"type:text" json:"content"`
	Date    string `gorm:"size:20" json:"date"`
	Image   string `gorm:"size:1024" json:"image"`
}

type Partner struct {
	gorm.Model
	Name  string `gorm:"size:255" json:"name"`
	Image string `gorm:"size:255" json:"image"`
}

type User struct {
	gorm.Model
	Login    string `gorm:"size:255" json:"login"`
	Password string `gorm:"size:255" json:"password"`
}
