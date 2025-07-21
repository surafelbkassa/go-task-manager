package main

import (
	"fmt"
	"log"

	"github.com/surafelbkassa/go-task-manager/data"
	"github.com/surafelbkassa/go-task-manager/router"
)

func main() {
	data.InitMongo()
	data.InitMongoUser()
	router := router.SetupRouter()
	if err := router.Run("localhost:8080"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
	fmt.Println("Starting server on localhost:8080")
}
