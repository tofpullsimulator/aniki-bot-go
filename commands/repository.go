package commands

import (
	"context"
	"io"
	"math/rand"
	"net/http"
	"net/url"
	"time"
)

type Post interface {
	GetID() string
	GetType() string
	GetTitle() string
	GetURL() string
	GetMediaURL() string
}

type PostResponse struct {
	Posts []Post
}

func (p PostResponse) GetPosts() []Post {
	return p.Posts
}

func (p PostResponse) GetRandomPost() Post {
	rand.Seed(time.Now().Unix())
	posts := p.GetPosts()
	n := rand.Int() % len(posts)

	return posts[n]
}

type Tag interface {
	GetValue() string
	GetName() string
}

type TagResponse struct {
	Tags []Tag
}

func (t TagResponse) GetTags() []Tag {
	return t.Tags
}

type Repository interface {
	GetPosts(ctx context.Context, tags, bannedTags []string) (*PostResponse, error)
	SearchTags(ctx context.Context, keyword string) (*TagResponse, error)
}

func DoGet(ctx context.Context, client *http.Client, endpoint string, params url.Values) (*http.Response, error) {
	var req *http.Request
	var err error
	if params == nil {
		req, err = http.NewRequest("GET", endpoint, nil)
	} else {
		req, err = http.NewRequest("GET", endpoint+"?"+params.Encode(), nil)
	}

	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req = req.WithContext(ctx)

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func DoPost(ctx context.Context, client *http.Client, endpoint string, body io.Reader) (*http.Response, error) {
	req, err := http.NewRequest("POST", endpoint, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req = req.WithContext(ctx)

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}
