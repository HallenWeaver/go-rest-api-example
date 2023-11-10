package routes

import (
	"alexandre/gorest/app/handler"
	"alexandre/gorest/app/middleware"
	"alexandre/gorest/app/repository"
	"alexandre/gorest/app/service"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func InitializeRoutes(router *gin.Engine, mongoClient *mongo.Client) {
	characterRepository := repository.NewCharacterRepository(mongoClient)
	characterService := service.NewCharacterService(characterRepository)
	characterHandler := handler.NewCharacterHandler(characterService)
	initializeCharacterRoutes(router, characterHandler)

	userRepository := repository.NewUserRepository(mongoClient)
	userService := service.NewUserService(userRepository)
	userHandler := handler.NewUserHandler(userService)
	initializeUserRoutes(router, userHandler)

	authenticationHandler := handler.NewAuthenticationHandler(userService)
	initializeAuthenticationRoutes(router, authenticationHandler)

	healthHandler := handler.NewHealthHandler()
	initializeHealthRoutes(router, healthHandler)
}

func initializeCharacterRoutes(router *gin.Engine, characterHandler *handler.CharacterHandler) {
	characterV1 := router.Group("/character").Use(middleware.Auth())
	characterV1.GET("/", characterHandler.GetCharacters)
	characterV1.GET("/:id", characterHandler.GetCharacter)
	characterV1.POST("", characterHandler.CreateCharacter)
	characterV1.PUT("/:id", characterHandler.UpdateCharacter)
	characterV1.DELETE("/:id", characterHandler.DeleteCharacter)
}

func initializeUserRoutes(router *gin.Engine, userHandler *handler.UserHandler) {
	userV1 := router.Group("/user")
	userV1.POST("", userHandler.CreateStandardUser)
}

func initializeAuthenticationRoutes(router *gin.Engine, authenticationHandler *handler.AuthenticationHandler) {
	userV1 := router.Group("/authenticate")
	userV1.POST("", authenticationHandler.LoginUser)
}

func initializeHealthRoutes(router *gin.Engine, healthHandler *handler.HealthHandler) {
	healthV1 := router.Group("/health")
	healthV1.GET("", healthHandler.GetAPIHealth)
}
