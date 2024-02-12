package main

import (
	"flag"
	"log/slog"
	"os"
	"utf-moodle-scraper/internal/backend"
)

func main() {
	debugFlag := flag.Bool("debug", false, "Enable debug mode")
	flag.Parse()
	if *debugFlag {
		logger := slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{
			Level: slog.LevelDebug,
		}))
		slog.SetDefault(logger)
		slog.Debug("Debug mode enabled.")
	}
	slog.Info("Starting service...")

	backend.Run(*debugFlag)
}
