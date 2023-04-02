package main

import (
	"alexandre/gorest/app/controller"

	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine) {
	initializeCharacterRoutes(router)
}

func initializeCharacterRoutes(router *gin.Engine) {
	characterV1 := router.Group("/character")
	characterV1.GET("", controller.GetCharacters)
	characterV1.GET("/:id", controller.GetCharacter)
	characterV1.POST("", controller.PostCharacter)
	characterV1.PUT("/:id", controller.PutCharacter)
	characterV1.DELETE("/:id", controller.DeleteCharacter)
}
