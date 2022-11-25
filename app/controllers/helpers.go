package controllers

import (
	"github.com/HazemNoor/movies-library/domain/entities"
	"github.com/gin-gonic/gin"
	"net/http"
)

func validationError(c *gin.Context, err error) {
	c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
}

func serverError(c *gin.Context, err error) {
	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
}

func getUserFromContext(c *gin.Context) *entities.User {
	user, ok := c.Get("user")
	if !ok {
		return nil
	}

	userStruct, ok := user.(*entities.User)
	if !ok {
		return nil
	}
	return userStruct
}
