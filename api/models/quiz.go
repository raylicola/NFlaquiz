package models

type Quiz struct {
	ID         int     `json:"id" gorm:"primary_key"`
	Hiragana   string  `json:"hiragana"`
	CountryID  string  `json:"country_id"`
	Country    Country
	Hint1      string  `json:"hint1"`
	Hint2      string  `json:"hint2"`
	Hint3      string  `json:"hint3"`
}