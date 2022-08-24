package main

import (
	"alexandre/gorest/app/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/characters", handler.GetCharacters)

	router.Run("localhost:8080")
}
