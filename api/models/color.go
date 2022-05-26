package models

type Color struct {
	ID string `gorm:"primary_key"`
	Name string
}