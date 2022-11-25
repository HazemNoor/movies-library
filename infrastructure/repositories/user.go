package repositories

import (
	"github.com/HazemNoor/movies-library/domain/entities"
	"github.com/HazemNoor/movies-library/infrastructure"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository() UserRepository {
	db, _ := infrastructure.DbConnection()
	return UserRepository{db: db}
}

func (r UserRepository) FindAll() (users *entities.Users, err error) {
	result := r.db.Find(&users)
	err = result.Error
	return
}

func (r UserRepository) FindByID(userID uint) (user *entities.User, err error) {
	result := r.db.First(&user, userID)
	err = result.Error
	return
}

func (r UserRepository) FindByEmail(email string) (user *entities.User, err error) {
	result := r.db.First(&user, "email = ?", email)
	err = result.Error
	return
}

func (r UserRepository) Create(user *entities.User) error {
	result := r.db.Create(user)
	return result.Error
}
