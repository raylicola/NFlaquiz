package models

type Country struct {
	ID            string  `json:"id" gorm:"primary_key"`
	Name          string  `json:"name"`
	AreaID        string  `json:"area_id"`
	Area          Area
	Description   string  `json:"description"`
}