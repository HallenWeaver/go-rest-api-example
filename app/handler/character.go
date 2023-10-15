package handler

import (
	"alexandre/gorest/app/model"
	"alexandre/gorest/app/service"
	"fmt"

	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CharacterHandler struct {
	CharacterService service.CharacterService
}

func NewCharacterHandler(characterService service.CharacterService) *CharacterHandler {
	return &CharacterHandler{
		CharacterService: characterService,
	}
}

func (h *CharacterHandler) GetCharacters(c *gin.Context) {
	ownerID := c.Param("ownerId")
	characters, err := h.CharacterService.GetCharacters(c, ownerID, 10)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, characters)
}

func (h *CharacterHandler) GetCharacter(c *gin.Context) {
	ownerID := c.Param("ownerId")
	characterId, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	character, err := h.CharacterService.GetCharacter(c, ownerID, characterId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, character)
}

func (h *CharacterHandler) CreateCharacter(c *gin.Context) {
	var newCharacter model.Character

	if err := c.BindJSON(&newCharacter); err != nil {
		return
	}

	success, err := h.CharacterService.CreateCharacter(c, newCharacter)

	if success {
		c.IndentedJSON(http.StatusCreated, newCharacter)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	}
}

func (h *CharacterHandler) UpdateCharacter(c *gin.Context) {
	var editCharacter model.Character

	if err := c.BindJSON(&editCharacter); err != nil {
		return
	}
	editCharacter.OwnerId = c.Param("ownerId")

	characterId, err := primitive.ObjectIDFromHex(c.Param("id"))

	if err != nil {
		return
	}

	fmt.Printf("Edit Character ID: %+v\n", characterId)

	editCharacter.ID = characterId

	success, err := h.CharacterService.UpdateCharacter(c, editCharacter)
	if success {
		c.IndentedJSON(http.StatusCreated, editCharacter)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	}
}

func (h *CharacterHandler) DeleteCharacter(c *gin.Context) {
	ownerId := c.Param("ownerId")
	characterId, err := primitive.ObjectIDFromHex(c.Param("id"))

	if err != nil {
		return
	}

	success, err := h.CharacterService.DeleteCharacter(c, ownerId, characterId)
	if success {
		c.IndentedJSON(http.StatusOK, gin.H{})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	}
}
