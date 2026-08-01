package main

import (
	"log/slog"
	"os"
)

func main() {
	config := Config{
		addr: ":8080",
		db:   DbConfig{},
	}

	api := Application{
		config: config,
	}

	// Logger
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	// Start API
	handler := api.mount()

	if error := api.run(handler); error != nil {
		slog.Error("Server has failed to start, error: %s", error)
		os.Exit(1)
	}
}
