package models

type QuizResult struct {
	ID uint `gorm:"primary_key"`
  Weight      int
  QuizId      int
  UserId      int
  User        User
}