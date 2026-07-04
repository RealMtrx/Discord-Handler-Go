package events

import (
	"fmt"

	"github.com/RealMtrx/Discord-Handler-Go/src/bot"
	"github.com/RealMtrx/Discord-Handler-Go/src/core"
	"github.com/bwmarrin/discordgo"
)

func InteractionCreate(s *discordgo.Session, i *discordgo.InteractionCreate) {
	if i.Type != discordgo.InteractionApplicationCommand {
		return
	}

	data := i.ApplicationCommandData()
	cmd, exists := bot.App.SlashCommands[data.Name]
	if !exists {
		return
	}

	var user *discordgo.User
	if i.Member != nil && i.Member.User != nil {
		user = i.Member.User
	} else if i.User != nil {
		user = i.User
	} else {
		return
	}

	onCooldown, remaining := core.Cooldowns.Check(user.ID, data.Name, 3000)
	if onCooldown {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "âڈ° Please wait " + fmt.Sprintf("%d", remaining) + " seconds before using this command again.",
				Flags:   1 << 6,
			},
		})
		return
	}

	guildName := "Direct Message"
	if i.GuildID != "" {
		guild, err := s.Guild(i.GuildID)
		if err == nil {
			guildName = guild.Name
		}
	}

	avatarURL := user.AvatarURL("256")
	core.SendSlashCommandUsage(user.ID, user.Username, data.Name, guildName, avatarURL)

	cmd.Handler(s, i)
}
