package models

type Quiz struct {
	ID        uint `gorm:"primary_key"`
	Hiragana  string
	CountryID string
	Hint1     string
	Hint2     string
	Hint3     string
}