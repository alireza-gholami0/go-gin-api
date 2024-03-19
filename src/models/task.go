package models

import "time"

type TaskStatus int

const (
	Active         TaskStatus = 1
	Inactive       TaskStatus = 2
	Undefined      TaskStatus = 0
	CollectionTask            = "task"
)

type Task struct {
	ID     uint       `gorm:"primaryKey;autoIncrement;column:id"`
	User   uint       `gorm:"column:user_id"`
	Title  string     `gorm:"size:255;not null"`
	Date   time.Time  `gorm:"type:date;default:current_date"`
	Status TaskStatus `gorm:"default:0;check:status < 3;check:status >= 0"`
}

type AddTaskRequest struct {
	Title string    `form:"title" binding:"required"`
	Date  time.Time `form:"date"`
}
