package models

import "context"

const (
	CollectionUser = "users"
)

type User struct {
	ID       uint   `gorm:"primaryKey;autoIncrement;column:id"`
	Username string `gorm:"size:255;not null"`
	Email    string `gorm:"size:255"`
	Password string `gorm:"size:255;not null"`
	Tasks    []Task `gorm:"foreignKey:user_id"`
}

type UserRepository interface {
	Create(c context.Context, user *User) error
	Fetch(c context.Context) ([]User, error)
	GetByEmail(c context.Context, email string) (User, error)
	GetByID(c context.Context, id string) (User, error)
}
