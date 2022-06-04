package models

import (
  "gorm.io/gorm"
)

type QuizResult struct {
	gorm.Model
  Weight      int  `json:"weight"`
  QuizId      int
  Quiz        Quiz
  UserId      int
  User        User
}