package models

type Quiz struct {
	ID        uint `gorm:"primary_key"`
	Hiragana  string
	CountryID string
}