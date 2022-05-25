package models

type Country struct {
	ID uint `gorm:"primary_key"`
	Name string
	AreaID uint
	Description string
}