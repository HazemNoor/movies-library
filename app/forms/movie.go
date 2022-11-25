package forms

import (
	"fmt"
	"github.com/HazemNoor/movies-library/domain/entities"
	"github.com/gin-gonic/gin"
	"time"
)

type MovieCreate struct {
	Name        string    `form:"name" binding:"required"`
	Description string    `form:"description" binding:"required"`
	Date        time.Time `form:"date" binding:"required" time_format:"2006-01-02"`
}

func NewMovieCreate(c *gin.Context) (f MovieCreate, err error) {
	err = c.ShouldBind(&f)
	return
}

func (f *MovieCreate) ToDomain() *entities.Movie {
	return &entities.Movie{
		Name:        f.Name,
		Description: f.Description,
		Date:        f.Date,
	}
}

type MovieUpdate struct {
	ID          uint      `uri:"id" binding:"required"`
	Name        string    `form:"name"`
	Description string    `form:"description"`
	Date        time.Time `form:"date" time_format:"2006-01-02"`
}

func NewMovieUpdate(c *gin.Context) (f MovieUpdate, err error) {
	if err = c.ShouldBindUri(&f); err == nil {
		if err = c.ShouldBind(&f); err == nil {
			return
		}
	}

	return
}

func (f *MovieUpdate) ToDomain() *entities.Movie {
	return &entities.Movie{
		Name:        f.Name,
		Description: f.Description,
		Date:        f.Date,
	}
}

type MovieView struct {
	ID uint `uri:"id" binding:"required"`
}

func NewMovieView(c *gin.Context) (f MovieView, err error) {
	err = c.ShouldBindUri(&f)
	return
}

type MovieFilters struct {
	Sort string `form:"sort" binding:"omitempty,oneof=date rate"`
	Dir  string `form:"dir" binding:"omitempty,oneof=asc desc"`
}

func NewMovieFilters(c *gin.Context) (f MovieFilters, err error) {
	err = c.ShouldBind(&f)
	if err == nil {
		if f.Sort == "" {
			f.Sort = "date"
		}

		if f.Dir == "" {
			f.Dir = "desc"
		}
	}
	return
}

func (f *MovieFilters) GetOrderBy() string {
	return fmt.Sprintf("%s %s", f.Sort, f.Dir)
}
