package models

import (
  "gorm.io/gorm"
)

type QuizResult struct {
	gorm.Model
  Weight      int  `json:"weight"`
  QuizID      int
  Quiz        Quiz
  UserID      int
  User        User
}