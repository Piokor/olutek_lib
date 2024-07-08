package bookapi

import (
	"fmt"
	"strings"
)

type Volume struct {
	ID              string
	Title, Subtitle string
	Authors         []string
	PublishDate     string
	Description     string
	PageCount       int64
	Publisher       string
	Image           VolumeImages
	Language        string
}

type VolumeImages struct {
	SmallThumbnail string
	Thumbnail      string
}

func (v Volume) String() string {
	return fmt.Sprintf("%s by %s", v.Title, strings.Join(v.Authors, ", "))
}
