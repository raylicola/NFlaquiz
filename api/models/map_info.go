package models

type MapInfo struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Bookmark    int     `json:"bookmark"`
	Weight      float64 `json:"weight"`
}