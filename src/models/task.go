package models

import "time"

const (
	CollectionTask = "task"
)

type Task struct {
	ID     uint      `gorm:"primaryKey;autoIncrement;column:id"`
	UserID uint      `gorm:"column:user_id;not null"`
	Title  string    `gorm:"size:255;not null"`
	Time   time.Time `gorm:"type:date;default:current_date"`
}
