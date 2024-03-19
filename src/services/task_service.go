package services

import (
	"github.com/alireza-gholami0/go-gin-api/src/models"
	"github.com/alireza-gholami0/go-gin-api/src/repositories"
	"github.com/gin-gonic/gin"
	"time"
)

type taskService struct {
	taskRepository repositories.TaskRepository
	contextTimeout time.Duration
}

type TaskService interface {
	AddTask(c *gin.Context, request models.AddTaskRequest, userId uint) (models.Task, error)
}

var ts *taskService

func GetTaskServiceIns(taskRepository repositories.TaskRepository, timeout time.Duration) TaskService {
	if ts == nil {
		ts = &taskService{
			taskRepository: taskRepository,
			contextTimeout: timeout,
		}
	}
	return ts
}

func (t taskService) AddTask(c *gin.Context, request models.AddTaskRequest, userId uint) (models.Task, error) {
	task := models.Task{
		User:  userId,
		Title: request.Title,
		Date:  request.Date,
	}
	task, err := t.taskRepository.Create(task)
	return task, err
}
