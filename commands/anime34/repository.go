package animer34

import (
	"bytes"
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
		"IncludeLinks": {"true"},
		"OrderBy":      {"0"},
		"Skip":         {"0"},
		"Take":         {"100"},
		"DisableTotal": {"false"},
	}
	if len(tags) > 0 {
		params.Add("Tag", tags[0])
	}

	res, err := commands.DoGet(ctx, r.client, r.serviceURL+"/api/post/search-light", params)
	if err != nil {
		return nil, err
	}

	var p PostsResponse
	if err := json.NewDecoder(res.Body).Decode(&p); err != nil {
		return nil, err
	}

	var posts []commands.Post
	for _, v := range p.Items {
		posts = append(posts, v)
	}

	return &commands.PostResponse{Posts: posts}, nil
}

func (r *repository) SearchTags(ctx context.Context, keyword string) (*commands.TagResponse, error) {
	inter := map[string]string{
		"searchText": keyword,
		"skip":       "0",
		"take":       "20",
	}
	postBody, _ := json.Marshal(inter)
	body := bytes.NewBuffer(postBody)

	res, err := commands.DoPost(ctx, r.client, r.serviceURL+"/api/tag/Search", body)
	if err != nil {
		return nil, err
	}

	var t TagsResponse
	if err := json.NewDecoder(res.Body).Decode(&t); err != nil {
		return nil, err
	}

	var tags []commands.Tag
	for _, v := range t.Items {
		tags = append(tags, v)
	}

	return &commands.TagResponse{Tags: tags}, nil
}
