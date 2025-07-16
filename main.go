package main

import (
	"fmt"

	"github.com/surafelbkassa/go-task-manager/router"
)

func main() {
	router := router.SetupRouter()
	fmt.Println("Starting server on localhost:8080")
	router.Run("localhost:8080")
}
