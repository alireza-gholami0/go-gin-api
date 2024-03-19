package services

import (
	"github.com/alireza-gholami0/go-gin-api/src/models"
	"github.com/alireza-gholami0/go-gin-api/src/repositories"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"time"
)

type userService struct {
	userRepository repositories.UserRepository
	contextTimeout time.Duration
}

type UserService interface {
	CreateUser(c *gin.Context, request models.SignupRequest) (*models.SignupResponse, error)
	Login(c *gin.Context, request models.LoginRequest) (*models.LoginResponse, error)
}

var us *userService

func GetUSerServiceIns(userRepository repositories.UserRepository, timeout time.Duration) UserService {
	if us == nil {
		us = &userService{
			userRepository: userRepository,
			contextTimeout: timeout,
		}
	}
	return us
}

func (u userService) CreateUser(c *gin.Context, request models.SignupRequest) (*models.SignupResponse, error) {
	var user models.User
	var response models.SignupResponse
	user, err := u.userRepository.GetByEmail(request.Email)
	if err == nil {
		c.JSON(http.StatusConflict, models.ErrorResponse{Message: "User already exists with the given email"})
		return nil, err
	}
	encryptedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(request.Password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Message: err.Error()})
		return nil, err
	}

	request.Password = string(encryptedPassword)

	user = models.User{
		Name:     request.Name,
		Email:    request.Email,
		Password: request.Password,
	}

	user, err = u.userRepository.Create(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Message: err.Error()})
		return nil, err
	}
	accessToken, err := CreateToken(user.ID, 5)
	response = models.SignupResponse{
		AccessToken: accessToken,
	}

	return &response, nil
}

func (u userService) Login(c *gin.Context, request models.LoginRequest) (*models.LoginResponse, error) {
	var response models.LoginResponse
	var user models.User
	var err error

	user, err = u.userRepository.GetByEmail(request.Email)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Message: "User not found"})
		return nil, err
	}

	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password)) != nil {
		c.JSON(http.StatusUnauthorized, models.ErrorResponse{Message: "Invalid credentials"})
		return nil, err
	}

	accessToken, err := CreateToken(user.ID, 5)
	response = models.LoginResponse{
		AccessToken: accessToken,
	}

	return &response, nil
}
