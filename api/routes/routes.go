package routes

import (
	"github.com/raylicola/NFlaquiz/controllers"
	"github.com/gin-gonic/gin"
)

func GetRouter() *gin.Engine {
	router := gin.Default()

	router.POST("/signup", controllers.Signup)
	router.POST("/login", controllers.Login)
	router.GET("/logout", controllers.Logout)

	return router
}