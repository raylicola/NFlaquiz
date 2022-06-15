package models

type QuizResult struct {
  ID          int
  Weight      float64  `json:"weight"`
  QuizID      int
  Quiz        Quiz
  UserID      int
  User        User
}