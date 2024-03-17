package repositories

import (
	"context"
	"github.com/alireza-gholami0/go-gin-api/src/models"
	"gorm.io/gorm"
)

type userRepository struct {
	database   gorm.DB
	collection string
}

func NewUserRepository(db gorm.DB, collection string) models.UserRepository {
	return &userRepository{
		database:   db,
		collection: collection,
	}
}

func (u userRepository) Create(c context.Context, user *models.User) error {
	//TODO implement me
	panic("implement me")
}

func (u userRepository) Fetch(c context.Context) ([]models.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u userRepository) GetByEmail(c context.Context, email string) (models.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u userRepository) GetByID(c context.Context, id string) (models.User, error) {
	//TODO implement me
	panic("implement me")
}
