package model

import (
	"time"
)

type Task struct {
	ID             uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID         uint      `gorm:"not null;index" json:"user_id"`
	Title          string    `gorm:"type:varchar(255);not null" json:"title"`
	Description    string    `gorm:"type:text;not null" json:"description"`
	URL            string    `gorm:"type:varchar(255);not null" json:"url"`
	OTGImageURL    string    `gorm:"type:varchar(255)" json:"ogp_image_url"`
	OTGDescription string    `gorm:"type:text" json:"ogp_description"`
	IsDone         bool      `gorm:"not null;default:false" json:"is_done"`
	CreatedAt      time.Time `gorm:"not null;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt      time.Time `gorm:"not null;default:CURRENT_TIMESTAMP" json:"updated_at"`
}

type TaskResponse struct {
	ID             uint   `json:"id" gorm:"primary_key" `
	UserID         uint   `json:"user_id" gorm:"not null;index"`
	Title          string `json:"title" gorm:"unique"`
	Description    string `json:"description" gorm:"unique"`
	URL            string `json:"url" gorm:"unique"`
	OTGImageURL    string `json:"ogp_image_url" gorm:"unique"`
	OTGDescription string `json:"ogp_description" gorm:"unique"`
	IsDone         bool   `json:"is_done" gorm:"unique"`
}
