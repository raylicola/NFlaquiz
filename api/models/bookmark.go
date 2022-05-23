package models

type Bookmark struct {
	ID uint `gorm:"primary_key"`
	CountryId string
	UserID    int
	User      User
}