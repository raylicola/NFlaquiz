package controllers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/raylicola/NFlaquiz/database"
	"github.com/raylicola/NFlaquiz/models"
	"github.com/raylicola/NFlaquiz/utils"
)

// 検索条件に当てはまるクイズを10問選ぶ
// ヒット数が10個以下なら全て選択する
// 受信：
//   colors(array): 選択された色
//    areas(array): 選択された地域
//  bookmark(int): ブックマークで絞り込むか否か
//                 0:絞り込みを行わない, 1:行う
// 戻り値：
//    10問以下のクイズセット
func SelectQuiz(c *gin.Context) {

	// クエリパラメータをバインド
	var req models.QuizFilter
	if err := c.Bind(&req); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	colors := req.Colors
	areas := req.Areas
	bookmark, _ := strconv.Atoi(req.Bookmark)
	user, err := utils.AuthUser(c)
	log.Println(bookmark)

	// クイズを選択
	// TODO: 可能であれば色はAND検索に修正
	var quizzes []models.Quiz

	if err != nil {
    // 1. ログインしていない場合
		// 条件に合うものをランダムに選択
		database.DB.Distinct("quizzes.id, hiragana, quizzes.country_id, hint1, hint2, hint3").Joins("left join countries on quizzes.country_id = countries.id").Joins("left join flag_colors on quizzes.country_id = flag_colors.country_id").Where("color_id in (?)", colors).Where("area_id in (?)", areas).Order("rand()").Limit(10).Find(&quizzes)

	} else {

		// 2. ログイン済みの場合
		// 2-1.ブックマークで絞る場合
		// 正答率が低い順に10問選択
		if bookmark == 1 {

			database.DB.Distinct("quizzes.id, hiragana, quizzes.country_id, hint1, hint2, hint3, weight").Joins("left join countries on quizzes.country_id = countries.id").Joins("left join flag_colors on quizzes.country_id = flag_colors.country_id").Joins("left join results on results.country_id = quizzes.country_id").Where("user_id=?", user.ID).Where("bookmark=1").Where("color_id in (?)", colors).Where("area_id in (?)", areas).Order("weight").Limit(10).Find(&quizzes)

		} else {

			// 2-2. ブックマークで絞らない場合
			var quizzes1 []models.Quiz
			var quizzes2 []models.Quiz
			// (1)該当ユーザーの正答率が低いクイズ
			// weight <= 0.5 と定義
			query1 := database.DB.Distinct("quizzes.id, hiragana, quizzes.country_id, hint1, hint2, hint3").Joins("left join results on quizzes.country_id = results.country_id").Joins("left join countries on quizzes.country_id = countries.id").Joins("left join flag_colors on quizzes.country_id = flag_colors.country_id").Where("user_id=?", user.ID).Where("weight<=0.5").Where("color_id in (?)", colors).Where("area_id in (?)", areas).Find(&quizzes1)

			// (2)該当ユーザーが未回答のクイズ
			query2 := database.DB.Distinct("quizzes.id, hiragana, quizzes.country_id, hint1, hint2, hint3").Joins("left join results on quizzes.country_id = results.country_id").Joins("left join countries on quizzes.country_id = countries.id").Joins("left join flag_colors on quizzes.country_id = flag_colors.country_id").Where("quizzes.country_id not in (?)", database.DB.Table("results").Select("country_id").Where("user_id=?", user.ID)).Where("color_id in (?)", colors).Where("area_id in (?)", areas).Find(&quizzes2)

			// (1)と(2)の中からランダムに選択
			database.DB.Raw(
				"? UNION ? ORDER BY rand() LIMIT 10",
				query1,
				query2,
			).Scan(&quizzes)

			// 10問に満たない場合は残りをweight>0.5のものから昇順で選択
			if len(quizzes) < 10 {
				var quizzes3 []models.Quiz
				query1 = database.DB.Distinct("quizzes.id, hiragana, quizzes.country_id, hint1, hint2, hint3, weight").Joins("left join results on quizzes.country_id = results.country_id").Joins("left join countries on quizzes.country_id = countries.id").Joins("left join flag_colors on quizzes.country_id = flag_colors.country_id").Where("user_id=?", user.ID).Where("weight>0.5").Where("color_id in (?)", colors).Where("area_id in (?)", areas).Order("weight").Limit(10-len(quizzes)).Find(&quizzes3)

				quizzes = append(quizzes, quizzes3...)
			}
		}
	}

	c.JSON(http.StatusOK, gin.H{"quizzes": quizzes})
}