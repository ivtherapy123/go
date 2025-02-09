package UserService

import "gorm.io/gorm"

type User struct {
	Id       uint   `json:"id" gorm:"primaryKey"`
	Email    string `json:"email"`
	Password string `json:"password"`
	gorm.Model
}
