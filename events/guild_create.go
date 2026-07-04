package events

import (
	"github.com/RealMtrx/Discord-Handler-Go/core"
	"github.com/bwmarrin/discordgo"
)

func GuildCreate(s *discordgo.Session, g *discordgo.GuildCreate) {
	iconURL := g.IconURL("256")
	core.SendJoinGuildWebhook(g.Name, g.ID, g.OwnerID, g.MemberCount, iconURL)
}
