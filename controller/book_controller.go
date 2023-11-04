package controller

import (
	"errors"
	"github.com/RianIhsan/go-postgres-redis/domain"
	"github.com/RianIhsan/go-postgres-redis/entities"
	"github.com/RianIhsan/go-postgres-redis/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

type BookController struct {
	bookUsecase domain.BookUseCase
}

func NewBookController(bookUsecase domain.BookUseCase) *BookController {
	return &BookController{bookUsecase: bookUsecase}
}

func (b *BookController) CreateBook(c *gin.Context) {
	var bookRequest entities.BookInput

	if err := c.ShouldBindJSON(&bookRequest); err != nil {
		response := model.GeneralResponse("Invalid payload", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}
	if bookRequest.Book == "" || bookRequest.Author == "" || bookRequest.Description == "" {
		response := model.GeneralResponse("Payload missing", errors.New("payload cannot be empty"))
		c.JSON(http.StatusBadRequest, response)
		return
	}

	book, err := b.bookUsecase.CreateBook(bookRequest)
	if err != nil {
		response := model.GeneralResponse("Internal server error: ", err.Error())
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	response := model.GeneralResponse("Successfully created book", model.FormatBook(book))
	c.JSON(http.StatusCreated, response)
}
