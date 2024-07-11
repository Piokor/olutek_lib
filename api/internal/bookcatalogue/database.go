package bookcatalogue

import (
	"github.com/Piokor/olutek_lib/internal/database"
)

var bookDbFields []string = []string{
	"api_id",
	"title",
	"authors",
	"subtitle",
	"page_count",
	"publish_date",
	"description",
	"publisher",
	"language",
	"small_thumbnail",
	"thumbnail",
}

func InsertBook(book *Book, dbService *database.DbService) error {
	options := "ON CONFLICT (api_id) DO NOTHING"
	return dbService.InsertWithOptions("books", book, options, bookDbFields...)
}

func InsertBooks(books []*Book, dbService *database.DbService) error {
	options := "ON CONFLICT (api_id) DO NOTHING"
	return dbService.InsertWithOptions("books", books, options, bookDbFields...)
}

func SelectBook(bookApiId string, dbService *database.DbService) (*Book, error) {
	queryStr := "SELECT * FROM books WHERE api_id=$1"
	book := Book{}
	err := dbService.Client.Get(&book, queryStr, bookApiId)
	if err != nil {
		return nil, err
	}
	return &book, nil
}
