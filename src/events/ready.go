package events

import (
	"github.com/RealMtrx/Discord-Handler-Go/src/bot"
	"github.com/RealMtrx/Discord-Handler-Go/src/core"
	"github.com/bwmarrin/discordgo"
)

func Ready(s *discordgo.Session, r *discordgo.Ready) {
	s.UpdateStatusComplex(discordgo.UpdateStatusData{
		Activities: []*discordgo.Activity{
			{Name: bot.App.Config.BotName, Type: discordgo.ActivityTypeGame},
		},
		Status: "online",
	})

	botUsername := s.State.User.Username
	botID := s.State.User.ID
	serverCount := len(s.State.Guilds)
	core.SendReadyWebhook(botUsername, botID, serverCount)
}
