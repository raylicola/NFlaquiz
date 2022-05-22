package models

import (
	"gorm.io/gorm"
)

type Bookmark struct {
	gorm.Model
	CountryId string
	UserId    int
	User      User
}