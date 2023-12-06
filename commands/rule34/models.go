package rule34

import (
	"fmt"
	"strconv"
	"strings"
)

type Post struct {
	ID           int    `json:"id"`
	ParentID     int    `json:"parent_id"`
	PreviewURL   string `json:"preview_url"`
	SampleURL    string `json:"sample_url"`
	FileURL      string `json:"file_url"`
	Directory    int    `json:"directory"`
	Hash         string `json:"hash"`
	Height       int    `json:"height"`
	Width        int    `json:"width"`
	Image        string `json:"image"`
	Change       int    `json:"change"`
	Owner        string `json:"owner"`
	Rating       string `json:"rating"`
	Sample       int    `json:"sample"`
	SampleHeight int    `json:"sample_height"`
	SampleWidth  int    `json:"sample_width"`
	Score        int    `json:"score"`
	Tags         string `json:"tags"`
}

func (p Post) GetID() string {
	return strconv.Itoa(p.ID)
}

func (p Post) GetType() string {
	if strings.HasSuffix(p.FileURL, ".mp4") {
		return "video"
	}

	return "image"
}

func (p Post) GetTitle() string {
	return fmt.Sprintf("ID: %d", p.ID)
}

func (p Post) GetURL() string {
	return fmt.Sprintf("https://rule34.xxx/index.php?page=post&s=view&id=%d", p.ID)
}

func (p Post) GetMediaURL() string {
	if p.GetType() == "video" {
		return p.PreviewURL
	}

	return p.FileURL
}

type Tag struct {
	Label string `json:"label"`
	Value string `json:"value"`
	Type  string `json:"type"`
}

func (t Tag) GetValue() string {
	return t.Value
}

func (t Tag) GetName() string {
	return t.Label
}
