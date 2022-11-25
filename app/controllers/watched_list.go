package controllers

import (
	"github.com/HazemNoor/movies-library/app/forms"
	"github.com/HazemNoor/movies-library/domain/services"
	"github.com/HazemNoor/movies-library/infrastructure/repositories"
	"github.com/gin-gonic/gin"
	"net/http"
)

type WatchedListController struct {
	MovieService       services.MovieService
	WatchedListService services.WatchedListService
}

func NewWatchedListController() *WatchedListController {
	return &WatchedListController{
		MovieService: services.MovieService{
			Repo: repositories.NewMovieRepository(),
		},
		WatchedListService: services.WatchedListService{
			Repo: repositories.NewWatchedListRepository(),
		},
	}
}

func (l *WatchedListController) Add(c *gin.Context) {
	form, err := forms.NewWatchedListAdd(c)
	if err != nil {
		validationError(c, err)
		return
	}

	watchedListItem := form.ToDomain()

	_ = watchedListItem.SetUser(getUserFromContext(c))

	movie, err := l.MovieService.GetById(form.MovieId)
	if err != nil {
		validationError(c, err)
		return
	}

	if err = watchedListItem.SetMovie(movie); err != nil {
		validationError(c, err)
		return
	}

	err = l.WatchedListService.AddItem(watchedListItem)
	if err != nil {
		serverError(c, err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"watched_list_item": watchedListItem})
}

func (l *WatchedListController) Update(c *gin.Context) {
	form, err := forms.NewWatchedListUpdate(c)
	if err != nil {
		validationError(c, err)
		return
	}

	watchedListItem, err := l.WatchedListService.GetUserItemById(form.ID, getUserFromContext(c).ID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	err = l.WatchedListService.UpdateItem(watchedListItem, form.ToDomain())
	if err != nil {
		serverError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"watched_list_item": watchedListItem})
}

func (l *WatchedListController) Delete(c *gin.Context) {
	form, err := forms.NewWatchedListView(c)
	if err != nil {
		validationError(c, err)
		return
	}

	watchedListItem, err := l.WatchedListService.GetUserItemById(form.ID, getUserFromContext(c).ID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	if err = l.WatchedListService.DeleteItem(watchedListItem); err != nil {
		validationError(c, err)
		return
	}

	c.Status(http.StatusNoContent)
}

func (l *WatchedListController) List(c *gin.Context) {
	watchedList, err := l.WatchedListService.FindListForUser(getUserFromContext(c))
	if err != nil {
		serverError(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"watchedList": watchedList})
}
