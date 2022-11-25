package repositories

import "github.com/HazemNoor/movies-library/domain/entities"

type WatchedListRepository interface {
	CreateItem(*entities.WatchedListItem) error
	UpdateItem(item *entities.WatchedListItem, data *entities.WatchedListItem) error
	DeleteItem(*entities.WatchedListItem) error
	FindUserItemByID(itemId uint, userId uint) (*entities.WatchedListItem, error)
	FindUserList(uint) (*entities.WatchedList, error)
}
