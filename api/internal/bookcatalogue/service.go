package bookcatalogue

import (
	"log"

	"github.com/Piokor/olutek_lib/internal"
	"github.com/Piokor/olutek_lib/internal/database"
	"github.com/Piokor/olutek_lib/internal/googleapi"
)

func getBookFromApi(bookApiId string, dbService *database.DbService) (*internal.Book, error) {
	apiResult, err := googleapi.GetBook(bookApiId)
	if err != nil {
		log.Printf("Book %s not found in the API", bookApiId)
		return nil, err
	}
	defer dbService.InsertBook(apiResult)
	return apiResult, nil
}

func GetBook(bookApiId string, dbService *database.DbService) (*internal.Book, error) {
	dbResult, err := dbService.GetBook(bookApiId)
	if err != nil {
		log.Printf("Book %s not found in the DB, fetching from API...", bookApiId)
		return getBookFromApi(bookApiId, dbService)
	}
	return dbResult, nil
}

func GetBooks(query *googleapi.BookQuery, dbService *database.DbService) ([]*internal.Book, error) {
	apiResults, err := googleapi.GetBooks(query)
	if err != nil {
		return nil, err
	}
	defer dbService.InsertBooks(apiResults)
	return apiResults, nil
}
