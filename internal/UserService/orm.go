package UserService

import (
	"awesomeProject1/internal/taskService"
	"gorm.io/gorm"
)

type User struct {
	Id       uint               `json:"id" gorm:"primaryKey"`
	Email    string             `json:"email"`
	Password string             `json:"password"`
	Tasks    []taskService.Task `json:"tasks"`
	gorm.Model
}
