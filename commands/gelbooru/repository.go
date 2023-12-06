package gelbooru

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/eoschaos/aniki-bot/commands"
)

type repository struct {
	client     *http.Client
	serviceURL string
}

func NewRepository(serviceURL string) commands.Repository {
	client := &http.Client{}
	return &repository{
		client:     client,
		serviceURL: serviceURL,
	}
}

func (r *repository) GetPosts(ctx context.Context, tags, bannedTags []string) (*commands.PostResponse, error) {
	params := url.Values{
		"limit": {"100"},
		"page":  {"dapi"},
		"json":  {"1"},
		"s":     {"post"},
		"q":     {"index"},
		"tags":  tags,
	}

	res, err := commands.DoGet(ctx, r.client, r.serviceURL+"/index.php", params)
	if err != nil {
		return nil, err
	}

	var p PostsResponse
	if err := json.NewDecoder(res.Body).Decode(&p); err != nil {
		return nil, err
	}

	var posts []commands.Post
	for _, v := range p.Posts {
		posts = append(posts, v)
	}

	return &commands.PostResponse{Posts: posts}, nil
}

func (r *repository) SearchTags(ctx context.Context, keyword string) (*commands.TagResponse, error) {
	params := url.Values{
		"page":  {"autocomplete2"},
		"term":  {keyword},
		"type":  {"tag_query"},
		"limit": {"20"},
	}
	res, err := commands.DoGet(ctx, r.client, r.serviceURL+"/index.php", params)
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
