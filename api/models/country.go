package models

type Country struct {
	ID string `gorm:"primary_key"`
	Name string
	AreaID string
	Description string
}