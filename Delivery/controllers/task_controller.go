package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	domain "github.com/surafelbkassa/go-task-manager/Domain"
	"github.com/surafelbkassa/go-task-manager/Usecases"
)

type Controller struct {
	TaskUC Usecases.TaskUseCaseInterface
	UserUC Usecases.UserUseCaseInterface
}

func NewController(taskUC Usecases.TaskUseCaseInterface, userUC Usecases.UserUseCaseInterface) *Controller {
	return &Controller{
		TaskUC: taskUC,
		UserUC: userUC,
	}
}

func (c *Controller) RegisterUser(ctx *gin.Context) {
	var user domain.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	_, err := c.UserUC.RegisterUser(user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
}

func (c *Controller) LoginUser(ctx *gin.Context) {
	var loginReq struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := ctx.ShouldBindJSON(&loginReq); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	token, err := c.UserUC.LoginUser(loginReq.Email, loginReq.Password)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"token": token})
}

func (c *Controller) PromoteUser(ctx *gin.Context) {
	userID := ctx.Param("id") // assuming URL is /users/:id/promote

	if userID == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "User ID required"})
		return
	}

	err := c.UserUC.PromoteUser(userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "User promoted successfully"})
}
func (c *Controller) GetTasks(ctx *gin.Context) {
	tasks, err := c.TaskUC.GetTasks()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, tasks)
}

func (c *Controller) GetTaskByID(ctx *gin.Context) {
	id := ctx.Param("id")
	task, err := c.TaskUC.GetTaskByID(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, task)
}

func (c *Controller) CreateTask(ctx *gin.Context) {
	var task domain.Task
	if err := ctx.ShouldBindJSON(&task); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	created, err := c.TaskUC.CreateTask(task)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, created)
}

func (c *Controller) UpdateTask(ctx *gin.Context) {
	id := ctx.Param("id")
	var task domain.Task
	if err := ctx.ShouldBindJSON(&task); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	updated, err := c.TaskUC.UpdateTask(id, task)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, updated)
}

func (c *Controller) DeleteTask(ctx *gin.Context) {
	id := ctx.Param("id")
	err := c.TaskUC.DeleteTask(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	ctx.Status(http.StatusNoContent)
}
