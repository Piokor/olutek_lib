package internal

import (
	"fmt"
	"strings"
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
