package core

import (
	"fmt"
	"time"

	"github.com/RealMtrx/Discord-Handler-Go/config"
)

func SendPrefixCommandUsage(userID, userName, commandName, guildName, avatarURL string) error {
	cfg := config.App
	if cfg.PrefixCommandWebhook == "" || cfg.PrefixCommandWebhook == "#" {
		return nil
	}

	embed := WebhookEmbed{
		Color:       0x57F287,
		Title:       fmt.Sprintf("%s Prefix Command Used", Emojis.Slash),
		Description: fmt.Sprintf("**Command:** `%s%s`", cfg.Prefix, commandName),
		Fields: []WebhookField{
			{Name: fmt.Sprintf("%s User Info", Emojis.User), Value: fmt.Sprintf("**UserName:** %s\n**ID:** %s", userName, userID), Inline: true},
			{Name: fmt.Sprintf("%s Server", Emojis.Server), Value: guildName, Inline: true},
			{Name: fmt.Sprintf("%s Time", Emojis.Loading), Value: fmt.Sprintf("<t:%d:R>", time.Now().Unix()), Inline: true},
		},
		Footer:    WebhookFooter{Text: footerText("Prefix Command Logger")},
		Timestamp: makeTimestamp(),
	}

	if avatarURL != "" {
		embed.Thumbnail = &WebhookThumbnail{URL: avatarURL}
	}

	return SendWebhook(cfg.PrefixCommandWebhook, embed)
}

func SendPrefixCommandError(userID, userName, commandName, guildName, errorMsg string) error {
	cfg := config.App
	if cfg.PrefixCommandWebhook == "" || cfg.PrefixCommandWebhook == "#" {
		return nil
	}

	embed := WebhookEmbed{
		Color:       0xFF0000,
		Title:       fmt.Sprintf("%s Prefix Command Error", Emojis.Error),
		Description: fmt.Sprintf("**Command:** `%s%s`\n**Error:** %s", cfg.Prefix, commandName, errorMsg),
		Fields: []WebhookField{
			{Name: fmt.Sprintf("%s User Info", Emojis.User), Value: fmt.Sprintf("%s (%s)", userName, userID), Inline: true},
			{Name: fmt.Sprintf("%s Server", Emojis.Server), Value: guildName, Inline: true},
			{Name: fmt.Sprintf("%s Time", Emojis.Loading), Value: fmt.Sprintf("<t:%d:F>", time.Now().Unix()), Inline: true},
		},
		Footer:    WebhookFooter{Text: footerText("Error Logger")},
		Timestamp: makeTimestamp(),
	}

	return SendWebhook(cfg.PrefixCommandWebhook, embed)
}