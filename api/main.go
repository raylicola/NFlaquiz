package main

import (
  "log"
  "github.com/raylicola/NFlaquiz/database"
  "github.com/gin-gonic/gin"
)

func main() {

  database.Connect()
  router := gin.Default()
  router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"key": "value"})
	})

  if err := router.Run(":8888"); err != nil {
    log.Fatal("Server Run Failed.: ", err)
  }
}