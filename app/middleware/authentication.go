package middleware

import (
	"alexandre/gorest/app/util"

	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, err := util.ParseUserDataFromToken(c)
		if err != nil {
			c.JSON(401, gin.H{"error": err.Error()})
			c.Abort()
			return
		}
		c.Next()
	}
}
