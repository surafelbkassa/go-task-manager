package main

import (
	"fmt"
	"net/http"

	"time"

	"github.com/gin-gonic/gin"
)

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

func main() {
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Pong"})
	})
	router.GET("/tasks", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"tasks": tasks})
	})
	router.GET("/tasks/:id", func(ctx *gin.Context) {
		id := ctx.Param("id")
		for _, task := range tasks {
			if task.ID == id {
				ctx.JSON(http.StatusAccepted, task)
				return
			}
		}
		ctx.JSON(http.StatusNotFound, gin.H{"message": "Task not found"})
	})
	router.PUT("/tasks/:id", func(ctx *gin.Context) {
		var updatedTask Task
		if err := ctx.ShouldBindJSON(&updatedTask); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
		id := ctx.Param("id")
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
				ctx.JSON(http.StatusAccepted, gin.H{"message": "Task updated successfully", "task": tasks[i]})
				return
			}
		}
		ctx.JSON(http.StatusNotFound, gin.H{"message": "Task not found"})
	})
	router.POST("/tasks", func(ctx *gin.Context) {
		var newTask Task
		if err := ctx.ShouldBindJSON(&newTask); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		tasks = append(tasks, newTask)
		ctx.JSON(http.StatusAccepted, gin.H{"message": "Task created"})
	})
	router.Run("localhost:8080")
	fmt.Println("Task manager")
}
