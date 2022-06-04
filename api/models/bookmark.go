package models

import (
  "gorm.io/gorm"
)

type Bookmark struct {
	gorm.Model
	CountryId   string
	Country     Country
	UserID      int
	User        User
}