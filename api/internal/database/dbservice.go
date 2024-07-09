package database

import (
	"fmt"
	"strings"

	"github.com/Piokor/olutek_lib/internal/bookcatalogue"
	"github.com/jmoiron/sqlx"
)

type DbService struct {
	client *sqlx.DB
}

func NewDbService() (*DbService, error) {
	dbConnection, err := GetClient()
	if err != nil {
		return nil, err
	}
	service := DbService{dbConnection}
	return &service, nil
}

func getBookInsertString() string {
	fields := []string{
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
	fieldsStr := strings.Join(fields, ", ")

	placeholders := make([]string, len(fields))
	for i, field := range fields {
		placeholders[i] = fmt.Sprintf(":%s", field)
	}
	placeholdersStr := strings.Join(placeholders, ", ")

	return fmt.Sprintf("INSERT INTO books (%s) VALUES (%s) ON CONFLICT (api_id) DO NOTHING", fieldsStr, placeholdersStr)
}

func (service *DbService) InsertBook(book bookcatalogue.Book) error {
	insertString := getBookInsertString()
	dbBook := toDbBook(book)
	_, err := service.client.NamedExec(insertString, &dbBook)
	return err
}

func (service *DbService) InsertBooks(books []bookcatalogue.Book) error {
	insertString := getBookInsertString()
	dbBooks := make([]DbBook, len(books))
	for i, book := range books {
		dbBooks[i] = toDbBook(book)
	}
	_, err := service.client.NamedExec(insertString, dbBooks)
	return err
}
