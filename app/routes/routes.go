package routes

import (
	"alexandre/gorest/app/handler"
	character_repository "alexandre/gorest/app/repository"
	character_service "alexandre/gorest/app/service"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func InitializeRoutes(router *gin.Engine, mongoClient *mongo.Client) {
	characterRepository, _ := character_repository.NewCharacterRepository(mongoClient)
	characterService := character_service.NewCharacterService(*characterRepository)
	characterHandler := handler.NewCharacterHandler(*characterService)
	initializeCharacterRoutes(router, characterHandler)

	healthHandler := handler.NewHealthHandler()
	initializeHealthRoutes(router, healthHandler)
}

func initializeCharacterRoutes(router *gin.Engine, characterHandler *handler.CharacterHandler) {
	characterV1 := router.Group("/character")
	characterV1.GET("/:ownerId", characterHandler.GetCharacters)
	characterV1.GET("/:ownerId/:id", characterHandler.GetCharacter)
	characterV1.POST("", characterHandler.CreateCharacter)
	characterV1.PUT("/:ownerId/:id", characterHandler.UpdateCharacter)
	characterV1.DELETE("/:ownerId/:id", characterHandler.DeleteCharacter)
}

func initializeHealthRoutes(router *gin.Engine, healthHandler *handler.HealthHandler) {
	healthV1 := router.Group("/health")
	healthV1.GET("", healthHandler.GetAPIHealth)
}
