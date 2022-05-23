package database

import (
  "github.com/raylicola/NFlaquiz/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
  user = "admin"
  password = "admin"
  protocol = "tcp(db:3306)"
  dbname = "nflaquiz"
	dsn = user + ":" + password + "@" + protocol + "/" + dbname + "?charset=utf8mb4&parseTime=True&loc=Local"
)

func Connect() {
  print(dsn)
  db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
  // 接続できなかった場合
  if err != nil {
    panic(err.Error())
  }

  db.AutoMigrate(
    &models.User{},
    &models.Bookmark{},
    &models.QuizResult{},
  )
}