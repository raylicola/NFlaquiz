package models

type QuizResult struct {
  ID          int      `json:"id" gorm:"primary_key"`
  Weight      float64  `json:"weight"`
  QuizID      int      `json:"quiz_id"`
  Quiz        Quiz
  UserID      int      `json:"user_id"`
  User        User
}