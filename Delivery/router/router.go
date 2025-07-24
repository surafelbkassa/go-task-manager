package router

import (
	"github.com/gin-gonic/gin"
	"github.com/surafelbkassa/go-task-manager/Delivery/controllers"
	"github.com/surafelbkassa/go-task-manager/Infrastructure"
	usecases "github.com/surafelbkassa/go-task-manager/Usecases"
)

func SetupRouter(taskUC usecases.TaskUseCaseInterface, userUC usecases.UserUseCaseInterface) *gin.Engine {
	r := gin.Default()
	ctr := controllers.NewController(taskUC, userUC)

	// Task routes
	tasks := r.Group("/tasks")
	tasks.Use(Infrastructure.AuthMiddleware("user"))
	tasks.GET("", ctr.GetTasks)
	tasks.POST("", ctr.CreateTask)
	tasks.GET(":id", ctr.GetTaskByID)
	tasks.PUT(":id", ctr.UpdateTask)
	tasks.DELETE(":id", ctr.DeleteTask)

	// Auth routes
	r.POST("/register", ctr.RegisterUser)
	r.POST("/login", ctr.LoginUser)
	r.POST("/promote/:id", Infrastructure.AuthMiddleware("admin"), ctr.PromoteUser)

	return r
}
