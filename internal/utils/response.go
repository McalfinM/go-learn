package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func JSON(c *gin.Context, status int, data interface{}) {
	c.JSON(status, data)
}

func Error(c *gin.Context, err error) {
	if apiErr, ok := err.(*ApiError); ok {
		c.JSON(apiErr.StatusCode, gin.H{
			"error": apiErr.Message,
		})
		return
	}

	c.JSON(http.StatusInternalServerError, gin.H{
		"error": "Internal Server Error",
	})
}
