package router

import (
	"github.com/gin-gonic/gin"
	"github.com/surafelbkassa/go-task-manager/controllers"
	"github.com/surafelbkassa/go-task-manager/middleware"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/tasks", middleware.AuthMiddleware(""), controllers.GetTasks)
	router.GET("/tasks/:id", middleware.AuthMiddleware("user"), controllers.GetTaskById)
	router.PUT("/tasks/:id", middleware.AuthMiddleware("user"), controllers.UpdatedTask)
	router.POST("/register", controllers.RegisterUser)
	router.POST("/login", controllers.LoginUser)
	router.POST("/tasks", middleware.AuthMiddleware("user"), controllers.CreateTask)
	router.DELETE("/tasks/:id", middleware.AuthMiddleware("user"), controllers.DeleteTask)
	router.POST("/promote/:id", middleware.AuthMiddleware("admin"), controllers.PromoteUser)
	return router
}
