package controllers

import (
	"github.com/HazemNoor/movies-library/app/forms"
	domainServices "github.com/HazemNoor/movies-library/domain/services"
	"github.com/HazemNoor/movies-library/infrastructure/repositories"
	"github.com/HazemNoor/movies-library/infrastructure/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AuthController struct {
	UserService domainServices.UserService
	AuthService domainServices.AuthService
}

func NewAuthController() *AuthController {
	return &AuthController{
		UserService: domainServices.UserService{
			Repo:      repositories.NewUserRepository(),
			Encryptor: services.NewEncryptor(),
		},
		AuthService: domainServices.AuthService{
			Repo: repositories.NewUserTokenRepository(),
		},
	}
}

func (uc *AuthController) Login(c *gin.Context) {
	form, err := forms.NewAuthLogin(c)
	if err != nil {
		validationError(c, err)
		return
	}

	user, err := uc.UserService.Login(form.Email, form.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	userToken, err := uc.AuthService.GenerateToken(user)
	if err != nil {
		serverError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"user_token": userToken})
}

func (uc *AuthController) Check(c *gin.Context) {
	user := getUserFromContext(c)
	c.JSON(http.StatusOK, gin.H{"user": user})
}
