package core

import (
	"fmt"

	"github.com/RealMtrx/Discord-Handler-Go/config"
)

func SendReadyWebhook(botUsername, botID string, serverCount int) error {
	cfg := config.App
	if cfg.ReadyWebhook == "" || cfg.ReadyWebhook == "#" {
		return nil
	}

	embed := WebhookEmbed{
		Color:       0x00FF00,
		Title:       "🟢 Bot is Online!",
		Description: fmt.Sprintf("**Bot:** %s\n**Status:** Online and Ready", botUsername),
		Fields: []WebhookField{
			{Name: "🤖 Bot Info", Value: fmt.Sprintf("**ID:** %s", botID), Inline: true},
			{Name: "🏠 Servers", Value: fmt.Sprintf("%d servers", serverCount), Inline: true},
		},
		Footer:    WebhookFooter{Text: footerText("System Logger")},
		Timestamp: makeTimestamp(),
	}

	return SendWebhook(cfg.ReadyWebhook, embed)
}