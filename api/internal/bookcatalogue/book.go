package bookcatalogue

import (
	"fmt"
	"strings"
)

type Book struct {
	ID             string
	Title          string
	Subtitle       string
	Authors        []string
	PublishDate    string
	Description    string
	PageCount      int64
	Publisher      string
	Language       string
	SmallThumbnail string
	Thumbnail      string
}

func (v Book) String() string {
	return fmt.Sprintf("%s by %s", v.Title, strings.Join(v.Authors, ", "))
}
