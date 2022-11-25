package middleware

import (
	domainServices "github.com/HazemNoor/movies-library/domain/services"
	"github.com/HazemNoor/movies-library/infrastructure/repositories"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

type AuthMiddleware struct {
	AuthService domainServices.AuthService
}

func NewAuthMiddleware() *AuthMiddleware {
	return &AuthMiddleware{
		AuthService: domainServices.AuthService{
			Repo: repositories.NewUserTokenRepository(),
		},
	}
}

func (m *AuthMiddleware) Handle(c *gin.Context) {
	tokenString := m.getTokenFromRequest(c)

	userToken, err := m.AuthService.ValidateToken(tokenString)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication failed"})
		c.Abort()
		return
	}

	c.Set("user", userToken.User)
	c.Next()
}

func (m *AuthMiddleware) getTokenFromRequest(c *gin.Context) string {
	bearerToken := c.Request.Header.Get("Authorization")
	splitToken := strings.Split(bearerToken, " ")
	if len(splitToken) == 2 {
		return splitToken[1]
	}
	return ""
}
