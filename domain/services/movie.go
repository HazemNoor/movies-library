package services

import (
	"errors"
	"github.com/HazemNoor/movies-library/domain/entities"
	"github.com/HazemNoor/movies-library/domain/repositories"
)

var ErrorMovieNotFound = errors.New("movie not found")

type MovieService struct {
	Repo repositories.MovieRepository
}

func (s *MovieService) GetAll() (movies *entities.Movies, err error) {
	movies, err = s.Repo.FindAll()
	return
}

func (s *MovieService) GetAllOrdered(orderBy string) (movies *entities.Movies, err error) {
	movies, err = s.Repo.FindAllOrdered(orderBy)
	return
}

func (s *MovieService) GetById(movieID uint) (*entities.Movie, error) {
	movie, err := s.Repo.FindByID(movieID)
	if err != nil {
		return nil, ErrorMovieNotFound
	}
	return movie, nil
}

func (s *MovieService) Create(movie *entities.Movie) error {
	return s.Repo.Create(movie)
}

func (s *MovieService) Update(movie *entities.Movie, data *entities.Movie) error {
	return s.Repo.Update(movie, data)
}

func (s *MovieService) Delete(movie *entities.Movie) error {
	return s.Repo.Delete(movie)
}
