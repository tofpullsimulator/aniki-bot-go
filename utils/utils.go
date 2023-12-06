package utils

import "github.com/bwmarrin/discordgo"

func CreateEmbedArticle(description string) *discordgo.MessageEmbed {
	return &discordgo.MessageEmbed{
		Description: description,
		Type:        discordgo.EmbedTypeRich,
	}
}

func SendEmbedArticle(s *discordgo.Session, i *discordgo.InteractionCreate, description string) {
	_ = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{CreateEmbedArticle(description)},
		},
	})
}
