package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/surafelbkassa/go-task-manager/Delivery/controllers"
	"github.com/surafelbkassa/go-task-manager/Infrastructure"
)

// ← accept the interface, not the concrete struct
func SetupRouter(
	r *gin.Engine,
	jwtSvc Infrastructure.JWTServiceInterface,
	taskCtrl *controllers.TaskController,
	userCtrl *controllers.UserController,
) {
	auth := Infrastructure.AuthMiddleware

	r.GET("/tasks", auth(jwtSvc, ""), taskCtrl.GetTasks)
	r.GET("/tasks/:id", auth(jwtSvc, "user"), taskCtrl.GetTaskById)
	r.POST("/tasks", auth(jwtSvc, "user"), taskCtrl.CreateTask)
	r.PUT("/tasks/:id", auth(jwtSvc, "user"), taskCtrl.UpdatedTask)
	r.DELETE("/tasks/:id", auth(jwtSvc, "user"), taskCtrl.DeleteTask)

	r.POST("/register", userCtrl.RegisterUser)
	r.POST("/login", userCtrl.LoginUser)
	r.POST("/promote/:id", auth(jwtSvc, "admin"), userCtrl.PromoteUser)
}
