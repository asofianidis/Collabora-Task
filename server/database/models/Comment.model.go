package model

import "time"

type Comment struct {
	ID        uint      `json:"id" sql:"id" gorm:"autoIncrement"`
	CreatedAt time.Time `json:"createdAt" sql:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updatedAt" sql:"updatedAt" gorm:"autoUpdateTime"`
	CreatedBy uint      `json:"createdBy" sql:"createdBy"`
	TaskID    uint      `json:"taskID" sql:"taskID"`
	Text      string    `json:"text" sql:"text"`
	UserID    uint      `json:"userID" sql:"userID"`
}
