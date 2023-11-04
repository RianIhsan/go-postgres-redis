package repository

import (
	"errors"
	"github.com/RianIhsan/go-postgres-redis/domain"
	"github.com/RianIhsan/go-postgres-redis/model"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type bookRepository struct {
	db  *gorm.DB
	rdb *redis.Client
}

func (b *bookRepository) CreateBook(createBook model.Book) error {
	if err := b.db.Create(&createBook).Error; err != nil {
		return errors.New("internal server error: cannot create book")
	}
	return nil
}

func NewBookRepository(db *gorm.DB, rdb *redis.Client) domain.BookRepo {
	return &bookRepository{
		db:  db,
		rdb: rdb,
	}
}
