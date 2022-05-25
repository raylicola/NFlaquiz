package models

type FlagColor struct {
	ID        uint `gorm:"primary_key"`
	CountryID string
	ColorID   string
}