package models

type Bookmark struct {
	ID          int     `json:"id" gorm:"primary_key"`
	CountryID   string  `json:"country_id"`
	Country     Country
	UserID      int     `json:"user_id"`
	User        User
}