package functions

import (
	"github.com/bwmarrin/discordgo"
	"time"
)

func EmbedCreate(title string, description string, thumbnail string) *discordgo.MessageEmbed {
		embed := &discordgo.MessageEmbed{
			Fields: []*discordgo.MessageEmbedField{&discordgo.MessageEmbedField{
				Name:   "Go helper",
				Value:  description,
				Inline: true,
			},
			},
			Thumbnail: &discordgo.MessageEmbedThumbnail{
				URL: thumbnail,
			},
			Footer: &discordgo.MessageEmbedFooter{
				Text: "Made by Leki#6796",
			},
			Timestamp: time.Now().Format(time.RFC3339),
			Title:     title,
		}
		return embed
	}

