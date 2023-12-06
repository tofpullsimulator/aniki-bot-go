package gelbooru

import (
	"github.com/bwmarrin/discordgo"
	"github.com/eoschaos/aniki-bot/commands"
)

func Command(name string) *discordgo.ApplicationCommand {
	return &discordgo.ApplicationCommand{
		Name:        name,
		Description: "Replies with an image from gelbooru.com",
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
}

func Handler() func(s *discordgo.Session, i *discordgo.InteractionCreate) {
	repo := NewRepository("https://gelbooru.com")
	svc := commands.NewService(repo)

	return commands.HandlerCreator(svc)
}
