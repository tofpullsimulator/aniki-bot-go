package rule34

import (
	"context"
	"encoding/json"
	"github.com/eoschaos/aniki-bot/commands"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type repository struct {
	client     *http.Client
	serviceURL string
	publicURL  string
}

func NewRepository(serviceURL, publicURL string) commands.Repository {
	client := &http.Client{}
	return &repository{
		client:     client,
		serviceURL: serviceURL,
		publicURL:  publicURL,
	}
}

func (r *repository) GetPosts(ctx context.Context, tags, bannedTags []string) (*commands.PostResponse, error) {
	var banned []string
	for _, v := range bannedTags {
		banned = append(banned, "-"+v)
	}

	params := url.Values{
		"limit": {"1000"},
		"page":  {"dapi"},
		"json":  {"1"},
		"s":     {"post"},
		"q":     {"index"},
	}
	requestTags := strings.Join(append(tags, banned...), "+")

	res, err := commands.DoGet(ctx, r.client, r.serviceURL+"/index.php?"+params.Encode()+"&tags="+requestTags, nil)
	if err != nil {
		return nil, err
	}

	var p []Post
	err = json.NewDecoder(res.Body).Decode(&p)
	if err != nil && err != io.EOF {
		return nil, err
	}

	var posts []commands.Post
	for _, v := range p {
		posts = append(posts, v)
	}

	return &commands.PostResponse{Posts: posts}, nil
}

func (r *repository) SearchTags(ctx context.Context, keyword string) (*commands.TagResponse, error) {
	params := url.Values{
		"q": {keyword},
	}
	res, err := commands.DoGet(ctx, r.client, r.publicURL+"/public/autocomplete.php", params)
	if err != nil {
		return nil, err
	}

	var t []Tag
	if err := json.NewDecoder(res.Body).Decode(&t); err != nil {
		return nil, err
	}

	var tags []commands.Tag
	for _, v := range t {
		tags = append(tags, v)
	}

	return &commands.TagResponse{Tags: tags}, nil
}
