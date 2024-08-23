package library

import (
	"time"

	"github.com/Piokor/olutek_lib/internal/database"
	"github.com/lib/pq"
)

func InsertLibraryBook(libraryBook *LibraryBook, dbService *database.DbService) error {
	fields := []string{
		"user_id",
		"catalogue_book_id",
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
		"notes",
		"status",
		"audiobook_duration",
	}
	query := dbService.BuildInsertQuery("library_books", fields, "")
	_, err := dbService.Client.Exec(query,
		libraryBook.UserId,
		libraryBook.CatalogueBook.DbId,
		libraryBook.Title,
		pq.Array(libraryBook.Authors),
		libraryBook.Subtitle,
		libraryBook.PageCount,
		libraryBook.PublishDate,
		libraryBook.Description,
		libraryBook.Publisher,
		libraryBook.Language,
		libraryBook.SmallThumbnail,
		libraryBook.Thumbnail,
		libraryBook.Notes,
		libraryBook.Status,
		libraryBook.AudiobookDuration,
	)
	return err
}

func SelectDailyReadingProgress(day time.Time, userId int64, dbService *database.DbService) (*DailyReadingProgressSummary, error) {
	query := `
	SELECT library_book.id, library_book.title, library_book.authors, book_progress.page_read, book_progress.notes,
	FROM book_progress,
	LEFT JOIN library_books ON book_progress.book_id = library_books.id
	WHERE library_books.user_id = $1 AND book_progress.reading_date = $2
	`

	rows, err := dbService.Client.Queryx(query, userId, day)
	if err != nil {
		return nil, err
	}

	dailyProgress := DailyReadingProgressSummary{
		day,
		make([]BookProgressSummary, 0),
	}

	for rows.Next() {
		bookProgress := BookProgressSummary{
			Date: day,
		}
		err = rows.Scan(&bookProgress.BookId, &bookProgress.BookTitle, &bookProgress.BookAuthors, &bookProgress.Pages, &bookProgress.Notes)
		if err != nil {
			return nil, err
		}
		dailyProgress.BooksProgress = append(dailyProgress.BooksProgress, bookProgress)
	}

	return &dailyProgress, nil
}
