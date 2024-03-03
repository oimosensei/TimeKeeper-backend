package main

import (
	"TverClip/db"
	"TverClip/model"
	"time"
)

func main() {
	db := db.NewDB()

	// データベースから全ユーザーを取得
	var users []model.User
	db.Find(&users)

	// 指定月の日数を取得（例：2024年2月）
	year, month := 2024, time.March
	daysInMonth := daysInMonth(year, month)

	// 各ユーザーに対して、指定月の全日にわたる出退勤記録を生成し、データベースに保存
	for _, user := range users {
		for day := 1; day <= daysInMonth; day++ {
			date := time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
			attendance := model.Attendance{
				UserID: user.ID,
				Date:   date.Format("2006-01-02"),
				Status: "通常", // 適宜設定
			}
			db.Create(&attendance)
		}
	}
}

// 指定された年と月の日数を計算
func daysInMonth(year int, month time.Month) int {
	return time.Date(year, month+1, 0, 0, 0, 0, 0, time.UTC).Day()
}
