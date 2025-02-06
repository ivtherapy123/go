package taskService

import (
	"gorm.io/gorm"
	"strconv"
)

type MessageRepository interface {
	// CreateTask - Передаем в функцию task типа Task из orm.go
	// возвращаем созданный Task и ошибку
	CreateTask(task Task) (Task, error)
	// GetAllTasks - Возвращаем массив из всех задач в БД и ошибку
	GetAllTasks() ([]Task, error)
	// UpdateTaskByID - Передаем id и Task, возвращаем обновленный Task
	// и ошибку
	UpdateTaskByID(id uint, task Task) (Task, error)
	// DeleteTaskByID - Передаем id для удаления, возвращаем только ошибку
	DeleteTaskByID(id string) error
}

type taskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) *taskRepository {
	return &taskRepository{db: db}
}

func (r *taskRepository) CreateTask(task Task) (Task, error) {
	result := r.db.Create(&task)
	if result.Error != nil {
		return Task{}, result.Error
	}
	return task, nil
}

func (r *taskRepository) GetAllTasks() ([]Task, error) {
	var tasks []Task
	err := r.db.Find(&tasks).Error
	return tasks, err
}

func (r *taskRepository) UpdateTaskByID(id uint, task Task) (Task, error) {

	var existingTask Task
	r.db.First(&existingTask, id)

	existingTask.Task = task.Task
	existingTask.IsDone = task.IsDone

	r.db.Save(&existingTask)
	return existingTask, nil
}

func (r *taskRepository) DeleteTaskByID(id string) error {
	var task Task
	uintId, err := strconv.ParseUint(id, 10, 32) // Преобразуем в uint
	if err != nil {
		return err
	}
	r.db.First(&task, uintId)
	r.db.Delete(&task)
	return nil
}
