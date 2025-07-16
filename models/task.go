package models

import "time"

type Task struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	DueDate     time.Time `json:"due_date"`
	Status      string    `json:"status"`
}

var tasks = []Task{
	{ID: "1", Title: "Task 1", Description: "Description for Task 1", DueDate: time.Now().Add(24 * time.Hour), Status: "Pending"},
	{ID: "2", Title: "Task 2", Description: "Description for Task 2", DueDate: time.Now().Add(48 * time.Hour), Status: "In Progress"},
	{ID: "3", Title: "Task 3", Description: "Description for Task 3", DueDate: time.Now().Add(72 * time.Hour), Status: "Completed"},
}
