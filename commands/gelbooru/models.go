package gelbooru

import (
	"fmt"
	"strconv"
	"strings"
)

type Post struct {
	ID            int    `json:"id"`
	ParentID      int    `json:"parent_id"`
	CreatorID     int    `json:"creator_id"`
	Score         int    `json:"score"`
	Height        int    `json:"height"`
	Width         int    `json:"width"`
	MD5           string `json:"md5"`
	Directory     string `json:"directory"`
	Image         string `json:"image"`
	Rating        string `json:"rating"`
	Source        string `json:"source"`
	Change        int    `json:"change"`
	Owner         string `json:"owner"`
	Sample        int    `json:"sample"`
	PreviewHeight int    `json:"preview_height"`
	PreviewWidth  int    `json:"preview_width"`
	Tags          string `json:"tags"`
	Title         string `json:"title"`
	HasNotes      string `json:"has_notes"`
	HasComments   string `json:"has_comments"`
	FileURL       string `json:"file_url"`
	PreviewURL    string `json:"preview_url"`
	SampleURL     string `json:"sample_url"`
	SampleHeight  int    `json:"sample_height"`
	SampleWidth   int    `json:"sample_width"`
	Status        string `json:"status"`
	PostLocked    int    `json:"post_locked"`
	HasChildren   string `json:"has_children"`
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
	return fmt.Sprintf("https://gelbooru.com/index.php?page=post&s=view&id=%d", p.ID)
}

func (p Post) GetMediaURL() string {
	if p.GetType() == "video" {
		return p.PreviewURL
	}

	return p.FileURL
}

type Attributes struct {
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
	Count  int `json:"count"`
}

type PostsResponse struct {
	Attributes Attributes `json:"@attributes"`
	Posts      []Post     `json:"post"`
}

type Tag struct {
	Type      string `json:"tag"`
	Label     string `json:"label"`
	Value     string `json:"value"`
	PostCount string `json:"post_count"`
	Category  string `json:"category"`
}

func (t Tag) GetValue() string {
	return t.Value
}

func (t Tag) GetName() string {
	return fmt.Sprintf("%s (%s)", t.GetValue(), t.PostCount)
}
