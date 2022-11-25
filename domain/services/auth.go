package services

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"github.com/HazemNoor/movies-library/domain/entities"
	"github.com/HazemNoor/movies-library/domain/repositories"
	"time"
)

type AuthService struct {
	Repo repositories.UserTokenRepository
}

var ErrorTokenNotFound = errors.New("token not found")

func (s *AuthService) GenerateToken(user *entities.User) (userToken entities.UserToken, err error) {
	randomToken, err := s.generateRandomToken()
	if err != nil {
		return userToken, err
	}

	userToken = entities.UserToken{
		UserId:    user.ID,
		Token:     randomToken,
		ExpiresAt: time.Now().Add(time.Hour * 24),
	}

	err = s.Repo.Create(&userToken)

	return userToken, err
}

func (s *AuthService) generateRandomToken() (string, error) {
	randomToken := make([]byte, 32)
	_, err := rand.Read(randomToken)
	if err != nil {
		return "", err
	}

	return base64.URLEncoding.EncodeToString(randomToken), nil
}

func (s *AuthService) ValidateToken(token string) (*entities.UserToken, error) {
	userToken, err := s.Repo.FindByToken(token)
	if err != nil {
		return nil, ErrorTokenNotFound
	}

	return userToken, nil
}
