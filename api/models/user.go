package models

type User struct {
  Email      string  `gorm:"primary_key"`
  Password   []byte
}