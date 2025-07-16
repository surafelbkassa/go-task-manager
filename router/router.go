package router

import (
	"github.com/gin-gonic/gin"
	"github.com/surafelbkassa/go-task-manager/controllers"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/tasks", controllers.GetTask)
	router.GET("/tasks/:id", controllers.GetTaskById)
	router.PUT("/tasks/:id", controllers.UpdatedTask)
	router.POST("/tasks", controllers.CreateTask)
	router.DELETE("/tasks/:id", controllers.DeleteTask)
	return router
}
