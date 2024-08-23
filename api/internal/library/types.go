package library

import (
	"time"

	"github.com/Piokor/olutek_lib/internal/bookcatalogue"
)

type BookStatus int64

const (
	notStarted BookStatus = iota
	inProgress BookStatus = iota
	completed  BookStatus = iota
)

type LibraryBook struct {
	Id                int64              `json:"id"`
	UserId            int64              `json:"user_id"`
	CatalogueBook     bookcatalogue.Book `json:"book"`
	Title             string             `json:"title"`
	Subtitle          string             `json:"subtitle"`
	Authors           []string           `json:"authors"`
	PublishDate       string             `json:"publish_date"`
	Description       string             `json:"description"`
	PageCount         int64              `json:"page_count"`
	Publisher         string             `json:"publisher"`
	Language          string             `json:"language"`
	SmallThumbnail    string             `json:"small_thumbnail"`
	Thumbnail         string             `json:"thumbnail"`
	Notes             string             `json:"notes"`
	Status            BookStatus         `json:"status"`
	AudiobookDuration time.Duration      `json:"audiobook_duraiton"`
}

type BookProgress struct {
	Id          int64       `json:"id"`
	LibraryBook LibraryBook `json:"book"`
	Date        time.Time   `json:"date"`
	Pages       int64       `json:"pages"`
	Notes       string      `json:"notes"`
}

type BookProgressSummary struct {
	BookTitle   string    `json:"book_title"`
	BookAuthors []string  `json:"book_authors"`
	BookId      int64     `json:"book_id"`
	Date        time.Time `json:"date"`
	Pages       int64     `json:"pages"`
	Notes       string    `json:"notes"`
}

type DailyReadingProgressSummary struct {
	Day           time.Time `json:"day"`
	BooksProgress []BookProgressSummary
}
