package repositories

import (
	"github.com/HazemNoor/movies-library/domain/entities"
)

type MovieRepository interface {
	FindAll() (*entities.Movies, error)
	FindAllOrdered(string) (*entities.Movies, error)
	Create(*entities.Movie) error
	FindByID(uint) (*entities.Movie, error)
	Update(movie *entities.Movie, data *entities.Movie) error
	Delete(*entities.Movie) error
}
