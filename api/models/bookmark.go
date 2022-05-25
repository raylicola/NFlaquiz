package models

import (
  "gorm.io/gorm"
)

type Bookmark struct {
	gorm.Model
	CountryId string
	UserID    int
	User      User
}