package entities

import "time"

type UserToken struct {
	ID        uint      `json:"-"`
	UserId    uint      `json:"user_id"`
	User      *User     `json:"-"`
	Token     string    `json:"token"`
	CreatedAt time.Time `json:"created_at"`
	ExpiresAt time.Time `json:"expires_at"`
}
