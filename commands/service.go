package commands

import (
	"context"
	"fmt"
	"regexp"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/eoschaos/aniki-bot/utils"
)

type Service interface {
	GetRandomPost(ctx context.Context, tags, bannedTags []string) (*discordgo.MessageEmbed, error)
	GetAutocompleteChoices(ctx context.Context, keyword string) ([]*discordgo.ApplicationCommandOptionChoice, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{
		repository: repository,
	}
}

func (s *service) GetRandomPost(ctx context.Context, tags, bannedTags []string) (*discordgo.MessageEmbed, error) {
	for i, e := range tags {
		exp := regexp.MustCompile(` \(\d+\)`)
		tags[i] = exp.ReplaceAllString(e, "")
	}

	res, err := s.repository.GetPosts(ctx, tags, bannedTags)
	if err != nil {
		return nil, err
	}

	if len(res.GetPosts()) == 0 {
		description := fmt.Sprintf("No posts found for tags: %s", strings.Join(tags, ", "))
		return utils.CreateEmbedArticle(description), nil
	}

	description := "Tags used: none"
	if len(tags) > 0 {
		description = fmt.Sprintf("Tags used: %s", strings.Join(tags, ", "))
	}

	p := res.GetRandomPost()
	if p.GetType() == "video" {
		return &discordgo.MessageEmbed{
			Title:       p.GetTitle(),
			URL:         p.GetURL(),
			Description: description,
			Thumbnail: &discordgo.MessageEmbedThumbnail{
				URL: p.GetMediaURL(),
			},
		}, nil
	}

	return &discordgo.MessageEmbed{
		Title:       p.GetTitle(),
		URL:         p.GetURL(),
		Description: description,
		Image: &discordgo.MessageEmbedImage{
			URL: p.GetMediaURL(),
		},
	}, nil
}

func (s *service) GetAutocompleteChoices(ctx context.Context, keyword string) ([]*discordgo.ApplicationCommandOptionChoice, error) {
	res, err := s.repository.SearchTags(ctx, keyword)
	if err != nil {
		return nil, err
	}

	tags := res.GetTags()
	choices := make([]*discordgo.ApplicationCommandOptionChoice, len(tags))
	for k, v := range tags {
		choices[k] = &discordgo.ApplicationCommandOptionChoice{
			Value: v.GetValue(),
			Name:  v.GetName(),
		}
	}

	return choices, nil
}
