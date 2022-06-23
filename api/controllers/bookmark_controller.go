package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/raylicola/NFlaquiz/database"
	"github.com/raylicola/NFlaquiz/models"
)

// ブックマーク有無の更新
// 受信：
//   id: URLで指定 (/bookmark/:id)
func UpdateBookmark(c *gin.Context) {
	var result models.Result
	id := c.Param("id")

	database.DB.Table("results").Where("id=?",id).Find(&result)

	if result.Bookmark == 1 {
		result.Bookmark = 0
	} else {
		result.Bookmark = 1
	}
	database.DB.Save(&result)

	c.JSON(http.StatusOK, gin.H{"msg": "Change Succeeded"})
}