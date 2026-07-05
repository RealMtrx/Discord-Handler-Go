# Discord Handler Go

A modern, feature-rich Discord bot handler built with **discordgo**, featuring both slash commands and prefix commands with a robust modular architecture designed for scalability and maintainability.

## 🚀 Features

- **Dual Command System**: Support for both slash commands and prefix commands
- **Modular Architecture**: Clean separation of concerns with dedicated handlers
- **Anti-Crash System**: Comprehensive error handling and monitoring
- **Event-Driven**: Fully event-driven architecture
- **Webhook Logging**: Real-time logging for errors and guild events
- **MongoDB Integration**: Persistent data storage with mongo-go-driver
- **Cooldown System**: Per-command cooldown management
- **Environment Configuration**: Secure configuration management with godotenv

## 📁 Project Structure

```
Discord-Handler-Go/
├── go.mod                       # Go module definition
├── go.sum                       # Go module checksums
├── src/                         # Source code
│   ├── main.go                  # Main bot entry point
│   ├── config/config.go         # Bot configuration from .env
│   ├── bot/bot.go               # Bot initialization
│   ├── Core/                    # Core utilities
│   │   ├── commandUtils.go      # Cooldown and utilities
│   │   ├── emojis.go            # Centralized emoji definitions
│   │   └── webhookUtil.go       # Webhook utility
│   ├── Database/
│   │   └── mongo.go             # MongoDB connection setup
│   ├── Events/                  # Discord event handlers
│   │   ├── guildCreate.go       # Handler when bot joins a server
│   │   ├── guildDelete.go       # Handler when bot leaves a server
│   │   ├── interactionCreate.go # Handles slash command interactions
│   │   ├── messageCreate.go     # Handles prefix commands
│   │   └── ready.go             # Bot ready event
│   ├── Handlers/                # Handlers for modularity
│   │   ├── AntiCrash.go         # Crash prevention and error handling
│   │   └── logger.go            # Logger for bot activity
│   ├── Models/
│   │   └── userModel.go         # User data model
│   └── Commands/
│       ├── Prefix/              # Prefix commands
│       │   └── ping.go          # Example prefix ping command
│       └── Slash/               # Slash commands
│           └── ping.go          # Example slash ping command
```

## 🔧 Installation

1. **Clone the repository**

   ```bash
   git clone https://github.com/RealMtrx/Discord-Handler-Go.git
   cd Discord-Handler-Go
   ```

2. **Install dependencies**

   ```bash
   go mod tidy
   ```

3. **Environment Setup**

   Copy `.env.example` to `.env` and fill in your values:

   ```env
   TOKEN=your_bot_token_here
   PREFIX=!
   BOT_NAME=Discord Handler
   MONGO_URI=mongodb://localhost:27017/discord-handler
   ERROR_WEBHOOK=https://discord.com/api/webhooks/your_webhook
   GUILD_LOG_WEBHOOK=https://discord.com/api/webhooks/your_webhook
   ```

4. **Run the bot**

   ```bash
   go run ./src/
   # or build and run
   go build ./src/
   ./Discord-Handler-Go
   ```

## 📋 Dependencies

- **discordgo**: v0.3 - Discord API wrapper
- **godotenv**: v1.5 - Environment variable management
- **mongo-go-driver**: v1 - MongoDB driver

## 📝 Command Development

### Creating Slash Commands

Create a new file in `src/Commands/Slash/[category]/[name].go`:

```go
package slash

import (
    "github.com/bwmarrin/discordgo"
)

var Ping = &Command{
    Name:        "ping",
    Description: "Replies with Pong!",
    Execute: func(s *discordgo.Session, i *discordgo.InteractionCreate) {
        s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
            Type: discordgo.InteractionResponseChannelMessageWithSource,
            Data: &discordgo.InteractionResponseData{
                Content: "Pong! 🏓",
            },
        })
    },
}
```

### Creating Prefix Commands

Create a new file in `src/Commands/Prefix/[category]/[name].go`:

```go
package prefix

import (
    "github.com/bwmarrin/discordgo"
)

var Ping = &Command{
    Name:    "ping",
    Execute: func(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {
        s.ChannelMessageSend(m.ChannelID, "Pong! 🏓")
    },
}
```

---

**Discord Handler** - A modern, scalable Discord bot framework built with Go.
