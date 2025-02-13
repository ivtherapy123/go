package taskService

import "gorm.io/gorm"

type Task struct {
	Id     uint   `json:"id" gorm:"primaryKey"`
	Task   string `json:"task"`
	IsDone bool   `json:"is_done"`
	UserID uint   `json:"user_id"`
	gorm.Model
}
