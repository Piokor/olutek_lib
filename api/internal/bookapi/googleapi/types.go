package googleapi

import (
	"github.com/Piokor/olutek_lib/internal/bookapi"
)

type GoogleVolumeResults struct {
	TotalItems int64          `json:"totalItems"`
	Items      []GoogleVolume `json:"items"`
}

type GoogleVolume struct {
	ID         string           `json:"id"`
	VolumeInfo GoogleVolumeInfo `json:"volumeInfo"`
}

type GoogleVolumeInfo struct {
	Title         string           `json:"title"`
	Subtitle      string           `json:"subtitle"`
	Authors       []string         `json:"authors"`
	PublishedDate string           `json:"publishedDate"`
	Description   string           `json:"description"`
	PageCount     int64            `json:"pageCount"`
	Publisher     string           `json:"publisher"`
	Language      string           `json:"language"`
	ImageLinks    GoogleImageLinks `json:"imageLinks"`
}

type GoogleImageLinks struct {
	SmallThumbnail string `json:"smallThumbnail"`
	Thumbnail      string `json:"thumbnail"`
}

func (gv GoogleVolume) ToVolume() bookapi.Volume {
	return bookapi.Volume{
		ID:          gv.ID,
		Title:       gv.VolumeInfo.Title,
		Subtitle:    gv.VolumeInfo.Subtitle,
		Authors:     gv.VolumeInfo.Authors,
		PublishDate: gv.VolumeInfo.PublishedDate,
		Description: gv.VolumeInfo.Description,
		PageCount:   gv.VolumeInfo.PageCount,
		Publisher:   gv.VolumeInfo.Publisher,
		Language:    gv.VolumeInfo.Language,
		Image: bookapi.VolumeImages{
			SmallThumbnail: gv.VolumeInfo.ImageLinks.SmallThumbnail,
			Thumbnail:      gv.VolumeInfo.ImageLinks.Thumbnail,
		},
	}
}
