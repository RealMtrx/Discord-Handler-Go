package core

import (
	"fmt"
	"time"

	"github.com/RealMtrx/Discord-Handler-Go/config"
)

func SendJoinGuildWebhook(guildName, guildID, ownerID string, memberCount int, iconURL string) error {
	cfg := config.App
	if cfg.JoinGuildWebhook == "" || cfg.JoinGuildWebhook == "#" {
		return nil
	}

	embed := WebhookEmbed{
		Color:       0x57F287,
		Title:       "🎉 Bot Joined New Server!",
		Description: fmt.Sprintf("**Server:** %s\n**ID:** %s", guildName, guildID),
		Fields: []WebhookField{
			{Name: "👑 Owner", Value: "<@" + ownerID + ">", Inline: true},
			{Name: "👥 Members", Value: fmt.Sprintf("%d members", memberCount), Inline: true},
			{Name: "📅 Joined At", Value: fmt.Sprintf("<t:%d:F>", time.Now().Unix()), Inline: true},
		},
		Footer:    WebhookFooter{Text: footerText("Guild Join Logger")},
		Timestamp: makeTimestamp(),
	}

	if iconURL != "" {
		embed.Thumbnail = &WebhookThumbnail{URL: iconURL}
	}

	return SendWebhook(cfg.JoinGuildWebhook, embed)
}