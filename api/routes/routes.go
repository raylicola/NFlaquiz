package routes

import (
	"github.com/raylicola/NFlaquiz/controllers"
	"github.com/gin-gonic/gin"
)

func GetRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/signup", controllers.GetSignup)
	router.POST("/signup", controllers.PostSignup)
	router.GET("/login", controllers.GetLogin)
	router.POST("/login", controllers.PostLogin)
	router.GET("/logout", controllers.Logout)
	router.GET("/auth", controllers.Auth)

	return router
}