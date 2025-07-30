package Usecases

import (
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	Domain "github.com/surafelbkassa/go-task-manager/Domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// --- Mock TaskRepository ---
type MockTaskRepo struct {
	mock.Mock
}

func (m *MockTaskRepo) GetAll() ([]Domain.Task, error) {
	args := m.Called()
	return args.Get(0).([]Domain.Task), args.Error(1)
}

func (m *MockTaskRepo) GetByID(id primitive.ObjectID) (*Domain.Task, error) {
	args := m.Called(id)
	return args.Get(0).(*Domain.Task), args.Error(1)
}

func (m *MockTaskRepo) Create(task Domain.Task) (*Domain.Task, error) {
	args := m.Called(task)
	return args.Get(0).(*Domain.Task), args.Error(1)
}

func (m *MockTaskRepo) Update(id primitive.ObjectID, task Domain.Task) (*Domain.Task, error) {
	args := m.Called(id, task)
	return args.Get(0).(*Domain.Task), args.Error(1)
}

func (m *MockTaskRepo) Delete(id primitive.ObjectID) error {
	return m.Called(id).Error(0)
}

// ---------- TESTS ----------

func TestGetTasks_Success(t *testing.T) {
	mockRepo := new(MockTaskRepo)
	uc := NewTaskUseCase(mockRepo)

	expected := []Domain.Task{
		{Title: "A", Description: "desc", CreatedAt: time.Now()},
	}
	mockRepo.On("GetAll").Return(expected, nil)

	result, err := uc.GetTasks()
	assert.NoError(t, err)
	assert.Equal(t, expected, result)
	mockRepo.AssertExpectations(t)
}

func TestGetTasks_Error(t *testing.T) {
	mockRepo := new(MockTaskRepo)
	uc := NewTaskUseCase(mockRepo)

	mockRepo.On("GetAll").Return([]Domain.Task(nil), errors.New("db fail"))

	_, err := uc.GetTasks()
	assert.EqualError(t, err, "db fail")
	mockRepo.AssertExpectations(t)
}

func TestGetTaskByID_Success(t *testing.T) {
	mockRepo := new(MockTaskRepo)
	uc := NewTaskUseCase(mockRepo)

	id := primitive.NewObjectID()
	expected := &Domain.Task{Title: "A"}
	mockRepo.On("GetByID", id).Return(expected, nil)

	result, err := uc.GetTaskByID(id.Hex())
	assert.NoError(t, err)
	assert.Equal(t, expected, result)
	mockRepo.AssertExpectations(t)
}

func TestGetTaskByID_InvalidID(t *testing.T) {
	uc := NewTaskUseCase(new(MockTaskRepo))
	_, err := uc.GetTaskByID("badhex")
	assert.EqualError(t, err, "the provided hex string is not a valid ObjectID")
}

func TestCreateTask_Success(t *testing.T) {
	mockRepo := new(MockTaskRepo)
	uc := NewTaskUseCase(mockRepo)

	task := Domain.Task{Title: "T1"}
	created := &Domain.Task{Title: "T1"}
	mockRepo.On("Create", task).Return(created, nil)

	res, err := uc.CreateTask(task)
	assert.NoError(t, err)
	assert.Equal(t, created, res)
	mockRepo.AssertExpectations(t)
}

func TestCreateTask_Error(t *testing.T) {
	mockRepo := new(MockTaskRepo)
	uc := NewTaskUseCase(mockRepo)

	task := Domain.Task{Title: "Fail"}
	mockRepo.On("Create", task).Return((*Domain.Task)(nil), errors.New("insert fail"))

	res, err := uc.CreateTask(task)
	assert.Nil(t, res)
	assert.EqualError(t, err, "insert fail")
}

func TestUpdateTask_Success(t *testing.T) {
	mockRepo := new(MockTaskRepo)
	uc := NewTaskUseCase(mockRepo)

	id := primitive.NewObjectID()
	task := Domain.Task{Title: "Upd"}
	updated := &Domain.Task{Title: "Upd"}
	mockRepo.On("Update", id, task).Return(updated, nil)

	res, err := uc.UpdateTask(id.Hex(), task)
	assert.NoError(t, err)
	assert.Equal(t, updated, res)
	mockRepo.AssertExpectations(t)
}

func TestUpdateTask_DBError(t *testing.T) {
	mockRepo := new(MockTaskRepo)
	uc := NewTaskUseCase(mockRepo)

	id := primitive.NewObjectID()
	task := Domain.Task{Title: "Fail"}
	mockRepo.On("Update", id, task).Return((*Domain.Task)(nil), errors.New("db error"))

	res, err := uc.UpdateTask(id.Hex(), task)
	assert.Nil(t, res)
	assert.EqualError(t, err, "db error")
}

func TestUpdateTask_InvalidID(t *testing.T) {
	uc := NewTaskUseCase(new(MockTaskRepo))
	_, err := uc.UpdateTask("nope", Domain.Task{})
	assert.EqualError(t, err, "invalid ID format")
}

func TestDeleteTask_Success(t *testing.T) {
	mockRepo := new(MockTaskRepo)
	uc := NewTaskUseCase(mockRepo)

	id := primitive.NewObjectID()
	mockRepo.On("Delete", id).Return(nil)

	err := uc.DeleteTask(id.Hex())
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestDeleteTask_InvalidID(t *testing.T) {
	uc := NewTaskUseCase(new(MockTaskRepo))
	err := uc.DeleteTask("xxx")
	assert.EqualError(t, err, "invalid ID format")
}

func TestGetTaskByID_DBError(t *testing.T) {
	mockRepo := new(MockTaskRepo)
	uc := NewTaskUseCase(mockRepo)

	id := primitive.NewObjectID()
	mockRepo.On("GetByID", id).Return((*Domain.Task)(nil), errors.New("not found"))

	res, err := uc.GetTaskByID(id.Hex())
	assert.Nil(t, res)
	assert.EqualError(t, err, "not found")
}
