package bookcatalogue

import (
	"fmt"
	"strings"

	"github.com/Piokor/olutek_lib/internal/googleapi"
)

type Book struct {
	ID             string   `json:"api_id"`
	Title          string   `json:"title"`
	Subtitle       string   `json:"subtitle"`
	Authors        []string `json:"authors"`
	PublishDate    string   `json:"publishDate"`
	Description    string   `json:"description"`
	PageCount      int64    `json:"pageCount"`
	Publisher      string   `json:"publisher"`
	Language       string   `json:"language"`
	SmallThumbnail string   `json:"smallThumbnail"`
	Thumbnail      string   `json:"thumbnail"`
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
