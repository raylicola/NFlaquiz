package controllers

import (
	"log"
	"net/http"
	"github.com/gin-gonic/gin"
)

func Logout(c *gin.Context) {
	log.Println(c.Cookie("jwt"))
	c.SetCookie("jwt", "", 3600, "/", "localhost", true, true)
	msg := "ログアウトしました。"
	c.JSON(http.StatusOK, gin.H{"msg": msg})
}