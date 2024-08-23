package library

import (
	"time"

	"github.com/Piokor/olutek_lib/internal/bookcatalogue"
	"github.com/Piokor/olutek_lib/internal/database"
)

func bookToLibraryBook(book *bookcatalogue.Book) *LibraryBook {
	return &(LibraryBook{
		CatalogueBook:  *book,
		Title:          book.Title,
		Subtitle:       book.Subtitle,
		Authors:        book.Authors,
		PublishDate:    book.PublishDate,
		Description:    book.Description,
		PageCount:      book.PageCount,
		Publisher:      book.Publisher,
		Language:       book.Language,
		SmallThumbnail: book.SmallThumbnail,
		Thumbnail:      book.Thumbnail,
	})
}

func AddBookToLibrary(bookApiId string, userId int64, dbService *database.DbService) error {
	book, err := bookcatalogue.GetBook(bookApiId, dbService)
	if err != nil {
		return err
	}
	libraryBook := bookToLibraryBook(book)
	libraryBook.Status = notStarted
	libraryBook.Notes = ""
	libraryBook.UserId = userId
	return InsertLibraryBook(libraryBook, dbService)
}

func GetDailyReadingProgress(day time.Time, userId int64, dbService *database.DbService) (*DailyReadingProgressSummary, error) {
	return SelectDailyReadingProgress(day, userId, dbService)
}
