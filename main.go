package main

import (
	"context"
	"github.com/1rd0/DeliveryGo/config"
	"github.com/1rd0/DeliveryGo/pkg/gp"

	"github.com/jackc/pgx/v5/pgxpool"
)

type User struct {
	ID   int
	Name string
	Age  int
}

func main() {

	cfg, err := config.New()
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	db, err := gp.NewPoolConn(ctx, cfg)
	if err != nil {
		panic(err)
	}
	defer db.Close()

}

func InsertUser(ctx context.Context, db *pgxpool.Pool, user User) error {

	query := `INSERT INTO users (id, name, age) VALUES ($1, $2, $3)`

	_, err := db.Exec(ctx, query, user.ID, user.Name, user.Age)
	if err != nil {
		return err
	}

	return nil
}
