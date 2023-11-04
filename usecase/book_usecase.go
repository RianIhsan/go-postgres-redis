package usecase

import (
	"github.com/RianIhsan/go-postgres-redis/domain"
	"github.com/RianIhsan/go-postgres-redis/entities"
	"github.com/RianIhsan/go-postgres-redis/model"
)

type bookUsecase struct {
	bookRepository domain.BookRepo
}

func NewBookUsecase(bookRepository domain.BookRepo) domain.BookUseCase {
	return &bookUsecase{
		bookRepository: bookRepository,
	}
}

func (b *bookUsecase) CreateBook(input entities.BookInput) (model.Book, error) {
	book := model.Book{}
	book.Book = input.Book
	book.Description = input.Description
	book.Author = input.Author
	err := b.bookRepository.CreateBook(book)
	if err != nil {
		return book, err
	}
	return book, err
}
