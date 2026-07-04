package handlers

import (
	"fmt"
	"time"

	"github.com/RealMtrx/Discord-Handler-Go/src/config"
	"github.com/fatih/color"
)

type StartupData struct {
	Name         string
	SlashCount   int
	PrefixCount  int
	EventsCount  int
	ModelsCount  int
	MongoStatus  bool
}

func StartupReport(data StartupData) {
	fmt.Println()
	color.Cyan("â•”â•گâ•گâ•گâ•گâ•گâ•گâ•گâ•گâ•گâ•گâ•گâ•گâ•گâ•گâ•گâ•گâ•گâ•گâ•گâ•گâ•گâ•گâ•گâ•گâ•گâ•گâ•گâ•گâ•گâ•گâ•گâ•گâ•گâ•گâ•—")
	color.Cyan(fmt.Sprintf("â•‘     %-30sâ•‘", config.App.BotName))
	color.Cyan("â•ڑâ•گâ•گâ•گâ•گâ•گâ•گâ•گâ•گâ•گâ•گâ•گâ•گâ•گâ•گâ•گâ•گâ•گâ•گâ•گâ•گâ•گâ•گâ•گâ•گâ•گâ•گâ•گâ•گâ•گâ•گâ•گâ•گâ•گâ•گâ•‌")
	fmt.Println()

	lines := []struct {
		label string
		ok    bool
	}{
		{fmt.Sprintf("Slash Commands: %d", data.SlashCount), true},
		{fmt.Sprintf("Prefix Commands: %d", data.PrefixCount), true},
		{fmt.Sprintf("Events Loaded: %d", data.EventsCount), true},
		{fmt.Sprintf("Models Loaded: %d", data.ModelsCount), true},
		{"AntiCrash: Active", true},
		{fmt.Sprintf("MongoDB: Connected = %v", data.MongoStatus), data.MongoStatus},
	}

	for _, l := range lines {
		if l.ok {
			color.Green("  âœ… %s", l.label)
		} else {
			color.Red("  â‌Œ %s", l.label)
		}
	}

	fmt.Println()
	color.Magenta("[ %s ] %s is now online and fully operational.", time.Now().Format("02/01/2006 15:04:05"), data.Name)
}
