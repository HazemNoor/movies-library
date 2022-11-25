package repositories

import (
	"github.com/HazemNoor/movies-library/domain/entities"
	"github.com/HazemNoor/movies-library/infrastructure"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type MovieRepository struct {
	db *gorm.DB
}

func NewMovieRepository() MovieRepository {
	db, _ := infrastructure.DbConnection()
	return MovieRepository{db: db}
}

func (r MovieRepository) FindAll() (movies *entities.Movies, err error) {
	result := r.db.Find(&movies)
	err = result.Error
	return
}

func (r MovieRepository) FindAllOrdered(orderBy string) (movies *entities.Movies, err error) {
	result := r.db.Order(orderBy).Find(&movies)
	err = result.Error
	return
}

func (r MovieRepository) FindByID(movieId uint) (movie *entities.Movie, err error) {
	result := r.db.Preload("User").First(&movie, movieId)
	err = result.Error
	return
}

func (r MovieRepository) Create(movie *entities.Movie) error {
	result := r.db.Omit(clause.Associations).Create(movie)
	return result.Error
}

func (r MovieRepository) Update(movie *entities.Movie, data *entities.Movie) error {
	result := r.db.Omit(clause.Associations).Model(&movie).Updates(data)
	return result.Error
}

func (r MovieRepository) Delete(movie *entities.Movie) error {
	result := r.db.Delete(&movie)
	return result.Error
}
