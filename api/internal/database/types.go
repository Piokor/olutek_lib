package database

import (
	"github.com/Piokor/olutek_lib/internal"
	"github.com/lib/pq"
)

type dbBook struct {
	ID             string         `db:"id"`
	ApiId          string         `db:"api_id"`
	Title          string         `db:"title"`
	Subtitle       string         `db:"subtitle"`
	Authors        pq.StringArray `db:"authors"`
	PublishDate    string         `db:"publish_date"`
	Description    string         `db:"description"`
	PageCount      int64          `db:"page_count"`
	Publisher      string         `db:"publisher"`
	Language       string         `db:"language"`
	SmallThumbnail string         `db:"small_thumbnail"`
	Thumbnail      string         `db:"thumbnail"`
}

func (dbv dbBook) toBook() internal.Book {
	return internal.Book{
		ID:             dbv.ApiId,
		Title:          dbv.Title,
		Subtitle:       dbv.Subtitle,
		Authors:        dbv.Authors,
		PublishDate:    dbv.PublishDate,
		Description:    dbv.Description,
		PageCount:      dbv.PageCount,
		Publisher:      dbv.Publisher,
		Language:       dbv.Language,
		SmallThumbnail: dbv.SmallThumbnail,
		Thumbnail:      dbv.Thumbnail,
	}
}

func toDbBook(v *internal.Book) dbBook {
	return dbBook{
		ApiId:          v.ID,
		Title:          v.Title,
		Subtitle:       v.Subtitle,
		Authors:        v.Authors,
		PublishDate:    v.PublishDate,
		Description:    v.Description,
		PageCount:      v.PageCount,
		Publisher:      v.Publisher,
		Language:       v.Language,
		SmallThumbnail: v.SmallThumbnail,
		Thumbnail:      v.Thumbnail,
	}
}
