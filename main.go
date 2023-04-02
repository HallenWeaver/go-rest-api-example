package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

var PORT = "8080"

func main() {
	// Map routes
	router := gin.Default()
	initializeRoutes(router)

	fmt.Printf("Starting server on port %s\n", PORT)

	router.Run(fmt.Sprintf("localhost:%s", PORT))
}
