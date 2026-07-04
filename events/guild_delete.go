package events

import (
	"github.com/RealMtrx/Discord-Handler-Go/core"
	"github.com/bwmarrin/discordgo"
)

func GuildDelete(s *discordgo.Session, g *discordgo.GuildDelete) {
	remainingServers := len(s.State.Guilds)
	core.SendLeaveGuildWebhook(g.ID, g.Name, 0, remainingServers)
}
