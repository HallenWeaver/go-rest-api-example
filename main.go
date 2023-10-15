package main

import (
	"alexandre/gorest/app/routes"
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func main() {
	// Initializing MongoDB connection
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)

	mongoClient, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("MONGODB_URI")))

	defer func() {
		cancel()
		if err := mongoClient.Disconnect(ctx); err != nil {
			log.Fatalf("MongoDB disconnect error : %v", err)
		}
	}()

	if err != nil {
		log.Fatalf("Connection Error :%v", err)
		return
	}

	err = mongoClient.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatalf("Ping MongoDB Error :%v", err)
		return
	}
	fmt.Println("Ping Success")

	// Initializing Routes
	router := gin.Default()
	routes.InitializeRoutes(router, mongoClient)

	// Finishing Startup Configuration
	defaultPort := "8080"
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
