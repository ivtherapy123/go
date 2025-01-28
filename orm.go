package main

import "gorm.io/gorm"

type Message struct {
	gorm.Model
	ID     uint   `json:"id" gorm:"primaryKey"`
	Task   string `json:"task"`
	IsDone bool   `json:"is_done"`
}
