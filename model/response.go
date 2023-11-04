package model

type response struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func GeneralResponse(message string, data interface{}) response {
	messageRes := response{
		Message: message,
		Data:    data,
	}

	return messageRes
}

type CreateBookFormatter struct {
	Id          int    `json:"id"`
	Book        string `json:"book"`
	Description string `json:"description"`
	Author      string `json:"author"`
}

func FormatBook(book Book) CreateBookFormatter {
	bookFormatter := CreateBookFormatter{}
	bookFormatter.Id = book.Id
	bookFormatter.Book = book.Book
	bookFormatter.Description = book.Description
	bookFormatter.Author = book.Author

	return bookFormatter
}
