package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/raylicola/NFlaquiz/database"
	"github.com/raylicola/NFlaquiz/models"
	"github.com/raylicola/NFlaquiz/utils"
)

// ブックマーク有無の更新
// 受信：
//   id: URLで指定 (/bookmark/:country_id)
func UpdateBookmark(c *gin.Context) {
	var result models.Result
	country_id := c.Param("country_id")
	user, err := utils.AuthUser(c)

	// ユーザーが認証されていない場合
	if err != nil {
    c.JSON(http.StatusUnauthorized, gin.H{"err_msg": "認証されていません"})
	}

	database.DB.Table("results").Where("country_id=?",country_id).Where("user_id=?",user.ID).Find(&result)

	if result.Bookmark == 1 {
		result.Bookmark = 0
	} else {
		result.Bookmark = 1
	}
	database.DB.Save(&result)

	c.JSON(http.StatusOK, gin.H{"msg": "Change Succeeded"})
}