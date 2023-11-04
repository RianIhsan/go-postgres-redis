package domain

import (
	"github.com/RianIhsan/go-postgres-redis/entities"
	"github.com/RianIhsan/go-postgres-redis/model"
)

type BookRepo interface {
	CreateBook(createBook model.Book) error
}

type BookUseCase interface {
	CreateBook(input entities.BookInput) (model.Book, error)
}
