package models

import (
  "gorm.io/gorm"
)

type User struct {
  gorm.Model
  Email       string  `json:"email" gorm:"primary_key"`
  Password    []byte  `json:"password"`
}