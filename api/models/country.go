package models

type Country struct {
	ID            string  `json:"id" gorm:"primary_key"`
	Name          string  `json:"name"`
	AreaID        string
	Area          Area
	Description   string  `json:"description"`
}