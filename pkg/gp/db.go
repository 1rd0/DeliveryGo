package gp

import (
	"context"
	"fmt"
	"github.com/1rd0/DeliveryGo/config"
	_ "github.com/golang-migrate/migrate/database/postgres"
	_ "github.com/golang-migrate/migrate/source/file"
	"github.com/jackc/pgx/v5/pgxpool"
)

func NewPoolConn(ctx context.Context, config *config.Config) (*pgxpool.Pool, error) {
	connStr := config.DatabaseURL()

	fmt.Println("Connecting to database with:", connStr) // Debug log

	conn, err := pgxpool.New(ctx, connStr)
	if err != nil {
		fmt.Println("Unable to connect to database:", err)
		return nil, fmt.Errorf("could not connect to gp: %w", err)
	}
	if err := conn.Ping(ctx); err != nil {
		conn.Close()
		fmt.Printf("Unable to ping database: %v\n", err)
		return nil, fmt.Errorf("could not ping gp: %w", err)
	}
	fmt.Println("Connected to database")
	return conn, nil
}
