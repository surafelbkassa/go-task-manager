package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	Domain "github.com/surafelbkassa/go-task-manager/Domain"
	"github.com/surafelbkassa/go-task-manager/Infrastructure"
	"github.com/surafelbkassa/go-task-manager/Usecases"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// TaskController
type TaskController struct {
	uc Usecases.TaskUseCaseInterface
}

func NewTaskController(u Usecases.TaskUseCaseInterface) *TaskController {
	return &TaskController{uc: u}
}

func (tc *TaskController) GetTasks(c *gin.Context) {
	list, err := tc.uc.GetTasks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, list)
}

func (tc *TaskController) GetTaskById(c *gin.Context) {
	id := c.Param("id")
	task, err := tc.uc.GetTaskByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, task)
}

func (tc *TaskController) CreateTask(c *gin.Context) {
	var t Domain.Task
	if err := c.ShouldBindJSON(&t); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	created, err := tc.uc.CreateTask(t)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, created)
}

func (tc *TaskController) UpdatedTask(c *gin.Context) {
	id := c.Param("id")
	var t Domain.Task
	if err := c.ShouldBindJSON(&t); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	updated, err := tc.uc.UpdateTask(id, t)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, updated)
}

func (tc *TaskController) DeleteTask(c *gin.Context) {
	id := c.Param("id")
	if err := tc.uc.DeleteTask(id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "deleted"})
}

// UserController
type UserController struct {
	uc     Usecases.UserUseCaseInterface
	jwtSvc Infrastructure.JWTServiceInterface
}

func NewUserController(u Usecases.UserUseCaseInterface, j Infrastructure.JWTServiceInterface) *UserController {
	return &UserController{uc: u, jwtSvc: j}
}

func (uc *UserController) RegisterUser(c *gin.Context) {
	var body struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// pass three primitives, not a Domain.User
	if err := uc.uc.RegisterUser(body.Name, body.Email, body.Password); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "user registered"})
}

func (uc *UserController) LoginUser(c *gin.Context) {
	var body struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user, err := uc.uc.LoginUser(body.Email, body.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	token, err := uc.jwtSvc.GenerateToken(user.UserID, user.Role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": token})
}

func (uc *UserController) PromoteUser(c *gin.Context) {
	id := c.Param("id")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID format"})
		return
	}
	if _, err := uc.uc.PromoteUser(objID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "promoted"})
}
