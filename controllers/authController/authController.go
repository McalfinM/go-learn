package authcontroller

import (
	"learn/api/domains/account"
	"net/http"

	"github.com/gin-gonic/gin"
)

type authcontroller struct {
	service account.Service
}

func NewAuthController(service account.Service) *authcontroller {
	return &authcontroller{service}
}

func (handler *authcontroller) FindAll(c *gin.Context) {
	user, err := handler.service.FindAll()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": user,
	})
}

func (handler *authcontroller) Register(c *gin.Context) {
	var input account.RegisterInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "validated!"})
}

func Logout(c *gin.Context) {

}
