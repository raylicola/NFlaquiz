package models

type Bookmark struct {
	ID          int
	CountryID   string
	Country     Country
	UserID      int
	User        User
}