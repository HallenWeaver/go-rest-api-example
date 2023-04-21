package handler

import (
	"alexandre/gorest/app/model"
	character_service "alexandre/gorest/app/service"

	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type CharacterHandler struct {
	CharacterService character_service.CharacterService
}

func NewCharacterHandler(characterService character_service.CharacterService) *CharacterHandler {
	return &CharacterHandler{
		CharacterService: characterService,
	}
}

func (h *CharacterHandler) GetCharacters(c *gin.Context) {
	ownerID := c.Param("ownerId")
	characters, err := h.CharacterService.GetCharacters(ownerID, 10)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, characters)
}

func (h *CharacterHandler) GetCharacter(c *gin.Context) {
	ownerID := c.Param("ownerId")
	characterId := c.Param("id")
	character, err := h.CharacterService.GetCharacter(ownerID, characterId)
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
	newCharacter.Id = uuid.NewString()

	success, err := h.CharacterService.CreateCharacter(newCharacter)

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
	editCharacter.Id = c.Param("id")

	success, err := h.CharacterService.UpdateCharacter(editCharacter)
	if success {
		c.IndentedJSON(http.StatusCreated, editCharacter)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	}
}

func (h *CharacterHandler) DeleteCharacter(c *gin.Context) {
	ownerId := c.Param("ownerId")
	characterId := c.Param("id")
	success, err := h.CharacterService.DeleteCharacter(ownerId, characterId)
	if success {
		c.IndentedJSON(http.StatusOK, gin.H{})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	}
}
