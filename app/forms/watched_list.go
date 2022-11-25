package forms

import (
	"github.com/HazemNoor/movies-library/domain/entities"
	"github.com/gin-gonic/gin"
)

type WatchedListAdd struct {
	MovieId uint   `form:"movie_id" binding:"required"`
	Rate    int    `form:"rate" binding:"required,min=1,max=5"`
	Review  string `form:"review" binding:"omitempty"`
}

func NewWatchedListAdd(c *gin.Context) (f WatchedListAdd, err error) {
	err = c.ShouldBind(&f)
	return
}

func (f *WatchedListAdd) ToDomain() *entities.WatchedListItem {
	return &entities.WatchedListItem{
		Rate:   f.Rate,
		Review: f.Review,
	}
}

type WatchedListUpdate struct {
	ID     uint   `uri:"id" binding:"required"`
	Rate   int    `form:"rate" binding:"omitempty,min=1,max=5"`
	Review string `form:"review" binding:"omitempty"`
}

func NewWatchedListUpdate(c *gin.Context) (f WatchedListUpdate, err error) {
	if err = c.ShouldBindUri(&f); err == nil {
		if err = c.ShouldBind(&f); err == nil {
			return
		}
	}

	return
}

func (f *WatchedListUpdate) ToDomain() *entities.WatchedListItem {
	return &entities.WatchedListItem{
		Rate:   f.Rate,
		Review: f.Review,
	}
}

type WatchedListView struct {
	ID uint `uri:"id" binding:"required"`
}

func NewWatchedListView(c *gin.Context) (f WatchedListView, err error) {
	err = c.ShouldBindUri(&f)
	return
}
