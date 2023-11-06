package errorhandler

import (
	"github.com/gin-gonic/gin"
)

func HandleRequestError(c *gin.Context, statusCode int, err error) {
	c.JSON(statusCode, gin.H{"error": err.Error()})
}
