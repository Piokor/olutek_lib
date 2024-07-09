package database

import (
	"fmt"
	"strings"

	"github.com/Piokor/olutek_lib/internal"
)

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

	return fmt.Sprintf("INSERT INTO books (%s) VALUES (%s) ON CONFLICT (api_id) DO NOTHING",
		fieldsStr,
		placeholdersStr,
	)
}

func (service *DbService) InsertBook(book *internal.Book) error {
	insertString := getBookInsertString()
	dbBook := toDbBook(book)
	_, err := service.client.NamedExec(insertString, &dbBook)
	return err
}

func (service *DbService) InsertBooks(books []*internal.Book) error {
	insertString := getBookInsertString()
	dbBooks := make([]dbBook, len(books))
	for i, book := range books {
		dbBooks[i] = toDbBook(book)
	}
	_, err := service.client.NamedExec(insertString, dbBooks)
	return err
}

func (service *DbService) GetBook(bookApiId string) (*internal.Book, error) {
	queryStr := "SELECT * FROM books WHERE api_id=$1"
	dbBook := dbBook{}
	err := service.client.Get(&dbBook, queryStr, bookApiId)
	if err != nil {
		return nil, err
	}
	book := dbBook.toBook()
	return &book, nil
}
