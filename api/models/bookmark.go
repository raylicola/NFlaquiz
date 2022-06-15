package models

import (
  "gorm.io/gorm"
)

type Bookmark struct {
	gorm.Model
	CountryID   string
	Country     Country
	UserID      int
	User        User
}