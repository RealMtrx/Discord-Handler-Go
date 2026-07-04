package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/RealMtrx/Discord-Handler-Go/src/bot"
	slashPublic "github.com/RealMtrx/Discord-Handler-Go/src/commands/slash/public"
	prefixPublic "github.com/RealMtrx/Discord-Handler-Go/src/commands/prefix/public"
	"github.com/RealMtrx/Discord-Handler-Go/src/config"
	"github.com/RealMtrx/Discord-Handler-Go/src/database"
	"github.com/RealMtrx/Discord-Handler-Go/src/events"
	"github.com/RealMtrx/Discord-Handler-Go/src/handlers"
	"github.com/fatih/color"
)

func main() {
	handlers.SetupAntiCrash()
	defer handlers.RecoverPanic()

	cfg := config.Load()
	config.App = cfg

	color.Cyan("â•”â•گâ•گâ•گâ•گâ•گâ•گâ•گâ•گâ•گâ•گâ•گâ•گâ•گâ•گâ•گâ•گâ•گâ•گâ•گâ•گâ•گâ•گâ•گâ•گâ•گâ•گâ•گâ•گâ•گâ•گâ•گâ•گâ•گâ•گâ•—")
	color.Cyan(fmt.Sprintf("â•‘     Starting %-20sâ•‘", cfg.BotName))
	color.Cyan("â•ڑâ•گâ•گâ•گâ•گâ•گâ•گâ•گâ•گâ•گâ•گâ•گâ•گâ•گâ•گâ•گâ•گâ•گâ•گâ•گâ•گâ•گâ•گâ•گâ•گâ•گâ•گâ•گâ•گâ•گâ•گâ•گâ•گâ•گâ•گâ•‌")
	println()

	b := bot.New(cfg)

	slashPublic.RegisterSlashPing(b)
	prefixPublic.RegisterPrefixPing(b)

	color.Blue("[System] Initializing AntiCrash...")
	color.Green("[System] AntiCrash active")

	color.Blue("[System] Connecting to MongoDB...")
	mongoConnected := database.Connect()
	if mongoConnected {
		defer database.Disconnect()
	}

	color.Blue("[System] Loading slash commands...")
	handlers.LoadSlashCommands(b)

	color.Blue("[System] Loading prefix commands...")
	handlers.LoadPrefixCommands(b)

	color.Blue("[System] Registering events...")
	eventsCount := handlers.RegisterEvents(b)
	b.Session.AddHandler(events.DiscordError)

	color.Blue("[System] Loading models...")
	modelsCount := handlers.LoadModels()

	color.Blue("[System] Starting bot...")
	if err := b.Start(); err != nil {
		color.Red("[System] Failed to start: %v", err)
		os.Exit(1)
	}

	color.Blue("[System] Registering slash commands with Discord API...")
	if err := handlers.RegisterSlashCommands(b); err != nil {
		color.Red("[System] %v", err)
	}

	handlers.StartupReport(handlers.StartupData{
		Name:         cfg.BotName,
		SlashCount:   len(b.SlashCommands),
		PrefixCount:  len(b.PrefixCommands),
		EventsCount:  eventsCount,
		ModelsCount:  modelsCount,
		MongoStatus:  mongoConnected,
	})

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	println()
	color.Yellow("[System] Shutting down...")
	b.Session.Close()
}
