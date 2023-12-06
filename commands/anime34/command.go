package animer34

import (
	"github.com/bwmarrin/discordgo"
	"github.com/eoschaos/aniki-bot/commands"
)

var Command = &discordgo.ApplicationCommand{
	Name:        "ar34",
	Description: "Replies with an image from anime.rule34.world",
	Type:        discordgo.ChatApplicationCommand,
	Options: []*discordgo.ApplicationCommandOption{
		{
			Type:         discordgo.ApplicationCommandOptionString,
			Name:         "tags",
			Description:  "The tag to search with",
			Autocomplete: true,
		},
	},
}

func Handler() func(s *discordgo.Session, i *discordgo.InteractionCreate) {
	repo := NewRepository("https://anime.rule34.world")
	svc := commands.NewService(repo)

	return commands.HandlerCreator(svc)
}
