package models

import "time"

const (
	CollectionTask = "task"
)

type Task struct {
	ID    uint      `gorm:"primaryKey;autoIncrement;column:id"`
	User  uint      `gorm:"column:user_id"`
	Title string    `gorm:"size:255;not null"`
	Date  time.Time `gorm:"type:date;default:current_date"`
}

type AddTaskRequest struct {
	Title string    `form:"title" binding:"required"`
	Date  time.Time `form:"date"`
}
