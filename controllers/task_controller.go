package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/surafelbkassa/go-task-manager/data"
	"github.com/surafelbkassa/go-task-manager/models"
)

func GetTask(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"tasks": data.GetTask()})
}
func GetTaskById(ctx *gin.Context) {
	id := ctx.Param("id")
	task, err := data.GetTaskById(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, task)
}
func CreateTask(c *gin.Context) {
	var newTask models.Task
	if err := c.ShouldBindJSON(&newTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	data.CreateTask(newTask)
	c.JSON(http.StatusCreated, gin.H{"message": "Task created successfully", "task": newTask})
}
func UpdatedTask(c *gin.Context) {
	var updatedTask models.Task
	if err := c.ShouldBindJSON(&updatedTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	id := c.Param("id")
	task, err := data.UpdatedTask(id, updatedTask)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Task updated successfully", "task": task})
}
func DeleteTask(x *gin.Context) {
	id := x.Param("id")
	err := data.DeleteTask(id)
	if err != nil {
		x.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}
	x.JSON(http.StatusOK, gin.H{"message": "Task deleted successfully"})
}
