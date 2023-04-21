package main

import (
	"fmt"
	"log"
	"os"

	"alexandre/gorest/app/handler"
	character_repository "alexandre/gorest/app/repository"
	character_service "alexandre/gorest/app/service"

	routing "alexandre/gorest/app/router"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

const defaultPort = "8080"

func main() {
	// Set up router and initialize routes
	router := gin.Default()
	characterRepository, _ := character_repository.NewCharacterRepository()
	characterService := character_service.NewCharacterService(*characterRepository)
	characterHandler := handler.NewCharacterHandler(*characterService)
	routing.InitializeRoutes(router, characterHandler)

	// Get port from environment variable or use default
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	// Start server
	log.Printf("Starting server on port %s\n", port)
	if err := router.Run(fmt.Sprintf(":%s", port)); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
