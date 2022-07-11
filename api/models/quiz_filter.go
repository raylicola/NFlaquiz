package models

type QuizFilter struct {
	Colors    []string  `form:"colors[]" binding:"required"`
	Areas     []string  `form:"areas[]" binding:"required"`
	Bookmark  string    `form:"bookmark" binding:"required"`
}