package controllers

import (
	"errors"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/raylicola/NFlaquiz/database"
	"github.com/raylicola/NFlaquiz/models"
	"golang.org/x/crypto/bcrypt"
)


// ログイン
//    email: メールアドレス
//    password: パスワード
// Returns
//    &user: models.User ログインユーザーの情報
//    err: error 処理中に発生したエラー
func Login(email, password string) (*models.User, error) {
  var user models.User

  // メールアドレスが間違っているor登録されていない場合
  res := database.DB.Where("email = ?", email).First(&user)
  if res.Error != nil {
    err := errors.New("メールアドレスまたはパスワードが違います。")
    return nil, err
  }

  // パスワードが間違っている場合
  if err := bcrypt.CompareHashAndPassword(user.Password, []byte(password)); err != nil {
    err := errors.New("メールアドレスまたはパスワードが違います。")
    return nil, err
  }

  return &user,nil
}


// Returns
//    {"msg": "Get Login"}
func GetLogin(c *gin.Context) {
  c.JSON(http.StatusOK, gin.H{"msg": "Get Login"})
}


// Returns
//    {"user": models.User}
func PostLogin(c *gin.Context) {
  email := c.PostForm("email")
  password := c.PostForm("password")

  user, err := Login(email, password)

  if err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{"err_msg": err.Error()})
    return
  }

  // ユーザー認証の設定
  claims := jwt.MapClaims{
		"email": user.Email,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	}
  // ヘッダーとペイロードの生成
  token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
  // トークンに署名を付与
  tokenString, _ := token.SignedString([]byte("SECRET_KEY"))

  cookie := new(http.Cookie)
  cookie.Value = tokenString

  // ローカル環境の場合
  c.SetSameSite(http.SameSiteNoneMode)
  if os.Getenv("ENV") == "local" {
    c.SetCookie("jwt", cookie.Value, 3600, "/", "localhost", true, true)
  }

  // 本番環境の場合
  if os.Getenv("ENV") == "production" {
      c.SetCookie("jwt", cookie.Value, 3600, "/", "your_domain", true, true)
  }

  c.JSON(http.StatusOK, gin.H{"user": user})
}

