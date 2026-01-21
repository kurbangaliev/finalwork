package models

import "gorm.io/gorm"

type Slogan struct {
	gorm.Model
	Slogan  string `gorm:"size:255"`
	Visible bool   `gorm:"default:false"`
}

type Manager struct {
	gorm.Model
	Name      string `gorm:"size:255"`
	JobTitle  string `gorm:"size:255"`
	Address   string `gorm:"size:255"`
	Phone     string `gorm:"size:255"`
	Email     string `gorm:"size:255"`
	Schedule  string `gorm:"size:255"`
	ImagePath string `gorm:"size:255"`
}

type News struct {
	gorm.Model
	Title     string `gorm:"type:text"`
	Content   string `gorm:"type:text"`
	Date      string `gorm:"size:20"`
	ImagePath string `gorm:"size:1024"`
}

type Partner struct {
	gorm.Model
	Name      string `gorm:"size:255"`
	ImagePath string `gorm:"size:255"`
}
type Image struct {
	gorm.Model
	Name       string
	Data       []byte
	ServerPath string `gorm:"size:255"`
	FolderPath string `gorm:"size:255"`
}
