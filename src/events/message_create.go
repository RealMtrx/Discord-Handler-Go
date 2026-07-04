package events

import (
	"fmt"
	"strings"

	"github.com/RealMtrx/Discord-Handler-Go/src/bot"
	"github.com/RealMtrx/Discord-Handler-Go/src/core"
	"github.com/bwmarrin/discordgo"
)

func MessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.Bot || m.GuildID == "" {
		return
	}

	prefix := bot.App.Config.Prefix
	if !strings.HasPrefix(m.Content, prefix) {
		return
	}

	args := strings.Fields(strings.TrimPrefix(m.Content, prefix))
	if len(args) == 0 {
		return
	}

	cmdName := strings.ToLower(args[0])
	cmd, exists := bot.App.PrefixCommands[cmdName]
	if !exists {
		return
	}

	onCooldown, remaining := core.Cooldowns.Check(m.Author.ID, cmdName, 3000)
	if onCooldown {
		s.ChannelMessageSend(m.ChannelID,
			"âڈ° Please wait "+fmt.Sprintf("%d", remaining)+" seconds before using this command again.")
		return
	}

	guildName := "Direct Message"
	guild, err := s.Guild(m.GuildID)
	if err == nil {
		guildName = guild.Name
	}

	avatarURL := m.Author.AvatarURL("256")
	core.SendPrefixCommandUsage(m.Author.ID, m.Author.Username, cmdName, guildName, avatarURL)

	go cmd.Handler(s, m, args[1:])
}
