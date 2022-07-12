package routes

import (
	"time"
	"github.com/raylicola/NFlaquiz/controllers"
	"github.com/gin-gonic/gin"
  "github.com/gin-contrib/cors"
)

func GetRouter() *gin.Engine {
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		// 許可するアクセス元
		AllowOrigins: []string{
				"http://localhost:8080",
		},
		// アクセス許可するHTTPメソッド
		AllowMethods: []string{
				"POST",
				"GET",
				"PUT",
		},
		// 許可するHTTPリクエストヘッダ
		AllowHeaders: []string{
				"Content-Type",
		},
		// cookieなどの情報を必要とする
		AllowCredentials: true,
		// preflightリクエストの結果をキャッシュする時間
		MaxAge: 24 * time.Hour,
}))

	router.POST("/signup", controllers.Signup)
	router.POST("/login", controllers.Login)
	router.GET("/logout", controllers.Logout)
	router.GET("/map", controllers.GetMapInfo)
	router.PUT("/bookmark/:country_id", controllers.UpdateBookmark)
	router.POST("/result/update", controllers.UpdateResult)
	router.GET("/quiz/select", controllers.SelectQuiz)

	return router
}