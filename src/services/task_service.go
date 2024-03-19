package services

import (
	"github.com/alireza-gholami0/go-gin-api/src/models"
	"github.com/alireza-gholami0/go-gin-api/src/repositories"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type taskService struct {
	taskRepository repositories.TaskRepository
	contextTimeout time.Duration
}

type TaskService interface {
	AddTask(c *gin.Context, request models.AddTaskRequest, userId uint) (models.Task, error)
	GetTask(c *gin.Context, taskId uint) (*models.Task, error)
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
		User:   userId,
		Title:  request.Title,
		Date:   request.Date,
		Status: models.Undefined,
	}
	task, err := t.taskRepository.Create(task)
	return task, err
}

func (t taskService) GetTask(c *gin.Context, taskId uint) (*models.Task, error) {
	task, err := t.taskRepository.GetByID(taskId)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Message: "Task not found"})
		return nil, err
	}
	return &task, nil
}
