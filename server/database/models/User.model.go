package model

type User struct {
	ID            uint        `json:"id" sql:"id" gorm:"autoIncrement"`
	Email         string      `json:"email" sql:"email"`
	Username      string      `json:"username" sql:"username"`
	Password      string      `json:"password" sql:"password"`
	CreatedAt     string      `json:"createdAt" sql:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt     string      `json:"updatedAt" sql:"updatedAt" gorm:"autoUpdateTime"`
	Admin         bool        `json:"admin" sql:"admin" default:"false"`
	AssignedTasks []Task      `json:"assignedTasks" sql:"assignedTasks" gorm:"foreignKey:AssigneeID;"`
	CreatedTasks  []Task      `json:"createdTasks" sql:"createdTasks" gorm:"foreignKey:CreatorID;"`
	Comments      []Comment   `json:"comments" sql:"comments"`
	WorkSpaces    []Workspace `gorm:"many2many:user_workspace;" json:"workspaces" sql:"workspaces"`
	TokenHash     string      `json:"tokenHash" sql:"tokenHash"`
}
