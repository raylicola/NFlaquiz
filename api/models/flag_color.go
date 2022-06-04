package models

type FlagColor struct {
	ID         uint    `gorm:"primary_key"`
	CountryID  string
	Country    Country
	ColorID    string
	Color      Color
}