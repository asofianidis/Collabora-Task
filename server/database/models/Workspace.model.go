package model

import "time"

type Workspace struct {
	ID          uint      `json:"id" sql:"id" gorm:"autoIncrement"`
	CreatedAt   time.Time `json:"createdAt" sql:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt   time.Time `json:"updatedAt" sql:"updatedAt" gorm:"autoUpdateTime"`
	Members     []User    `json:"members" sql:"members" gorm:"many2many:user_workspace;"`
	Tasks       []Task    `json:"tasks" sql:"tasks"`
	Name        string    `json:"name" sql:"name"`
	Description string    `json:"description" sql:"description"`
}
