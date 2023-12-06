package rule34

import (
	"github.com/bwmarrin/discordgo"
	"github.com/eoschaos/aniki-bot/commands"
)

var Command = &discordgo.ApplicationCommand{
	Name:        "r34",
	Description: "Replies with an image from rule34.xxx",
	Type:        discordgo.ChatApplicationCommand,
	Options: []*discordgo.ApplicationCommandOption{
		{
			Type:         discordgo.ApplicationCommandOptionString,
			Name:         "tags",
			Description:  "The tags to search with",
			Autocomplete: true,
		},
		{
			Type:        discordgo.ApplicationCommandOptionBoolean,
			Name:        "furries",
			Description: "Enable the furry tags, you dirty bastard",
		},
	},
}

func Handler() func(s *discordgo.Session, i *discordgo.InteractionCreate) {
	repo := NewRepository("https://api.rule34.xxx", "https://rule34.xxx")
	svc := commands.NewService(repo)

	return commands.HandlerCreator(svc)
}
