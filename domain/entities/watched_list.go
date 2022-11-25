package entities

import (
	"errors"
	"time"
)

var ErrorItemBelongsToAnotherUser = errors.New("this item belongs to another user, you can't edit or delete")

type WatchedList []WatchedListItem

type WatchedListItem struct {
	ID        uint      `json:"id"`
	UserId    uint      `json:"-"`
	User      *User     `json:"-"`
	MovieId   uint      `json:"movie_id"`
	Movie     *Movie    `json:"-"`
	Rate      int       `json:"rate"`
	Review    string    `json:"review"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (m *WatchedListItem) SetUser(user *User) error {
	if m.User != nil && m.UserId != user.ID {
		return ErrorItemBelongsToAnotherUser
	}
	m.UserId = user.ID
	m.User = user
	return nil
}

func (m *WatchedListItem) SetMovie(movie *Movie) error {
	m.MovieId = movie.ID
	m.Movie = movie
	return nil
}
