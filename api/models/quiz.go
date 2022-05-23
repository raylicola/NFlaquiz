package models

type Quiz struct {
	ID uint `gorm:"primary_key"`
	Name string
}