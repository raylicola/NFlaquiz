package models

import (
	"gorm.io/gorm"
)

type QuizResult struct {
	gorm.Model
  Weight      int
  QuizId      int
  UserId      int
  User        User
}