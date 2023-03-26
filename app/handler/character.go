package handler

import (
	"alexandre/gorest/app/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

var baseCharacters = []model.Character{
	{ID: "0760b871-55d4-4316-a7b9-767581bc73bf", Name: "Hallen Weaver", Age: 19},
}

func GetCharacters(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, baseCharacters)
}

func PostCharacters(c *gin.Context) {
	var newCharacter model.Character

	if err := c.BindJSON(&newCharacter); err != nil {
		return
	}

	baseCharacters = append(baseCharacters, newCharacter)
	c.IndentedJSON(http.StatusCreated, newCharacter)
}
