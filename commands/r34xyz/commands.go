package r34xyz

import (
	"github.com/bwmarrin/discordgo"
	"github.com/eoschaos/aniki-bot/commands"
)

var Command = &discordgo.ApplicationCommand{
	Name:        "xyz34",
	Description: "Replies with an image from r-34.xyz",
	Type:        discordgo.ChatApplicationCommand,
	Options: []*discordgo.ApplicationCommandOption{
		{
			Type:         discordgo.ApplicationCommandOptionString,
			Name:         "tags",
			Description:  "The tags to search with",
			Autocomplete: true,
		},
	},
}

func Handler() func(s *discordgo.Session, i *discordgo.InteractionCreate) {
	repo := NewRepository("https://r-34.xyz")
	svc := commands.NewService(repo)

	return commands.HandlerCreator(svc)
}
