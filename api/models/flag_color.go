package models

type FlagColor struct {
	ID         int    `gorm:"primary_key"`
	CountryID  string
	Country    Country
	ColorID    string
	Color      Color
}