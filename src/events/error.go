package events

import (
	"fmt"

	"github.com/RealMtrx/Discord-Handler-Go/src/core"
	"github.com/bwmarrin/discordgo"
)

func DiscordError(s *discordgo.Session, err interface{}) {
	errorMsg := fmt.Sprintf("%v", err)
	core.SendErrorWebhook(errorMsg)
}
