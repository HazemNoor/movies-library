package services

import (
	"errors"
	"github.com/HazemNoor/movies-library/domain/entities"
	"github.com/HazemNoor/movies-library/domain/repositories"
)

var ErrorExistsBefore = errors.New("email exists before")
var ErrorInvalidLogin = errors.New("invalid login details")

type UserService struct {
	Repo      repositories.UserRepository
	Encryptor Encryptor
}

func (s *UserService) Register(user *entities.User) error {
	_, err := s.Repo.FindByEmail(user.Email)
	if err == nil {
		return ErrorExistsBefore
	}

	user.Password, err = s.Encryptor.Encrypt(user.Password)
	if err != nil {
		return err
	}

	return s.Repo.Create(user)
}

func (s *UserService) GetAll() (users *entities.Users, err error) {
	users, err = s.Repo.FindAll()

	return
}

func (s *UserService) GetById(userID uint) (*entities.User, error) {
	return s.Repo.FindByID(userID)
}

func (s *UserService) GetByEmail(email string) (*entities.User, error) {
	return s.Repo.FindByEmail(email)
}

func (s *UserService) Login(email string, password string) (*entities.User, error) {
	user, err := s.GetByEmail(email)
	if err != nil {
		return nil, ErrorInvalidLogin
	}

	err = s.Encryptor.Check(user.Password, password)
	if err != nil {
		return nil, ErrorInvalidLogin
	}

	return user, nil
}
