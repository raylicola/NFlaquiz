package models

type User struct {
  ID          int     `json:"id" gorm:"primary_key"`
  Email       string  `json:"email"`
  Password    []byte  `json:"password"`
}