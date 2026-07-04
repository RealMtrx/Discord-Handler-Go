package core

import (
	"fmt"
	"time"

	"github.com/RealMtrx/Discord-Handler-Go/config"
)

func SendSlashCommandUsage(userID, userName, commandName, guildName, avatarURL string) error {
	cfg := config.App
	if cfg.SlashCommandWebhook == "" || cfg.SlashCommandWebhook == "#" {
		return nil
	}

	embed := WebhookEmbed{
		Color:       0x5865F2,
		Title:       fmt.Sprintf("%s Slash Command Used", Emojis.Slash),
		Description: fmt.Sprintf("**Command:** `/%s`", commandName),
		Fields: []WebhookField{
			{Name: fmt.Sprintf("%s User Info", Emojis.User), Value: fmt.Sprintf("**UserName:** %s\n**ID:** %s", userName, userID), Inline: true},
			{Name: fmt.Sprintf("%s Server", Emojis.Server), Value: guildName, Inline: true},
			{Name: fmt.Sprintf("%s Time", Emojis.Loading), Value: fmt.Sprintf("<t:%d:R>", time.Now().Unix()), Inline: true},
		},
		Footer:    WebhookFooter{Text: footerText("Slash Command Logger")},
		Timestamp: makeTimestamp(),
	}

	if avatarURL != "" {
		embed.Thumbnail = &WebhookThumbnail{URL: avatarURL}
	}

	return SendWebhook(cfg.SlashCommandWebhook, embed)
}

func SendSlashCommandError(userID, userName, commandName, guildName, errorMsg string) error {
	cfg := config.App
	if cfg.SlashCommandWebhook == "" || cfg.SlashCommandWebhook == "#" {
		return nil
	}

	embed := WebhookEmbed{
		Color:       0xFF0000,
		Title:       fmt.Sprintf("%s Slash Command Error", Emojis.Error),
		Description: fmt.Sprintf("**Command:** `/%s`\n**Error:** %s", commandName, errorMsg),
		Fields: []WebhookField{
			{Name: fmt.Sprintf("%s User Info", Emojis.User), Value: fmt.Sprintf("%s (%s)", userName, userID), Inline: true},
			{Name: fmt.Sprintf("%s Server", Emojis.Server), Value: guildName, Inline: true},
			{Name: fmt.Sprintf("%s Time", Emojis.Loading), Value: fmt.Sprintf("<t:%d:F>", time.Now().Unix()), Inline: true},
		},
		Footer:    WebhookFooter{Text: footerText("Error Logger")},
		Timestamp: makeTimestamp(),
	}

	return SendWebhook(cfg.SlashCommandWebhook, embed)
}