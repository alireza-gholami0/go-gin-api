package routes

import (
	"github.com/alireza-gholami0/go-gin-api/src/controller"
	"github.com/alireza-gholami0/go-gin-api/src/repositories"
	"github.com/alireza-gholami0/go-gin-api/src/services"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"time"
)

func NewLoginRouter(timeout time.Duration, db gorm.DB, group *gin.RouterGroup) {
	ur := repositories.GetUserRepositoryIns(db)
	us := services.GetUSerServiceIns(ur, timeout)
	lc := controller.LoginController{
		UserService: us,
	}
	group.POST("/login", lc.Login)
}
