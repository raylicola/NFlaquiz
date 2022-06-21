package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/raylicola/NFlaquiz/database"
	"github.com/raylicola/NFlaquiz/models"
)

// トップページ地図情報を取得
// 返り値:
//   成功時：[country_id,name,description, bookmark, weight]の配列
func GetMapInfo(c *gin.Context) {
	user, err := User(c)
	var map_info []models.MapInfo
	if err != nil {
		// ログインしていない場合
		database.DB.Table("countries").Scan(&map_info)
	} else {
		// ログイン済みの場合
		database.DB.Table("countries").Select("countries.id, countries.name, countries.description, results.bookmark, results.weight").Joins("left outer join results on countries.id = results.country_id").Where("results.user_id=?", user.ID).Or("results.user_id IS NULL").Scan(&map_info)
	}

	c.JSON(http.StatusOK, gin.H{"map_info": map_info})
}