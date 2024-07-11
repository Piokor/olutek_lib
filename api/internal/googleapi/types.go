package googleapi

type GoogleVolumeResults struct {
	TotalItems int64          `json:"totalItems"`
	Items      []GoogleVolume `json:"items"`
}

type GoogleVolume struct {
	ID         string           `json:"id"`
	VolumeInfo GoogleVolumeInfo `json:"volumeInfo"`
}

type GoogleVolumeInfo struct {
	Title       string           `json:"title"`
	Subtitle    string           `json:"subtitle"`
	Authors     []string         `json:"authors"`
	PublishDate string           `json:"publishedDate"`
	Description string           `json:"description"`
	PageCount   int64            `json:"pageCount"`
	Publisher   string           `json:"publisher"`
	Language    string           `json:"language"`
	ImageLinks  GoogleImageLinks `json:"imageLinks"`
}

type GoogleImageLinks struct {
	SmallThumbnail string `json:"smallThumbnail"`
	Thumbnail      string `json:"thumbnail"`
}
