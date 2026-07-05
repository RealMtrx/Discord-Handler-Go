<div align="center">
  <h1>Discord Handler — Go</h1>
  <p><strong>A production-ready Discord bot framework built with DiscordGo and MongoDB — slash commands, prefix commands, anti-crash, webhook logging, and a modular <code>src/</code> architecture.</strong></p>

  <p>
    <a href="https://github.com/RealMtrx/Discord-Handler-Go/blob/main/LICENSE"><img src="https://img.shields.io/badge/license-MIT-blue.svg" alt="License"></a>
    <a href="https://github.com/RealMtrx/Discord-Handler-Go/releases"><img src="https://img.shields.io/badge/version-0.9.0--beta-yellow" alt="Version 0.9.0 Beta"></a>
    <a href="https://github.com/RealMtrx/Discord-Handler-Go/stargazers"><img src="https://img.shields.io/github/stars/RealMtrx/Discord-Handler-Go" alt="Stars"></a>
    <a href="https://github.com/RealMtrx/Discord-Handler-Go/issues"><img src="https://img.shields.io/github/issues/RealMtrx/Discord-Handler-Go" alt="Issues"></a>
    <a href="https://github.com/RealMtrx/Discord-Handler-Go/network"><img src="https://img.shields.io/github/forks/RealMtrx/Discord-Handler-Go" alt="Forks"></a>
    <a href="https://github.com/RealMtrx/Discord-Handler/graphs/contributors"><img src="https://img.shields.io/badge/ecosystem-26%20repos-brightgreen" alt="26 Repos"></a>
    <a href="https://discord.gg/0hu2"><img src="https://img.shields.io/badge/discord-0hu2-5865F2" alt="Discord"></a>
  </p>

  <br>

  <p>
    <a href="#-features">Features</a> •
    <a href="#-quick-start">Quick Start</a> •
    <a href="#-project-structure">Structure</a> •
    <a href="#-api-reference">API</a> •
    <a href="#-database-edition">SQL Edition</a> •
    <a href="#-related-repositories">Ecosystem</a>
  </p>
</div>

---

## Overview

Discord Handler Go is a production-ready Discord bot framework built on **DiscordGo** with **MongoDB** storage. It provides a complete foundation for building Discord bots with slash commands, prefix commands, event handling, anti-crash protection, and webhook-based logging — all organized in a clean, scalable `src/` directory structure.

> **Version:** 0.9.0 (Stable Beta) — Part of the [Discord Handler](https://github.com/RealMtrx/Discord-Handler) ecosystem (26 repos across 13 languages).

## Features

- **Dual Command System** — Slash commands and prefix commands with auto-registration via DiscordGo
- **MongoDB Integration** — Persistent data storage with the official mongo-go-driver
- **Modular Architecture** — Clean separation: Commands, Events, Handlers, Core, Database, Models
- **Anti-Crash Protection** — Panic recovery with `recover()` and deferred handlers
- **Webhook Logging** — Dedicated webhooks for errors, slash commands, prefix commands, guild joins/leaves, and bot ready events
- **Cooldown System** — Per-command rate limiting using `CooldownManager` with goroutine-based cleanup
- **Emoji System** — Centralized emoji constants via `Emojis` struct for consistent rendering
- **Environment Configuration** — Secure token and secrets management via `os.Getenv` with fallbacks
- **Startup Report** — Color-coded terminal banner using `fatih/color` showing loaded commands, events, and connection status
- **Graceful Shutdown** — Proper cleanup on SIGINT, SIGTERM, and os.Interrupt signals

## Quick Start

```bash
# Clone the repository
git clone https://github.com/RealMtrx/Discord-Handler-Go.git
cd Discord-Handler-Go

# Download dependencies
go mod tidy

# Configure environment
cp .env.example .env
# Edit .env with your bot token, client ID, and MongoDB URI

# Run the bot
go run ./src/
```

### Prerequisites

- **Go 1.22+** — Runtime environment
- **MongoDB** — Local or Atlas instance
- **Discord Application** — Bot token and client ID from the [Discord Developer Portal](https://discord.com/developers/applications)

### Environment Variables

```env
TOKEN=your_bot_token
CLIENT_ID=your_client_id
BOT_NAME=Discord Handler
OWNER_IDS=owner_id_1,owner_id_2
PREFIX=$
MONGODB_URI=mongodb://localhost:27017/discord_bot
ERROR_WEBHOOK=your_webhook_url
SLASH_WEBHOOK=your_webhook_url
PREFIX_WEBHOOK=your_webhook_url
JOIN_WEBHOOK=your_webhook_url
LEAVE_WEBHOOK=your_webhook_url
READY_WEBHOOK=your_webhook_url
```

## Project Structure

```
Discord-Handler-Go/
├── go.mod                        # Go module definition
├── go.sum                        # Go module checksums
├── .env.example                  # Environment template
├── LICENSE
├── src/
│   ├── main.go                   # Entry point — initializes everything
│   ├── config/
│   │   └── config.go             # Config struct loaded from env vars
│   ├── bot/
│   │   └── bot.go                # Bot struct, Session, command maps, Start()
│   ├── Core/                     # Shared utilities
│   │   ├── command_utils.go      # FormatError, LogCommandUsage, ErrorReport
│   │   ├── cooldown.go           # CooldownManager with goroutine cleanup
│   │   ├── emojis.go             # Unicode emoji constants
│   │   ├── error_webhook.go      # Error reporting webhook
│   │   ├── join_guild_webhook.go # Guild join notification webhook
│   │   ├── leave_guild_webhook.go# Guild leave notification webhook
│   │   ├── prefix_command_webhook.go
│   │   ├── ready_webhook.go      # Bot ready event webhook
│   │   ├── slash_command_webhook.go
│   │   └── webhooks.go           # WebhookEmbed, SendWebhook, shared helpers
│   ├── Database/
│   │   └── mongo.go              # MongoDB connection, ping, disconnect
│   ├── Events/                   # Discord event handlers
│   │   ├── error.go              # DiscordError — sends error webhook
│   │   ├── guild_create.go       # Sends guild join webhook
│   │   ├── guild_delete.go       # Sends guild leave webhook
│   │   ├── interaction_create.go # Slash command dispatch + cooldown
│   │   ├── message_create.go     # Prefix command dispatch + cooldown
│   │   └── ready.go              # Bot ready — sets activity, sends webhook
│   ├── Handlers/                 # Loaders and registrars
│   │   ├── anticrash.go          # SetupAntiCrash, RecoverPanic
│   │   ├── commands.go           # LoadSlashCommands, RegisterSlashCommands
│   │   ├── events.go             # RegisterEvents — attaches all listeners
│   │   ├── logger.go             # StartupReport with color-coded banner
│   │   ├── models.go             # LoadModels — model reporting
│   │   └── prefix.go             # LoadPrefixCommands — in-memory registration
│   ├── Models/
│   │   └── user.go               # User struct with GetUser, CreateUser
│   └── Commands/
│       ├── Slash/
│       │   └── public/
│       │       └── ping.go       # Shows bot, WebSocket, and REST latency
│       └── Prefix/
│           └── public/
│               └── ping.go       # Shows bot, WebSocket, and REST latency
```

## API Reference

### Core Types

| Type | Location | Description |
| ---- | -------- | ----------- |
| `Config` | `config/config.go` | Application configuration loaded from environment variables |
| `Bot` | `bot/bot.go` | Main bot struct holding Session, Config, and command maps |
| `SlashCommand` | `bot/bot.go` | Slash command definition with `ApplicationCommand` data and handler |
| `PrefixCommand` | `bot/bot.go` | Prefix command definition with name and handler |
| `CooldownManager` | `core/cooldown.go` | Thread-safe per-command cooldown tracker |
| `WebhookEmbed` | `core/webhooks.go` | Discord embed struct for webhook payloads |
| `User` | `models/user.go` | MongoDB user model with `userId` and `points` fields |

### Core Functions

| Function | Location | Description |
| -------- | -------- | ----------- |
| `config.Load()` | `config/config.go` | Loads env vars into global `Config` singleton |
| `bot.New(cfg)` | `bot/bot.go` | Creates and returns a new Bot instance |
| `b.Start()` | `bot/bot.go` | Opens Discord Gateway session with configured intents |
| `Cooldowns.Check(userID, command, ms)` | `core/cooldown.go` | Returns `(onCooldown, remainingSeconds)` |
| `SendWebhook(url, embed)` | `core/webhooks.go` | Posts an embed to a Discord webhook URL |
| `FormatError(err, commandName)` | `core/command_utils.go` | Returns an `ErrorReport` struct |
| `LogCommandUsage(userID, userName, cmd, guild)` | `core/command_utils.go` | Prints command usage to stdout |

### Event Handlers

| Function | Location | Description |
| -------- | -------- | ----------- |
| `Ready(s, r)` | `events/ready.go` | Sets bot activity, sends ready webhook |
| `InteractionCreate(s, i)` | `events/interaction_create.go` | Routes slash commands with cooldown + webhook logging |
| `MessageCreate(s, m)` | `events/message_create.go` | Routes prefix commands with cooldown + webhook logging |
| `GuildCreate(s, g)` | `events/guild_create.go` | Sends guild join notification |
| `GuildDelete(s, g)` | `events/guild_delete.go` | Sends guild leave notification |
| `DiscordError(s, err)` | `events/error.go` | Sends error to webhook |

### Handlers

| Function | Location | Description |
| -------- | -------- | ----------- |
| `SetupAntiCrash()` | `handlers/anticrash.go` | Configures log output and panic recovery |
| `RecoverPanic()` | `handlers/anticrash.go` | Deferred panic recovery with stack trace |
| `RegisterEvents(b)` | `handlers/events.go` | Attaches all 5 event handlers to the session |
| `LoadSlashCommands(b)` | `handlers/commands.go` | Reports loaded slash commands count |
| `RegisterSlashCommands(b)` | `handlers/commands.go` | Bulk-overwrites global slash commands with Discord API |
| `LoadPrefixCommands(b)` | `handlers/prefix.go` | Reports loaded prefix commands count |
| `LoadModels()` | `handlers/models.go` | Reports model count |
| `StartupReport(data)` | `handlers/logger.go` | Prints color-coded startup summary |

## Adding Commands

### Slash Command

Create a new file in `src/Commands/Slash/[category]/[name].go`:

```go
package public

import (
    "github.com/RealMtrx/Discord-Handler-Go/src/bot"
    "github.com/bwmarrin/discordgo"
)

func RegisterSlashHello(b *bot.Bot) {
    b.SlashCommands["hello"] = &bot.SlashCommand{
        Data: &discordgo.ApplicationCommand{
            Name:        "hello",
            Description: "Say hello!",
        },
        Handler: func(s *discordgo.Session, i *discordgo.InteractionCreate) {
            s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
                Type: discordgo.InteractionResponseChannelMessageWithSource,
                Data: &discordgo.InteractionResponseData{
                    Content: "Hello! 👋",
                },
            })
        },
    }
}
```

Then register it in `main.go`:

```go
import helloPublic "github.com/RealMtrx/Discord-Handler-Go/src/commands/slash/public"

func main() {
    // ... after bot.New(cfg)
    helloPublic.RegisterSlashHello(b)
    // ...
}
```

### Prefix Command

Create a new file in `src/Commands/Prefix/[category]/[name].go`:

```go
package public

import (
    "github.com/RealMtrx/Discord-Handler-Go/src/bot"
    "github.com/bwmarrin/discordgo"
)

func RegisterPrefixHello(b *bot.Bot) {
    b.PrefixCommands["hello"] = &bot.PrefixCommand{
        Name: "hello",
        Handler: func(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {
            s.ChannelMessageSend(m.ChannelID, "Hello! 👋")
        },
    }
}
```

## Database Edition

This is the **MongoDB edition**. A **SQL edition** using Sequelize ORM is also available:

| Feature | MongoDB Edition | SQL Edition |
| ------- | --------------- | ----------- |
| Repository | [Discord-Handler-Go](https://github.com/RealMtrx/Discord-Handler-Go) | [Discord-Handler-Go-Sequelize](https://github.com/RealMtrx/Discord-Handler-Go-Sequelize) |
| Database | MongoDB | SQLite, PostgreSQL, MySQL, MSSQL |
| Driver | mongo-go-driver | Sequelize / GORM |
| Dialects | MongoDB only | Multi-dialect via config |

## Related Repositories

Discord Handler Go is part of a **26-repo ecosystem**. Here are the other repositories:

### Core Framework (MongoDB)

| Language | Repository |
| -------- | ---------- |
| JavaScript | [Discord-Handler-Js](https://github.com/RealMtrx/Discord-Handler-Js) |
| TypeScript | [Discord-Handler-Ts](https://github.com/RealMtrx/Discord-Handler-Ts) |
| Rust | [Discord-Handler-Rs](https://github.com/RealMtrx/Discord-Handler-Rs) |
| Python | [Discord-Handler-Py](https://github.com/RealMtrx/Discord-Handler-Py) |
| C# | [Discord-Handler-Cs](https://github.com/RealMtrx/Discord-Handler-Cs) |
| Java | [Discord-Handler-Java](https://github.com/RealMtrx/Discord-Handler-Java) |
| Kotlin | [Discord-Handler-Kt](https://github.com/RealMtrx/Discord-Handler-Kt) |
| C++ | [Discord-Handler-Cpp](https://github.com/RealMtrx/Discord-Handler-Cpp) |
| Dart | [Discord-Handler-Dart](https://github.com/RealMtrx/Discord-Handler-Dart) |
| Ruby | [Discord-Handler-Rb](https://github.com/RealMtrx/Discord-Handler-Rb) |
| Lua | [Discord-Handler-Lua](https://github.com/RealMtrx/Discord-Handler-Lua) |
| PHP | [Discord-Handler-Php](https://github.com/RealMtrx/Discord-Handler-Php) |

### Database Editions (SQL)

| Language | Repository |
| -------- | ---------- |
| JavaScript | [Discord-Handler-Js-Sequelize](https://github.com/RealMtrx/Discord-Handler-Js-Sequelize) |
| TypeScript | [Discord-Handler-Ts-Sequelize](https://github.com/RealMtrx/Discord-Handler-Ts-Sequelize) |
| Go | [Discord-Handler-Go-Sequelize](https://github.com/RealMtrx/Discord-Handler-Go-Sequelize) |
| Rust | [Discord-Handler-Rs-Sequelize](https://github.com/RealMtrx/Discord-Handler-Rs-Sequelize) |
| Python | [Discord-Handler-Py-Sequelize](https://github.com/RealMtrx/Discord-Handler-Py-Sequelize) |
| C# | [Discord-Handler-Cs-Sequelize](https://github.com/RealMtrx/Discord-Handler-Cs-Sequelize) |
| Java | [Discord-Handler-Java-Sequelize](https://github.com/RealMtrx/Discord-Handler-Java-Sequelize) |
| Kotlin | [Discord-Handler-Kt-Sequelize](https://github.com/RealMtrx/Discord-Handler-Kt-Sequelize) |
| C++ | [Discord-Handler-Cpp-Sequelize](https://github.com/RealMtrx/Discord-Handler-Cpp-Sequelize) |
| Dart | [Discord-Handler-Dart-Sequelize](https://github.com/RealMtrx/Discord-Handler-Dart-Sequelize) |
| Ruby | [Discord-Handler-Rb-Sequelize](https://github.com/RealMtrx/Discord-Handler-Rb-Sequelize) |
| Lua | [Discord-Handler-Lua-Sequelize](https://github.com/RealMtrx/Discord-Handler-Lua-Sequelize) |
| PHP | [Discord-Handler-Php-Sequelize](https://github.com/RealMtrx/Discord-Handler-Php-Sequelize) |

### Hub

| Repository | Description |
| ---------- | ----------- |
| [Discord-Handler](https://github.com/RealMtrx/Discord-Handler) | Central hub — documentation, examples, changelog, roadmap |

## License

MIT License — Copyright © 2026 Mtrx

---

<div align="center">
  <sub>Built by <strong>Mtrx</strong> — Discord: <strong>0hu2</strong></sub>
  <br>
  <sub><a href="https://github.com/RealMtrx/Discord-Handler">Discord Handler Ecosystem</a></sub>
</div>
