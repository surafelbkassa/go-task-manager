package main

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"github.com/surafelbkassa/go-task-manager/Delivery/router"
	"github.com/surafelbkassa/go-task-manager/Infrastructure"
	"github.com/surafelbkassa/go-task-manager/Repositories"
	"github.com/surafelbkassa/go-task-manager/Usecases"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file:", err)
	}
	Infrastructure.InitMongo()
	Infrastructure.InitMongoUser()

	taskRepo := Repositories.NewTaskRepository(Infrastructure.TaskCollection)
	userRepo := Repositories.NewUserRepository(Infrastructure.UserCollection)

	taskUC := Usecases.NewTaskUseCase(taskRepo)
	userUC := Usecases.NewUserUseCase(userRepo)

	router := router.SetupRouter(taskUC, userUC)
	if err := router.Run("localhost:8080"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
	fmt.Println("Starting server on localhost:8080")
}
