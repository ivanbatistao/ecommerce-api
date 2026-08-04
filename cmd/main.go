package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/ivanbatistao/ecommerce-api/internal/env"
	"github.com/jackc/pgx/v5"
)

func main() {
	ctx := context.Background()

	config := Config{
		addr: ":8080",
		db: DbConfig{
			dsn: env.GetString("GOOSE_DBSTRING", "host=localhost user=postgres password=postgres dbname=ecommerce sslmode=disable"),
		},
	}

	// Logger
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	// Database connection
	conn, error := pgx.Connect(ctx, config.db.dsn)
	if error != nil {
		panic(error)
	}
	defer conn.Close(ctx)

	logger.Info("connected to database", "dns", config.db.dsn)

	// Application instance
	api := Application{
		config: config,
		db:     conn,
	}

	// Start API
	handler := api.mount()

	if error := api.run(handler); error != nil {
		slog.Error("Server has failed to start, error: %s", error)
		os.Exit(1)
	}
}
