package bookcatalogue

import (
	"fmt"
	"strings"

	"github.com/Piokor/olutek_lib/internal/googleapi"
)

type Book struct {
	ID             string   `json:"api_id" db:"api_id"`
	Title          string   `json:"title" db:"title"`
	Subtitle       string   `json:"subtitle" db:"subtitle"`
	Authors        []string `json:"authors" db:"authors"`
	PublishDate    string   `json:"publishDate" db:"publish_date"`
	Description    string   `json:"description" db:"description"`
	PageCount      int64    `json:"pageCount" db:"page_count"`
	Publisher      string   `json:"publisher" db:"publisher"`
	Language       string   `json:"language" db:"language"`
	SmallThumbnail string   `json:"smallThumbnail" db:"small_thumbnail"`
	Thumbnail      string   `json:"thumbnail" db:"thumbnail"`
}

func (v Book) String() string {
	return fmt.Sprintf("%s by %s", v.Title, strings.Join(v.Authors, ", "))
}

func bookFromGoogleVolume(gv *googleapi.GoogleVolume) *Book {
	book := Book{
		ID:             gv.ID,
		Title:          gv.VolumeInfo.Title,
		Subtitle:       gv.VolumeInfo.Subtitle,
		Authors:        gv.VolumeInfo.Authors,
		PublishDate:    gv.VolumeInfo.PublishDate,
		Description:    gv.VolumeInfo.Description,
		PageCount:      gv.VolumeInfo.PageCount,
		Publisher:      gv.VolumeInfo.Publisher,
		Language:       gv.VolumeInfo.Language,
		SmallThumbnail: gv.VolumeInfo.ImageLinks.SmallThumbnail,
		Thumbnail:      gv.VolumeInfo.ImageLinks.Thumbnail,
	}
	return &book
}
