package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/Shenith404/go-ecom/internal/env"
	"github.com/jackc/pgx/v5"
)

func main() {
	
	//Database
	ctx := context.Background()
	
	cfg := config{
		addr: ":8080",
		db: dbConfig{
			dsn:env.GestString("DB_DSN", "host=127.0.0.1 user=admin-user password=password dbname=go_ecommerce_db sslmode=disable"),
		},
	}
	api := application{
		config: cfg,
	}
	//loggger
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	slog.SetDefault(logger)

    //Database connection
	conn, err := pgx.Connect(ctx, cfg.db.dsn)
	if err != nil {
		panic(err)
	} 
	defer conn.Close(ctx)
	logger.Info("database connection established", "dsn", cfg.db.dsn)

	h := api.mount()
	if err := api.run(h); err != nil {
		slog.Error("server error", "error", err)
		os.Exit(1)
	}

}