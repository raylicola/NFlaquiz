package models

type QuizResult struct {
  ID          int
  Weight      int  `json:"weight"`
  QuizID      int
  Quiz        Quiz
  UserID      int
  User        User
}