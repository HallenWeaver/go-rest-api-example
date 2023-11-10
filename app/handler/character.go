package handler

import (
	"alexandre/gorest/app/helper"
	"alexandre/gorest/app/model"
	"alexandre/gorest/app/service"

	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CharacterHandler struct {
	CharacterService service.ICharacterService
}

func NewCharacterHandler(characterService service.ICharacterService) *CharacterHandler {
	return &CharacterHandler{
		CharacterService: characterService,
	}
}

func (h *CharacterHandler) GetCharacters(c *gin.Context) {
	ownerID, err := helper.ParseUserDataFromToken(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	characters, err := h.CharacterService.GetCharacters(c, ownerID, 10)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, characters)
}

func (h *CharacterHandler) GetCharacter(c *gin.Context) {
	ownerID, err := helper.ParseUserDataFromToken(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

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
	ownerID, err := helper.ParseUserDataFromToken(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var newCharacter model.Character

	if err := c.BindJSON(&newCharacter); err != nil {
		return
	}

	newCharacter.OwnerId = ownerID

	savedCharacter, err := h.CharacterService.CreateCharacter(c, newCharacter)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.IndentedJSON(http.StatusCreated, savedCharacter)
	}
}

func (h *CharacterHandler) UpdateCharacter(c *gin.Context) {
	var editCharacter model.Character

	if err := c.BindJSON(&editCharacter); err != nil {
		return
	}

	ownerID, err := helper.ParseUserDataFromToken(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	editCharacter.OwnerId = ownerID

	characterId, err := primitive.ObjectIDFromHex(c.Param("id"))

	if err != nil {
		return
	}

	editCharacter.ID = characterId

	success, err := h.CharacterService.UpdateCharacter(c, editCharacter)
	if success {
		c.IndentedJSON(http.StatusCreated, editCharacter)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}

func (h *CharacterHandler) DeleteCharacter(c *gin.Context) {
	ownerID, err := helper.ParseUserDataFromToken(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	characterId, err := primitive.ObjectIDFromHex(c.Param("id"))

	if err != nil {
		return
	}

	success, err := h.CharacterService.DeleteCharacter(c, ownerID, characterId)
	if success {
		c.IndentedJSON(http.StatusOK, gin.H{})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}
