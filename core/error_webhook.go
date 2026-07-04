package core

import (
	"fmt"

	"github.com/RealMtrx/Discord-Handler-Go/config"
)

func SendErrorWebhook(errorMsg string) error {
	cfg := config.App
	if cfg.ErrorWebhook == "" || cfg.ErrorWebhook == "#" {
		return nil
	}

	embed := WebhookEmbed{
		Color:       0xFF0000,
		Title:       "❌ Bot Error Report",
		Description: fmt.Sprintf("**Error:** %s", errorMsg),
		Fields: []WebhookField{
			{Name: "📅 Timestamp", Value: makeTimestamp(), Inline: true},
		},
		Footer:    WebhookFooter{Text: footerText("Error Logger")},
		Timestamp: makeTimestamp(),
	}

	return SendWebhook(cfg.ErrorWebhook, embed)
}