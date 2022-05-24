package database

import (
  "os"
  "github.com/raylicola/NFlaquiz/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
  user = os.Getenv("MYSQL_USER")
  password = os.Getenv("MYSQL_PASSWORD")
  protocol = "tcp(db:3306)"
  dbname = os.Getenv("MYSQL_DATABASE")
	dsn = user + ":" + password + "@" + protocol + "/" + dbname + "?charset=utf8mb4&parseTime=True&loc=Local"
  DB *gorm.DB
)

func Connect() {
  print(dsn)
  db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
  // 接続できなかった場合
  if err != nil {
    panic(err.Error())
  }

  DB = db

  db.AutoMigrate(
    &models.User{},
    &models.Bookmark{},
    &models.QuizResult{},
  )
}