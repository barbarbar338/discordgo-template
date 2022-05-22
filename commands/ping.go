package commands

import "github.com/bwmarrin/discordgo"

var PingData = discordgo.ApplicationCommand {
	Name: "ping",
	Description: "Replies with pong!",
}

func PingHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData {
			Content: ":ping_pong: Pong!",
		},
	})
}
