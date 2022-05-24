package controllers

import (
	"errors"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/raylicola/NFlaquiz/database"
	"github.com/raylicola/NFlaquiz/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// ユーザー登録
// Parameters
//    email: メールアドレス
//    password: パスワード
//    password_confirm: 確認用パスワード
//
// Returns
//    err: error 処理中に発生したエラー
//    &user: models.User 新規登録したユーザーの情報
func Signup(email, password, password_confirm string) (*models.User, error) {
  var user models.User

  // 登録済みの場合
  err := database.DB.Where("email = ?", email).First(&user).Error
  if !errors.Is(err, gorm.ErrRecordNotFound) {
    err := errors.New("メールアドレスが既に登録されています。")
    return nil, err
  }

  // パスワードが一致しない場合
  if password != password_confirm {
    err := errors.New("パスワードが一致しません。")
    return nil, err
  }
  // パスワードのエンコード
  hashed_password, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

  if err != nil {
		err := errors.New("パスワードの暗号化でエラーが発生しました。")
		return nil, err
	}

  user = models.User{
    Email: email,
    Password: hashed_password,
  }

  database.DB.Create(&user)

  return &user, nil
}

func GetSignup(c *gin.Context) {
  c.JSON(http.StatusOK, gin.H{"message": "Get Signup"})
}

func PostSignup(c *gin.Context) {
  log.Println(c.PostForm("email"))
  email := c.PostForm("email")
  password := c.PostForm("password")
  password_confirm := c.PostForm("password_confirm")

  user, err := Signup(email, password, password_confirm)

  if err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
    return
  }
  c.JSON(http.StatusOK, gin.H{"user": user})
}
