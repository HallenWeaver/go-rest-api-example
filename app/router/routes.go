package routing

import (
	"alexandre/gorest/app/handler"

	"github.com/gin-gonic/gin"
)

func InitializeRoutes(router *gin.Engine, characterHandler *handler.CharacterHandler) {
	initializeCharacterRoutes(router, characterHandler)
}

func initializeCharacterRoutes(router *gin.Engine, characterHandler *handler.CharacterHandler) {
	characterV1 := router.Group("/character")
	characterV1.GET("/:ownerId", characterHandler.GetCharacters)
	characterV1.GET("/:ownerId/:id", characterHandler.GetCharacter)
	characterV1.POST("", characterHandler.CreateCharacter)
	characterV1.PUT("/:ownerId/:id", characterHandler.UpdateCharacter)
	characterV1.DELETE(":ownerId/:id", characterHandler.DeleteCharacter)
}
