package utils

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/raylicola/NFlaquiz/database"
	"github.com/raylicola/NFlaquiz/models"
)

// Cookieからユーザー認証を行う
// 戻り値:
//   成功時：ユーザー情報
//   失敗時：エラー情報
func AuthUser(c *gin.Context) (*models.User, error){
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