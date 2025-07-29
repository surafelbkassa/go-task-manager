package Usecases

import (
	"errors"

	domain "github.com/surafelbkassa/go-task-manager/Domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TaskUseCaseInterface interface {
	GetTasks() ([]domain.Task, error)
	GetTaskByID(string) (*domain.Task, error)
	CreateTask(domain.Task) (*domain.Task, error)
	UpdateTask(string, domain.Task) (*domain.Task, error)
	DeleteTask(string) error
}

type TaskUseCase struct {
	repo domain.TaskRepository
}

func NewTaskUseCase(r domain.TaskRepository) *TaskUseCase {
	return &TaskUseCase{repo: r}
}

// ← ADD THIS
func (u *TaskUseCase) GetTasks() ([]domain.Task, error) {
	return u.repo.GetAll()
}

func (u *TaskUseCase) GetTaskByID(id string) (*domain.Task, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	return u.repo.GetByID(objID)
}

func (u *TaskUseCase) CreateTask(task domain.Task) (*domain.Task, error) {
	return u.repo.Create(task)
}

func (u *TaskUseCase) UpdateTask(id string, task domain.Task) (*domain.Task, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New("invalid ID format")
	}
	return u.repo.Update(objID, task)
}

func (u *TaskUseCase) DeleteTask(id string) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.New("invalid ID format")
	}
	return u.repo.Delete(objID)
}
