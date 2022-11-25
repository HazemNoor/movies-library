package controllers

import (
	"github.com/HazemNoor/movies-library/app/forms"
	"github.com/HazemNoor/movies-library/domain/services"
	"github.com/HazemNoor/movies-library/infrastructure/repositories"
	"github.com/gin-gonic/gin"
	"net/http"
)

type MovieController struct {
	MovieService services.MovieService
}

func NewMovieController() *MovieController {
	return &MovieController{
		MovieService: services.MovieService{
			Repo: repositories.NewMovieRepository(),
		},
	}
}

func (mc *MovieController) Create(c *gin.Context) {
	form, err := forms.NewMovieCreate(c)
	if err != nil {
		validationError(c, err)
		return
	}

	movie := form.ToDomain()

	_ = movie.SetUser(getUserFromContext(c))

	err = mc.MovieService.Create(movie)
	if err != nil {
		serverError(c, err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"movie": movie})
}

func (mc *MovieController) Update(c *gin.Context) {
	form, err := forms.NewMovieUpdate(c)
	if err != nil {
		validationError(c, err)
		return
	}

	movie, err := mc.MovieService.GetById(form.ID)
	if err != nil {
		validationError(c, err)
		return
	}

	err = movie.SetUser(getUserFromContext(c))
	if err != nil {
		validationError(c, err)
		return
	}

	err = mc.MovieService.Update(movie, form.ToDomain())
	if err != nil {
		serverError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"movie": movie})
}

func (mc *MovieController) Show(c *gin.Context) {
	form, err := forms.NewMovieView(c)
	if err != nil {
		validationError(c, err)
		return
	}

	movie, err := mc.MovieService.GetById(form.ID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"movie": movie})
}

func (mc *MovieController) Delete(c *gin.Context) {
	form, err := forms.NewMovieView(c)
	if err != nil {
		validationError(c, err)
		return
	}

	movie, err := mc.MovieService.GetById(form.ID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	if err = movie.SetUser(getUserFromContext(c)); err != nil {
		validationError(c, err)
		return
	}

	if err = mc.MovieService.Delete(movie); err != nil {
		validationError(c, err)
		return
	}

	c.Status(http.StatusNoContent)
}

func (mc *MovieController) List(c *gin.Context) {
	form, err := forms.NewMovieFilters(c)
	if err != nil {
		validationError(c, err)
		return
	}

	movies, err := mc.MovieService.GetAllOrdered(form.GetOrderBy())
	if err != nil {
		serverError(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"movies": movies})
}
