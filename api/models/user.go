package models

import (
  "gorm.io/gorm"
)

type User struct {
  gorm.Model
  Email      string  `gorm:"primary_key"`
  Password   []byte
}