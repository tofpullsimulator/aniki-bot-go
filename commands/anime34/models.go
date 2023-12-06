package animer34

import (
	"fmt"
	"strconv"
	"strings"

	"golang.org/x/exp/slices"
)

type ImageLink struct {
	Type int    `json:"type"`
	URL  string `json:"url"`
}

type Post struct {
	ID         int         `json:"id"`
	Likes      int         `json:"likes"`
	Views      int         `json:"views"`
	Comments   int         `json:"comments"`
	Type       int         `json:"type"`
	Duration   string      `json:"duration"`
	Posted     string      `json:"posted"`
	ImageLinks []ImageLink `json:"imageLinks"`
	Tags       []string    `json:"tags"`
}

func (p Post) GetID() string {
	return strconv.Itoa(p.ID)
}

func (p Post) GetType() string {
	if p.Type == 1 {
		return "video"
	}

	return "string"
}

func (p Post) GetTitle() string {
	return fmt.Sprintf("ID: %d", p.ID)
}

func (p Post) GetURL() string {
	return fmt.Sprintf("https://anime.rule34.world/post/%d", p.ID)
}

func (p Post) GetMediaURL() string {
	idx := slices.IndexFunc(p.ImageLinks, func(imageLink ImageLink) bool {
		return imageLink.Type == 3
	})

	imageLink := p.ImageLinks[idx]
	imageLink.URL = strings.ReplaceAll(imageLink.URL, "pic256", "pic")
	if !strings.HasPrefix(imageLink.URL, "http") {
		imageLink.URL = fmt.Sprintf("https://anime-rule34-world.b-cdn.net%s", imageLink.URL)
	}

	return imageLink.URL
}

type PostsResponse struct {
	Items      []*Post `json:"items"`
	Skip       int     `json:"skip"`
	TotalCount int     `json:"totalCount"`
}

type Tag struct {
	ID    int    `json:"id"`
	Value string `json:"value"`
	Count int    `json:"count"`
	Type  int    `json:"type"`
}

func (t Tag) GetValue() string {
	return t.Value
}

func (t Tag) GetName() string {
	return fmt.Sprintf("%s (%d)", t.GetValue(), t.Count)
}

type TagsResponse struct {
	Items []*Tag `json:"items"`
}
