package models

type User struct {
  ID          int
  Email       string  `json:"email" gorm:"primary_key"`
  Password    []byte  `json:"password"`
}