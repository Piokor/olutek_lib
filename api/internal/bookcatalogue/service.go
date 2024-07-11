package bookcatalogue

import (
	"log"

	"github.com/Piokor/olutek_lib/internal/database"
	"github.com/Piokor/olutek_lib/internal/googleapi"
)

func getBookFromApi(bookApiId string, dbService *database.DbService) (*Book, error) {
	apiResult, err := googleapi.GetBook(bookApiId)
	if err != nil {
		log.Printf("Book %s not found in the API", bookApiId)
		return nil, err
	}
	book := bookFromGoogleVolume(apiResult)
	defer InsertBook(book, dbService)
	return book, nil
}

func GetBook(bookApiId string, dbService *database.DbService) (*Book, error) {
	dbResult, err := SelectBook(bookApiId, dbService)
	if err != nil {
		log.Printf("Book %s not found in the DB, fetching from API...", bookApiId)
		return getBookFromApi(bookApiId, dbService)
	}
	return dbResult, nil
}

func GetBooks(query *googleapi.BookQuery, dbService *database.DbService) ([]*Book, error) {
	apiResults, err := googleapi.GetBooks(query)
	if err != nil {
		return nil, err
	}
	books := make([]*Book, len(apiResults))
	for i, apiResult := range apiResults {
		books[i] = bookFromGoogleVolume(apiResult)
	}
	defer InsertBooks(books, dbService)
	return books, nil
}
