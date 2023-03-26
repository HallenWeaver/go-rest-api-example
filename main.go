package main

import (
	"alexandre/gorest/app/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/characters", handler.GetCharacters)
	router.GET("/character/:id", handler.GetCharacter)
	router.POST("/character", handler.PostCharacter)
	router.PUT("/character", handler.PutCharacter)
	router.DELETE("/character/:id", handler.DeleteCharacter)

	router.Run("localhost:8080")
}
