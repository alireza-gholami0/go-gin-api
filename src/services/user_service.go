package services

import (
	"context"
	"github.com/alireza-gholami0/go-gin-api/src/models"
	"time"
)

type userService struct {
	userRepository models.UserRepository
	contextTimeout time.Duration
}

type UserService interface {
	createUser(c context.Context, user *models.User) (models.User, error)
}

func NewUserService(userRepository models.UserRepository, timeout time.Duration) UserService {
	return &userService{
		userRepository: userRepository,
		contextTimeout: timeout,
	}
}

func (u userService) createUser(c context.Context, user *models.User) (models.User, error) {
	//TODO implement me
	panic("implement me")
}
