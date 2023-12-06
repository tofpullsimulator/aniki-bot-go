package r34xyz

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"

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
	inter := map[string]any{
		"cursor":         nil,
		"excludeTags":    nil,
		"includeTags":    tags,
		"postedFromDays": nil,
		"skipCache":      false,
		"sortBy":         0,
		"sortOrder":      1,
		"status":         2,
		"take":           1000,
		"type":           nil,
	}
	postBody, _ := json.Marshal(inter)
	body := bytes.NewBuffer(postBody)

	res, err := commands.DoPost(ctx, r.client, r.serviceURL+"/api/post/search/root", body)
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
	res, err := commands.DoGet(ctx, r.client, r.serviceURL+"/api/tags/search/"+keyword, nil)
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
