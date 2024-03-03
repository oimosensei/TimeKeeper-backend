package repository

import (
	"TverClip/model"
	"fmt"

	"gorm.io/gorm"
)

type ITaskRepository interface {
	GetAllTasks(tasks *[]model.Task) error
	GetTasksByUserID(tasks *[]model.Task, userID uint) error
	GetTaskByID(task *model.Task, id uint) error
	CreateTask(task *model.Task) error
	UpdateTask(taskID uint, userID uint, newTask *model.Task) error
	DeleteTask(taskID, userID uint) error
}

type TaskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) ITaskRepository {
	return &TaskRepository{db}
}

func (tr *TaskRepository) GetAllTasks(tasks *[]model.Task) error {
	if err := tr.db.Find(tasks).Error; err != nil {
		return err
	}
	return nil
}

func (tr *TaskRepository) GetTasksByUserID(tasks *[]model.Task, userID uint) error {
	if err := tr.db.Where("user_id = ?", userID).Find(tasks).Error; err != nil {
		return err
	}
	return nil
}

func (tr *TaskRepository) GetTaskByID(task *model.Task, id uint) error {
	if err := tr.db.Where("id = ?", id).First(task).Error; err != nil {
		return err
	}
	return nil
}

func (tr *TaskRepository) CreateTask(task *model.Task) error {
	if err := tr.db.Create(task).Error; err != nil {
		return err
	}
	return nil
}

func (tr *TaskRepository) UpdateTask(taskID uint, userID uint, newTask *model.Task) error {
	result := tr.db.Model(&model.Task{}).Where("id = ? AND user_id = ?", taskID, userID).Updates(newTask)

	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected < 1 {
		return fmt.Errorf("object does not exist")
	}

	return nil
}

func (tr *TaskRepository) DeleteTask(taskID uint, userID uint) error {
	result := tr.db.Where("id = ? AND user_id = ?", taskID, userID).Delete(&model.Task{})
	fmt.Println(taskID, userID)

	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected < 1 {
		return fmt.Errorf("object does not exist")
	}

	return nil
}
