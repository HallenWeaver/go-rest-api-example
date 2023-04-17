package controller

import (
	character_service "alexandre/gorest/app/service"

	"net/http"

	"github.com/gin-gonic/gin"
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
	ownerID := c.Query("ownerId")
	characters, err := h.CharacterService.GetCharacters(ownerID, 10)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, characters)
}

// func GetCharacter(c *gin.Context) {
// 	characterId := c.Param("id")
// 	fmt.Printf("Getting character with id %s\n", characterId)
// 	connectionError := model.ConnectDatabase()
// 	if connectionError != nil {
// 		customError.SendResponse(c, customError.Response{Status: http.StatusInternalServerError, Error: []string{connectionError.Error()}})
// 		return
// 	}
// 	character, queryError := model.GetCharacterById(characterId)
// 	if queryError != nil {
// 		customError.SendResponse(c, customError.Response{Status: http.StatusNotFound, Error: []string{queryError.Error()}})
// 		return
// 	}
// 	if character.Id == characterId {
// 		c.IndentedJSON(http.StatusOK, character)
// 		return
// 	}
// 	customError.SendResponse(c, customError.Response{Status: http.StatusNotFound, Error: []string{"Character ID not found"}})
// }

// func PostCharacter(c *gin.Context) {
// 	var newCharacter model.Character

// 	if err := c.BindJSON(&newCharacter); err != nil {
// 		return
// 	}
// 	newCharacter.Id = uuid.NewString()

// 	connectionError := model.ConnectDatabase()
// 	if connectionError != nil {
// 		customError.SendResponse(c, customError.Response{Status: http.StatusInternalServerError, Error: []string{connectionError.Error()}})
// 		return
// 	}

// 	success, err := model.AddCharacter(newCharacter)

// 	if success {
// 		c.IndentedJSON(http.StatusCreated, newCharacter)
// 	} else {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err})
// 	}
// }

// func PutCharacter(c *gin.Context) {
// 	var editCharacter model.Character

// 	if err := c.BindJSON(&editCharacter); err != nil {
// 		return
// 	}

// 	editCharacter.Id = c.Param("id")

// 	success, err := model.UpdateCharacter(editCharacter)
// 	if success {
// 		c.IndentedJSON(http.StatusCreated, editCharacter)
// 	} else {
// 		customError.SendResponse(c, customError.Response{Status: http.StatusNotFound, Error: []string{err.Error()}})
// 	}
// }

// func DeleteCharacter(c *gin.Context) {
// 	characterId := c.Param("id")
// 	success, err := model.DeleteCharacter(characterId)
// 	if success {
// 		c.IndentedJSON(http.StatusOK, gin.H{})
// 	} else {
// 		customError.SendResponse(c, customError.Response{Status: http.StatusNotFound, Error: []string{err.Error()}})
// 	}
// }
