package models

type Hint struct {
	ID        uint `gorm:"primary_key"`
	Content   string
	CountryID string
}