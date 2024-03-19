package routes

import (
	"github.com/alireza-gholami0/go-gin-api/src/controller"
	"github.com/alireza-gholami0/go-gin-api/src/repositories"
	"github.com/alireza-gholami0/go-gin-api/src/services"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"time"
)

func NewTaskRouter(timeout time.Duration, db gorm.DB, group *gin.RouterGroup) {
	tr := repositories.GetTaskRepositoryIns(db)
	ts := services.GetTaskServiceIns(tr, timeout)
	tc := controller.TaskController{
		TaskService: ts,
	}
	group.POST("/task/add", tc.AddTask)
}
