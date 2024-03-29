package repositories

import (
	"fmt"
	"github.com/alireza-gholami0/go-gin-api/src/models"
	"gorm.io/gorm"
)

type taskRepository struct {
	database gorm.DB
}

type TaskRepository interface {
	Create(task models.Task) (models.Task, error)
	GetByID(id uint) (models.Task, error)
}

var tr *taskRepository

func GetTaskRepositoryIns(db gorm.DB) TaskRepository {
	if tr == nil {
		tr = &taskRepository{
			database: db,
		}
	}
	return tr
}

func (t taskRepository) Create(task models.Task) (models.Task, error) {
	err := t.database.Create(&task).Error
	if err != nil {
		fmt.Println(err.Error())
	}
	return task, err
}

func (t taskRepository) GetByID(id uint) (models.Task, error) {
	var task models.Task
	err := t.database.Where("id = ?", id).First(&task).Error
	return task, err
}
