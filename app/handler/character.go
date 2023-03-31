package handler

import (
	"alexandre/gorest/app/error"
	"alexandre/gorest/app/model"

	"net/http"

	"github.com/gin-gonic/gin"
)

var baseCharacters = []model.Character{
	{Id: "0760b871-55d4-4316-a7b9-767581bc73bf", OwnerId: "7f485de9-7af0-4c93-8ba9-d6562984ad80", Name: "Hallen Weaver", Age: 19},
}

func GetCharacters(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, baseCharacters)
}

func GetCharacter(c *gin.Context) {
	characterId := c.Param("id")
	for _, character := range baseCharacters {
		if character.Id == characterId {
			c.IndentedJSON(http.StatusOK, character)
			return
		}
	}
	error.SendResponse(c, error.Response{Status: http.StatusNotFound, Error: []string{"Character ID not found"}})
}

func PostCharacter(c *gin.Context) {
	var newCharacter model.Character

	if err := c.BindJSON(&newCharacter); err != nil {
		return
	}

	baseCharacters = append(baseCharacters, newCharacter)
	c.IndentedJSON(http.StatusCreated, newCharacter)
}

func PutCharacter(c *gin.Context) {
	var editCharacter model.Character

	if err := c.BindJSON(&editCharacter); err != nil {
		return
	}

	for index, character := range baseCharacters {
		if character.Id == editCharacter.Id {
			baseCharacters[index] = editCharacter
			c.IndentedJSON(http.StatusCreated, editCharacter)
			return
		}
	}
	error.SendResponse(c, error.Response{Status: http.StatusNotFound, Error: []string{"Character ID not found"}})
}

func DeleteCharacter(c *gin.Context) {
	characterId := c.Param("id")
	for index, character := range baseCharacters {
		if character.Id == characterId {
			baseCharacters = append(baseCharacters[:index], baseCharacters[index+1:]...)
			return
		}
	}
	error.SendResponse(c, error.Response{Status: http.StatusNotFound, Error: []string{"Character ID not found"}})
}
