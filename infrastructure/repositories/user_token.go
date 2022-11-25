package repositories

import (
	"github.com/HazemNoor/movies-library/domain/entities"
	"github.com/HazemNoor/movies-library/infrastructure"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"time"
)

type UserTokenRepository struct {
	db *gorm.DB
}

func NewUserTokenRepository() UserTokenRepository {
	db, _ := infrastructure.DbConnection()
	return UserTokenRepository{db: db}
}

func (r UserTokenRepository) Create(userToken *entities.UserToken) error {
	result := r.db.Omit(clause.Associations).Create(userToken)
	return result.Error
}

func (r UserTokenRepository) FindByToken(token string) (userToken *entities.UserToken, err error) {
	result := r.db.
		Preload("User").
		Where("`token` = ?", token).
		Where("`expires_at` >= ?", time.Now()).
		First(&userToken)
	err = result.Error
	return
}
