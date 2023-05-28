package model

import "time"

type Task struct {
	ID          uint      `json:"id" sql:"id"`
	CreatorID   uint      `json:"creatorid" sql:"creatorid"`
	CreatedAt   time.Time `json:"createdat" sql:"createdat"`
	UpdatedAt   time.Time `json:"updatedat" sql:"updatedat"`
	Title       string    `json:"title" sql:"title"`
	AssigneeID  uint      `json:"assigneeid" sql:"assigneeid"`
	Description string    `json:"description" sql:"description"`
	Status      string    `json:"status" sql:"status" default:"To Do"`
	WorkspaceID uint      `json:"workspaceid" sql:"workspaceid"`
	Comments    []Comment `json:"comments" sql:"comments"`
}
