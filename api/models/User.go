package models

type User struct {
  ID uint `gorm:"primary_key"`
  Email      string
  Password   string
}