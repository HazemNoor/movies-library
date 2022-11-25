package repositories

import (
	"github.com/HazemNoor/movies-library/domain/entities"
)

type UserRepository interface {
	FindAll() (*entities.Users, error)
	FindByID(uint) (*entities.User, error)
	FindByEmail(string) (*entities.User, error)
	Create(*entities.User) error
}
