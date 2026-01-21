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
	Title     string `gorm:"size:255"`
	Content   string `gorm:"size:255"`
	Date      string `gorm:"size:255"`
	ImagePath string `gorm:"size:255"`
}

type Partner struct {
	gorm.Model
	Name      string `gorm:"size:255"`
	ImagePath string `gorm:"size:255"`
}
type Image struct {
	gorm.Model
	Name         string
	Data         []byte
	RelativePath string `gorm:"size:255"`
}
