package controllers

import (
	"errors"
	"net/http"
	"os"
	"log"
	"time"

	"gorm.io/gorm"
	"golang.org/x/crypto/bcrypt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/raylicola/NFlaquiz/database"
	"github.com/raylicola/NFlaquiz/models"
)


// ユーザー情報の取得
func User(c *gin.Context) (*models.User, error){
	cookie, err := c.Cookie("jwt")

	if err != nil {
		return nil, err
	}

	token, err := jwt.Parse(cookie, func(token *jwt.Token) (interface{}, error) {
		return []byte("SECRET_KEY"), nil
	})
	if err != nil {
		return nil, err
	}
	claims := token.Claims.(jwt.MapClaims)
	id := claims["sub"]

	var user models.User
	database.DB.Where("id = ?", id).First(&user)

	return &user, nil
}

/*
ログイン
  email: メールアドレス
  password: パスワード
*/
func Login(c *gin.Context) {
	var user models.User
  email := c.PostForm("email")
  password := c.PostForm("password")

  res := database.DB.Where("email = ?", email).First(&user)
	if res.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err_msg": res.Error})
		return
	}

	// パスワードが間違っている場合
	if err := bcrypt.CompareHashAndPassword(user.Password, []byte(password)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err_msg": "メールアドレスまたはパスワードが違います"})
		return
	}

  // ユーザー認証の設定
  claims := jwt.MapClaims{
		"sub": user.ID,
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

  c.JSON(http.StatusOK, gin.H{"jwt": tokenString})
}


// ログアウト
func Logout(c *gin.Context) {
	log.Println(c.Cookie("jwt"))
	c.SetCookie("jwt", "", 3600, "/", "localhost", true, true)
	c.JSON(http.StatusOK, gin.H{"msg": "Logout"})
}


/*
ユーザー登録
Parameters
  email: メールアドレス
  password: パスワード
  password_confirm: 確認用パスワード
*/
func Signup(c *gin.Context) {
	var user models.User
	email := c.PostForm("email")
	password := c.PostForm("password")
	password_confirm := c.PostForm("password_confirm")

	// 登録済みの場合
	res := database.DB.Where("email = ?", email).First(&user)

	if !errors.Is(res.Error, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusInternalServerError, gin.H{"err_msg": "メールアドレスが既に登録されています"})
		return
	}

	// パスワードが一致しない場合
	if password != password_confirm {
		c.JSON(http.StatusInternalServerError, gin.H{"err_msg": "パスワードが一致しません"})
		return
	}
	// パスワードのエンコード
	hashed_password, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err_msg": "パスワードの暗号化でエラーが発生しました"})
		return
	}

	user = models.User{
		Email:    email,
		Password: hashed_password,
	}

	database.DB.Create(&user)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err_msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"user": user})
}