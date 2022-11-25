package repositories

import (
	"github.com/HazemNoor/movies-library/domain/entities"
	"github.com/HazemNoor/movies-library/infrastructure"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type WatchedListRepository struct {
	db *gorm.DB
}

func NewWatchedListRepository() WatchedListRepository {
	db, _ := infrastructure.DbConnection()
	return WatchedListRepository{db: db}
}

func (w WatchedListRepository) CreateItem(item *entities.WatchedListItem) error {
	tx := w.db.Begin()
	result := w.db.Omit(clause.Associations).Create(item)
	if result.Error != nil {
		tx.Rollback()
		return result.Error
	}
	if err := w.updateMovieRate(item, tx); err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func (w WatchedListRepository) UpdateItem(item *entities.WatchedListItem, data *entities.WatchedListItem) error {
	tx := w.db.Begin()
	result := tx.Omit(clause.Associations).Model(&item).Updates(data)
	if result.Error != nil {
		tx.Rollback()
		return result.Error
	}
	if err := w.updateMovieRate(item, tx); err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func (w WatchedListRepository) DeleteItem(item *entities.WatchedListItem) error {
	tx := w.db.Begin()
	result := w.db.Delete(&item)
	if result.Error != nil {
		tx.Rollback()
		return result.Error
	}
	if err := w.updateMovieRate(item, tx); err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func (w WatchedListRepository) updateMovieRate(item *entities.WatchedListItem, tx *gorm.DB) error {
	itemsList := new(entities.WatchedList)
	w.db.
		Where("`movie_id` = ?", item.MovieId).
		Where("`id` != ?", item.ID).
		Find(&itemsList)

	count := len(*itemsList) + 1
	sum := item.Rate
	for _, itemList := range *itemsList {
		sum += itemList.Rate
	}

	rate := float32(sum) / float32(count)
	result := tx.Omit(clause.Associations).Model(&item.Movie).Updates(&entities.Movie{Rate: rate})
	return result.Error
}

func (w WatchedListRepository) FindUserItemByID(itemId uint, userId uint) (item *entities.WatchedListItem, err error) {
	result := w.db.
		Preload("User").
		Preload("Movie").
		Where("`user_id` = ?", userId).
		First(&item, itemId)
	err = result.Error
	return
}

func (w WatchedListRepository) FindUserList(userId uint) (list *entities.WatchedList, err error) {
	result := w.db.
		Preload("User").
		Preload("Movie").
		Where("`user_id` = ?", userId).
		Find(&list)
	err = result.Error
	return
}
