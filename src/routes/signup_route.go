package routes

import (
	"github.com/alireza-gholami0/go-gin-api/src/controller"
	"github.com/alireza-gholami0/go-gin-api/src/models"
	"github.com/alireza-gholami0/go-gin-api/src/repositories"
	"github.com/alireza-gholami0/go-gin-api/src/services"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"time"
)

func NewSignupRouter(timeout time.Duration, db gorm.DB, group *gin.RouterGroup) {
	ur := repositories.NewUserRepository(db, models.CollectionUser)
	us := services.NewUserService(ur, timeout)
	sc := controller.SignupController{
		UserService: us,
	}
	group.POST("/signup", sc.Signup)
}
