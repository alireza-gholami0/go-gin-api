package repositories

import (
	"github.com/alireza-gholami0/go-gin-api/src/models"
	"gorm.io/gorm"
)

type userRepository struct {
	database gorm.DB
}

type UserRepository interface {
	Create(user models.User) (models.User, error)
	GetByEmail(email string) (models.User, error)
	GetByID(id uint) (models.User, error)
}

var ur *userRepository

func GetUserRepositoryIns(db gorm.DB) UserRepository {
	if ur == nil {
		ur = &userRepository{
			database: db,
		}
	}
	return ur
}

func (ur userRepository) Create(user models.User) (models.User, error) {
	err := ur.database.Create(&user).Error
	return user, err
}

func (ur userRepository) Fetch() ([]models.User, error) {
	//TODO implement me
	panic("implement me")
}

func (ur userRepository) GetByEmail(email string) (models.User, error) {
	var user models.User
	err := ur.database.Where("email = ?", email).First(&user).Error
	return user, err
}

func (ur userRepository) GetByID(id uint) (models.User, error) {
	var user models.User
	err := ur.database.Where("id = ?", id).First(&user).Error
	return user, err
}
