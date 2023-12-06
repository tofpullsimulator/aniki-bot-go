package r34xyz

import (
	"fmt"
	"strconv"
)

type Post struct {
	Posted     string        `json:"posted"`
	PostSearch string        `json:"postSearch"`
	Likes      int           `json:"likes"`
	Views      int           `json:"views"`
	Comments   int           `json:"comments"`
	Type       int           `json:"type"`
	Status     int           `json:"status"`
	UploaderID int           `json:"uploaderId"`
	Uploader   string        `json:"uploader"`
	Duration   int           `json:"duration"`
	Error      string        `json:"error"`
	Width      int           `json:"width"`
	Height     int           `json:"height"`
	Files      map[int][]int `json:"files"`
	Tags       []string      `json:"tags"`
	FullTags   []string      `json:"fullTags"`
	Cursor     string        `json:"cursor"`
	Sources    []string      `json:"sources"`
	ID         int           `json:"id"`
	Created    string        `json:"created"`
}

func (p Post) GetID() string {
	return strconv.Itoa(p.ID)
}

func (p Post) GetType() string {
	if p.Type == 1 {
		return "video"
	}

	return "image"
}

func (p Post) GetTitle() string {
	return fmt.Sprintf("ID: %d", p.ID)
}

func (p Post) GetURL() string {
	return fmt.Sprintf("https://r-34.xyz/post/%d", p.ID)
}

func (p Post) GetMediaURL() string {
	id := p.GetID()
	prefix := id[0:2]
	if len(id) == 4 {
		prefix = id[0:1]
	}

	if p.GetType() == "video" {
		return fmt.Sprintf("https://r34xyz.b-cdn.net/posts/%s/%d/%d.thumbnail.jpg", prefix, p.ID, p.ID)
	}

	return fmt.Sprintf("https://r34xyz.b-cdn.net/posts/%s/%d/%d.jpg", prefix, p.ID, p.ID)
}

type PostsResponse struct {
	Items  []*Post `json:"items"`
	Cursor int     `json:"cursor"`
}

type Tag struct {
	ID    int    `json:"id"`
	Value string `json:"value"`
	Count int    `json:"count"`
	Type  string `json:"type"`
}

func (t Tag) GetValue() string {
	return t.Value
}

func (t Tag) GetName() string {
	return fmt.Sprintf("%s (%d)", t.GetValue(), t.Count)
}
