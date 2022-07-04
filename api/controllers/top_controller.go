package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/raylicola/NFlaquiz/database"
	"github.com/raylicola/NFlaquiz/models"
	"github.com/raylicola/NFlaquiz/utils"
)

// トップページ地図情報を取得
// 返り値:
//   成功時：[country_id,name,description, bookmark, weight]の配列
func GetMapInfo(c *gin.Context) {
	user, err := utils.AuthUser(c)
	var mapInfo []models.MapInfo
	if err != nil {
		// ログインしていない場合
		database.DB.Table("countries").Scan(&mapInfo)
	} else {
		// ログイン済みの場合
		var mapInfo1 []models.MapInfo
		var mapInfo2 []models.MapInfo

		// 該当ユーザーが結果を持つ国
		query1 := database.DB.Table("countries").Select("countries.id, countries.name, countries.description, results.bookmark, results.weight").Joins("left outer join results on countries.id = results.country_id").Where("results.user_id=?", user.ID).Scan(&mapInfo1)

		// 結果を持たない国
		// Bookmark, Weightに0を代入
		query2 := database.DB.Table("countries").Select("countries.id, countries.name, countries.description, 0, 0").Joins("left outer join results on countries.id = results.country_id").Where("countries.id not in (?)", database.DB.Table("results").Select("country_id").Where("user_id=?", user.ID)).Scan(&mapInfo2)

		log.Println("---")
		log.Println(mapInfo2)

		database.DB.Raw(
			"? UNION ?",
			query1,
			query2,
		).Scan(&mapInfo)
	}

	c.JSON(http.StatusOK, gin.H{"map_info": mapInfo})
}