package main

import (
	"context"
	"database/sql"
	"fmt"
	"os"

	"github.com/pressly/goose/v3"
	"github.com/pressly/goose/v3/lock"

	_ "github.com/jackc/pgx/v5/stdlib"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "example"
	dbname   = "postgres"
)

func main() {
	var db *sql.DB
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("pgx", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	sessionLocker, err := lock.NewPostgresSessionLocker(
		// Timeout after 30min. Try every 15s up to 120 times.
		lock.WithLockTimeout(1, 1),
	)
	if err != nil {
		panic(err)
	}

	provider, err := goose.NewProvider(
		goose.DialectPostgres,
		db,
		os.DirFS("migrations"),
		goose.WithSessionLocker(sessionLocker), // Use session-level advisory lock.
	)
	if err != nil {
		panic(err)
	}

	results, err := provider.Up(context.Background())
	if err != nil {
		panic(err)
	}

	if len(results) == 0 {
		fmt.Println("no migrations to apply")
		return
	}
	for _, result := range results {
		fmt.Println(result.Source.Path, "done in ", result.Duration.String())
	}
}
