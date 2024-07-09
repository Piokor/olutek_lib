package googleapi

import "github.com/Piokor/olutek_lib/internal"

type googleVolumeResults struct {
	TotalItems int64          `json:"totalItems"`
	Items      []googleVolume `json:"items"`
}

type googleVolume struct {
	ID         string           `json:"id"`
	VolumeInfo googleVolumeInfo `json:"volumeInfo"`
}

type googleVolumeInfo struct {
	Title       string           `json:"title"`
	Subtitle    string           `json:"subtitle"`
	Authors     []string         `json:"authors"`
	PublishDate string           `json:"publishedDate"`
	Description string           `json:"description"`
	PageCount   int64            `json:"pageCount"`
	Publisher   string           `json:"publisher"`
	Language    string           `json:"language"`
	ImageLinks  googleImageLinks `json:"imageLinks"`
}

type googleImageLinks struct {
	SmallThumbnail string `json:"smallThumbnail"`
	Thumbnail      string `json:"thumbnail"`
}

func (gv googleVolume) toBook() *internal.Book {
	book := internal.Book{
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
