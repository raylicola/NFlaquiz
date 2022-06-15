package models

type FlagColor struct {
	ID         int     `json:"id" gorm:"primary_key"`
	CountryID  string  `json:"country_id"`
	Country    Country
	ColorID    string  `json:"color_id"`
	Color      Color
}