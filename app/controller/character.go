package controller

import (
	"alexandre/gorest/app/error"
	"alexandre/gorest/app/model"
	"encoding/json"
	"fmt"

	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var baseCharacters = []model.Character{}

func GetCharacters(c *gin.Context) {
	fmt.Println("Getting characters...")
	connectionError := model.ConnectDatabase()
	if connectionError != nil {
		error.SendResponse(c, error.Response{Status: http.StatusInternalServerError, Error: []string{connectionError.Error()}})
		return
	}
	results, queryError := model.GetCharacters(10)
	if queryError != nil {
		error.SendResponse(c, error.Response{Status: http.StatusNotFound, Error: []string{queryError.Error()}})
		return
	}

	for _, result := range results {
		res, _ := json.Marshal(result)
		fmt.Println(string(res))
	}

	c.IndentedJSON(http.StatusOK, results)
}

func GetCharacter(c *gin.Context) {
	characterId := c.Param("id")
	fmt.Printf("Getting character with id %s\n", characterId)
	connectionError := model.ConnectDatabase()
	if connectionError != nil {
		error.SendResponse(c, error.Response{Status: http.StatusInternalServerError, Error: []string{connectionError.Error()}})
		return
	}
	character, queryError := model.GetCharacterById(characterId)
	if queryError != nil {
		error.SendResponse(c, error.Response{Status: http.StatusNotFound, Error: []string{queryError.Error()}})
		return
	}
	if character.Id == characterId {
		c.IndentedJSON(http.StatusOK, character)
		return
	}
	error.SendResponse(c, error.Response{Status: http.StatusNotFound, Error: []string{"Character ID not found"}})
}

func PostCharacter(c *gin.Context) {
	var newCharacter model.Character

	if err := c.BindJSON(&newCharacter); err != nil {
		return
	}
	newCharacter.Id = uuid.NewString()

	connectionError := model.ConnectDatabase()
	if connectionError != nil {
		error.SendResponse(c, error.Response{Status: http.StatusInternalServerError, Error: []string{connectionError.Error()}})
		return
	}

	success, err := model.AddCharacter(newCharacter)

	if success {
		c.IndentedJSON(http.StatusCreated, newCharacter)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	}
}

func PutCharacter(c *gin.Context) {
	var editCharacter model.Character

	if err := c.BindJSON(&editCharacter); err != nil {
		return
	}

	editCharacter.Id = c.Param("id")

	success, err := model.UpdateCharacter(editCharacter)
	if success {
		c.IndentedJSON(http.StatusCreated, editCharacter)
	} else {
		error.SendResponse(c, error.Response{Status: http.StatusNotFound, Error: []string{err.Error()}})
	}
}

func DeleteCharacter(c *gin.Context) {
	characterId := c.Param("id")
	success, err := model.DeleteCharacter(characterId)
	if success {
		c.IndentedJSON(http.StatusOK, gin.H{})
	} else {
		error.SendResponse(c, error.Response{Status: http.StatusNotFound, Error: []string{err.Error()}})
	}
}
