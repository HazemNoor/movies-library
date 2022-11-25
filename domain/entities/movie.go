package entities

import (
	"errors"
	"time"
)

var ErrorMovieBelongsToAnotherUser = errors.New("this movie belongs to another user, you can't edit or delete")

type Movies []Movie

type Movie struct {
	ID          uint      `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Date        time.Time `json:"date"`
	Cover       string    `json:"cover"`
	Rate        float32   `json:"rate"`
	UserId      uint      `json:"user_id"`
	User        *User     `json:"-"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (m *Movie) SetUser(user *User) error {
	if m.User != nil && m.UserId != user.ID {
		return ErrorMovieBelongsToAnotherUser
	}
	m.UserId = user.ID
	m.User = user
	return nil
}
