package models

import "gorm.io/gorm"

type Manager struct {
	gorm.Model
	Name     string `gorm:"size:255"`
	JobTitle string `gorm:"size:255"`
	Address  string `gorm:"size:255"`
	Phone    string `gorm:"size:255"`
	Email    string `gorm:"size:255"`
	Schedule string `gorm:"size:255"`
	Image    string `gorm:"size:255"`
}

type News struct {
	gorm.Model
	Title   string `gorm:"type:text"`
	Content string `gorm:"type:text"`
	Date    string `gorm:"size:20"`
	Image   string `gorm:"size:1024"`
}

type Partner struct {
	gorm.Model
	Name  string `gorm:"size:255"`
	Image string `gorm:"size:255"`
}
