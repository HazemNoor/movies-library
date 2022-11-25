package services

import (
	"errors"
	"github.com/HazemNoor/movies-library/domain/entities"
	"github.com/HazemNoor/movies-library/domain/repositories"
)

var ErrorItemNotFound = errors.New("item not found")

type WatchedListService struct {
	Repo repositories.WatchedListRepository
}

func (s *WatchedListService) FindListForUser(user *entities.User) (list *entities.WatchedList, err error) {
	list, err = s.Repo.FindUserList(user.ID)
	return
}

func (s *WatchedListService) GetUserItemById(itemId uint, userId uint) (*entities.WatchedListItem, error) {
	movie, err := s.Repo.FindUserItemByID(itemId, userId)
	if err != nil {
		return nil, ErrorItemNotFound
	}
	return movie, nil
}

func (s *WatchedListService) AddItem(item *entities.WatchedListItem) error {
	return s.Repo.CreateItem(item)
}

func (s *WatchedListService) UpdateItem(item *entities.WatchedListItem, data *entities.WatchedListItem) error {
	return s.Repo.UpdateItem(item, data)
}

func (s *WatchedListService) DeleteItem(item *entities.WatchedListItem) error {
	return s.Repo.DeleteItem(item)
}
