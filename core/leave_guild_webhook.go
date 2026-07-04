package core

import (
	"fmt"
	"time"

	"github.com/RealMtrx/Discord-Handler-Go/config"
)

func SendLeaveGuildWebhook(guildID, guildName string, memberCount int, remainingServers int) error {
	cfg := config.App
	if cfg.LeaveGuildWebhook == "" || cfg.LeaveGuildWebhook == "#" {
		return nil
	}

	embed := WebhookEmbed{
		Color:       0xFF0000,
		Title:       "👋 Bot Left Server",
		Description: fmt.Sprintf("**Server:** %s\n**ID:** %s", guildName, guildID),
		Fields: []WebhookField{
			{Name: "👥 Members", Value: fmt.Sprintf("%d members", memberCount), Inline: true},
			{Name: "📅 Left At", Value: fmt.Sprintf("<t:%d:F>", time.Now().Unix()), Inline: true},
			{Name: "📊 Remaining Servers", Value: fmt.Sprintf("%d servers", remainingServers), Inline: true},
		},
		Footer:    WebhookFooter{Text: footerText("Guild Leave Logger")},
		Timestamp: makeTimestamp(),
	}

	return SendWebhook(cfg.LeaveGuildWebhook, embed)
}