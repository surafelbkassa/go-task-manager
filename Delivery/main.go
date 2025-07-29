package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/surafelbkassa/go-task-manager/Delivery/controllers"
	routers "github.com/surafelbkassa/go-task-manager/Delivery/router"
	"github.com/surafelbkassa/go-task-manager/Infrastructure"
	"github.com/surafelbkassa/go-task-manager/Repositories"
	"github.com/surafelbkassa/go-task-manager/Usecases"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	r := gin.Default()
	ctx := context.Background()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
	if err := client.Ping(ctx, nil); err != nil {
		log.Fatal(err)
	}

	// returns the interface type
	jwtSvc := Infrastructure.NewJWTService("secret-key", 24*time.Hour)
	hasher := Infrastructure.NewPasswordService()

	// repositories
	taskRepo := Repositories.NewTaskRepository(client.Database("task_manager").Collection("tasks"))
	userRepo := Repositories.NewUserRepository(client.Database("task_manager").Collection("users"), ctx)

	// use‐cases
	taskUC := Usecases.NewTaskUseCase(taskRepo)
	userUC := Usecases.NewUserUseCase(userRepo, hasher)

	// controllers
	taskCtrl := controllers.NewTaskController(taskUC)
	userCtrl := controllers.NewUserController(userUC, jwtSvc)

	// routes
	routers.SetupRouter(r, jwtSvc, taskCtrl, userCtrl)

	fmt.Println("Starting server on :8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatal(fmt.Sprintf("Failed to start server: %v", err))
	}
}
