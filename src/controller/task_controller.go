package controller

import (
	"github.com/alireza-gholami0/go-gin-api/src/models"
	"github.com/alireza-gholami0/go-gin-api/src/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type TaskController struct {
	TaskService services.TaskService
}

func (tc *TaskController) AddTask(c *gin.Context) {
	var request models.AddTaskRequest
	err := c.ShouldBind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Message: err.Error()})
		return
	}

	userIdPtr := getUserID(c)
	task, err := tc.TaskService.AddTask(c, request, *userIdPtr)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Message: err.Error()})
		return
	}
	c.JSON(200, task)
}

func (tc *TaskController) GetTask(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Message: "id not exist"})
		return
	}
	userIdPtr := getUserID(c)
	taskId, _ := strconv.Atoi(id)
	task, err := tc.TaskService.GetTask(c, uint(taskId), *userIdPtr)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, task)

}

func getUserID(c *gin.Context) *uint {
	userId, ok := c.Get("x-user-id")
	if !ok {
		c.JSON(http.StatusUnauthorized, models.ErrorResponse{Message: "user not found"})
		return nil
	}
	userIdPtr, ok := userId.(*uint)
	return userIdPtr
}
