package models

const (
	CollectionUser = "users"
)

type User struct {
	ID       uint   `gorm:"primaryKey;autoIncrement;column:id"`
	Name     string `gorm:"size:255;not null"`
	Email    string `gorm:"size:255"`
	Password string `gorm:"size:255;not null"`
	Tasks    []Task `gorm:"foreignKey:user_id"`
}
