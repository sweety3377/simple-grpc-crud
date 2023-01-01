package database

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"simple-crud/config"
)

func Connect(ctx context.Context, cfg *config.Config) (*pgx.Conn, error) {
	connString := getConnString(cfg)

	conn, err := pgx.Connect(ctx, connString)
	if err != nil {
		return nil, err
	}

	// Database healthcheck
	err = conn.Ping(ctx)
	if err != nil {
		return nil, err
	}

	return conn, nil
}

// Format connection string from config
func getConnString(cfg *config.Config) string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", cfg.DbHost, cfg.DbPort, cfg.DbUser, cfg.DbPassword, cfg.DbName)
}
