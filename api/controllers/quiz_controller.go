package controllers

import (
	"errors"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/raylicola/NFlaquiz/database"
	"github.com/raylicola/NFlaquiz/models"
	"github.com/raylicola/NFlaquiz/utils"
	"gorm.io/gorm"
)

// クイズの回答状況をもとにResultを更新する
// ログインしているときのみ
// 受信：
//   [{country_id: 国ID, answer: (0|1), bookmark: (0|1)}, ...]
//     answer -> 1:正解, 0: 不正解
//   bookmark -> 1:登録する, 0:しない（既に登録済みであれば変更しない）
func UpdateResult(c *gin.Context) {

	// データをバインド
	var req []models.AnswerStatus
	err := c.BindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err_msg": "データのバインドに失敗しました"})
		return
	}

	// ユーザーが認証されていない場合
	user, err := utils.AuthUser(c)
	if err != nil {
    c.JSON(http.StatusUnauthorized, gin.H{"err_msg": "認証されていません"})
		return
	}

	// resultテーブル更新処理
	for _, v := range req {

		// クイズ結果が登録済みか否か判定
		var result models.Result
		query := database.DB.Where("country_id=?", v.CountryID).Where("user_id=?", user.ID).First(&result)

		if errors.Is(query.Error, gorm.ErrRecordNotFound) {

			// 未登録の場合は新規作成
			weight := 1.0
			// 不正解の場合
			if v.Answer == 0 {
				weight = 0.5
			}
			new_result := models.Result{
				CountryID: v.CountryID,
				UserID: user.ID,
				Weight: weight,
				Bookmark: v.Bookmark,
			}
			log.Println(new_result)
			database.DB.Create(&new_result)

		} else {
			// 登録済みの場合はレコードを更新
			// 重みの更新
			weight := result.Weight
			if v.Answer == 1 && result.Weight < 1.0{
				weight += 0.25
			} else if v.Answer == 0 && result.Weight > 0.25 {
				weight -= 0.25
			}

			database.DB.Model(&result).Where("user_id=?", user.ID).Where("country_id=?", v.CountryID).Update("weight", weight)

			// ブックマークの更新
			if v.Bookmark == 1 {
				database.DB.Model(&result).Where("user_id=?", user.ID).Where("country_id=?", v.CountryID).Update("bookmark", 1)
			}
		}
	}

	c.JSON(http.StatusOK, gin.H{})
}

