package commands

import (
	"context"
	"log"
	"os"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/eoschaos/aniki-bot/utils"
)

func HandlerCreator(svc Service) func(s *discordgo.Session, i *discordgo.InteractionCreate) {
	return func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		options := i.ApplicationCommandData().Options
		optionMap := make(map[string]*discordgo.ApplicationCommandInteractionDataOption, len(options))
		for _, opt := range options {
			optionMap[opt.Name] = opt
		}

		var tags []string
		if option, ok := optionMap["tags"]; ok {
			tags = strings.Split(option.StringValue(), ",")
		}

		var furries bool
		if option, ok := optionMap["furries"]; ok {
			furries = option.BoolValue()
		}
		ctx := context.Background()

		switch i.Type {
		case discordgo.InteractionApplicationCommand:
			var bannedTags []string
			if !furries {
				furryTags := os.Getenv("FURRY_TAGS")
				bannedTags = strings.Split(furryTags, ",")
			}

			messageEmbed, err := svc.GetRandomPost(ctx, tags, bannedTags)
			if err != nil {
				log.Println(err)
				utils.SendEmbedArticle(s, i, "Something went wrong with getting an image!")
				return
			}

			err = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Embeds: []*discordgo.MessageEmbed{messageEmbed},
				},
			})
			if err != nil {
				log.Println(err)
			}
		case discordgo.InteractionApplicationCommandAutocomplete:
			choices, err := svc.GetAutocompleteChoices(ctx, tags[0])
			if err != nil {
				log.Println(err)
			}

			err = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionApplicationCommandAutocompleteResult,
				Data: &discordgo.InteractionResponseData{
					Choices: choices,
				},
			})
			if err != nil {
				log.Println(err)
			}
		}
	}
}
