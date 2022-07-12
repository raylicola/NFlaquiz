package models

type AnswerStatus struct {
	CountryID  string  `json:"country_id"`
	Answer     int     `json:"answer"`
	Bookmark   int     `json:"bookmark"`
}