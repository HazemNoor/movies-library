package forms

import (
	"github.com/HazemNoor/movies-library/domain/entities"
	"github.com/gin-gonic/gin"
)

type UserCreate struct {
	Name     string `form:"name" binding:"required"`
	Age      uint   `form:"age" binding:"required"`
	Email    string `form:"email" binding:"required,email"`
	Password string `form:"password" binding:"required"`
}

func NewUserCreate(c *gin.Context) (f UserCreate, err error) {
	err = c.ShouldBind(&f)
	return
}

func (f *UserCreate) ToDomain() *entities.User {
	return &entities.User{
		Name:     f.Name,
		Age:      f.Age,
		Email:    f.Email,
		Password: f.Password,
	}
}
