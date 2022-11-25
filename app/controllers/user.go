package controllers

import (
	"github.com/HazemNoor/movies-library/app/forms"
	domainServices "github.com/HazemNoor/movies-library/domain/services"
	"github.com/HazemNoor/movies-library/infrastructure/repositories"
	"github.com/HazemNoor/movies-library/infrastructure/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserController struct {
	UserService domainServices.UserService
}

func NewUserController() *UserController {
	return &UserController{
		UserService: domainServices.UserService{
			Repo:      repositories.NewUserRepository(),
			Encryptor: services.NewEncryptor(),
		},
	}
}

func (uc *UserController) Register(c *gin.Context) {
	form, err := forms.NewUserCreate(c)
	if err != nil {
		validationError(c, err)
		return
	}

	user := form.ToDomain()

	err = uc.UserService.Register(user)
	if err != nil {
		// todo: set proper http status code according to error type
		validationError(c, err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"user": user})
}
