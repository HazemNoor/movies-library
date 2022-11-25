package repositories

import (
	"github.com/HazemNoor/movies-library/domain/entities"
)

type UserTokenRepository interface {
	Create(*entities.UserToken) error
	FindByToken(string) (*entities.UserToken, error)
}
