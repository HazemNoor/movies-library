package forms

import (
	"github.com/gin-gonic/gin"
)

type AuthLogin struct {
	Email    string `form:"email" binding:"required,email"`
	Password string `form:"password" binding:"required"`
}

func NewAuthLogin(c *gin.Context) (f AuthLogin, err error) {
	err = c.ShouldBind(&f)
	return
}
