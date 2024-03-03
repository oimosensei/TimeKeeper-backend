package model

import (
	"time"
)

// Attendance 構造体は出退勤記録を表します。
type Attendance struct {
	ID           uint       `gorm:"primaryKey;autoIncrement" json:"id"`                   // レコードのユニークな識別子
	UserID       uint       `gorm:"not null;index" json:"user_id"`                        // ユーザー識別子（外部キー）
	ClockInTime  *time.Time `json:"clock_in_time"`                                        // 出勤時刻
	ClockOutTime *time.Time `json:"clock_out_time"`                                       // 退勤時刻
	Date         string     `gorm:"not null" json:"date"`                                 // 日付
	Status       string     `gorm:"type:varchar(255);not null" json:"status"`             // 出勤状態（通常、遅刻、早退）
	CreatedAt    time.Time  `gorm:"not null;default:CURRENT_TIMESTAMP" json:"created_at"` // レコード作成時刻
	UpdatedAt    time.Time  `gorm:"not null;default:CURRENT_TIMESTAMP" json:"updated_at"` // レコード更新時刻
}
