package data

import (
	"errors"
	"time"

	"github.com/surafelbkassa/go-task-manager/models"
)

var tasks = []models.Task{
	{ID: "1", Title: "Task 1", Description: "Description for Task 1", DueDate: time.Now().Add(24 * time.Hour), Status: "Pending"},
	{ID: "2", Title: "Task 2", Description: "Description for Task 2", DueDate: time.Now().Add(48 * time.Hour), Status: "In Progress"},
	{ID: "3", Title: "Task 3", Description: "Description for Task 3", DueDate: time.Now().Add(72 * time.Hour), Status: "Completed"},
}

func GetTask() []models.Task {
	return tasks
}

func GetTaskById(id string) (*models.Task, error) {
	for _, task := range tasks {
		if task.ID == id {
			return &task, nil
		}
	}
	return nil, errors.New("task not found")
}
func UpdatedTask(id string, updatedTask models.Task) (*models.Task, error) {

	for i, task := range tasks {
		if id == task.ID {
			if updatedTask.Title != "" {
				tasks[i].Title = updatedTask.Title
			}
			if updatedTask.Description != "" {
				tasks[i].Description = updatedTask.Description
			}
			if !updatedTask.DueDate.IsZero() {
				tasks[i].DueDate = updatedTask.DueDate
			}
			if updatedTask.Status != "" {
				tasks[i].Status = updatedTask.Status
			}
			return &tasks[i], nil
		}
	}
	return nil, errors.New("task not found")

}
func CreateTask(newTask models.Task) {
	tasks = append(tasks, newTask)
}
func DeleteTask(id string) error {
	for i, task := range tasks {
		if task.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			return nil
		}
	}
	return errors.New("task not found")
}
