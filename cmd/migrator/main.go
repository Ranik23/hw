package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {

	command := flag.String("command", "up", "Command to execute: up, down, or force")
	version := flag.Int("version", 0, "Version to force")
	flag.Parse()

	m, err := migrate.New(
		"file://internal/storage/migrations",
		"postgres://postgres:postgres@localhost:5432/library?sslmode=disable&search_path=library",
	)
	if err != nil {
		log.Fatalf("Unable to create migration instance: %v", err)
	}
	defer m.Close()


	switch *command {
	case "up":
		fmt.Println("Applying all up migrations...")
		if err := m.Up(); err != nil && err != migrate.ErrNoChange {
			log.Fatalf("Migration up failed: %v", err)
		}
		fmt.Println("Migrations applied successfully.")
	case "down":
		fmt.Println("Rolling back all down migrations...")
		if err := m.Down(); err != nil && err != migrate.ErrNoChange {
			log.Fatalf("Migration down failed: %v", err)
		}
		fmt.Println("Migrations rolled back successfully.")
	case "force":
		if *version == 0 {
			log.Fatal("Version is required for the 'force' command.")
		}
		fmt.Printf("Forcing migration to version %d...\n", *version)
		if err := m.Force(*version); err != nil {
			log.Fatalf("Force migration failed: %v", err)
		}
		fmt.Println("Migration forced successfully.")
	default:
		log.Fatalf("Unknown command: %s. Use 'up', 'down', or 'force'.", *command)
	}
}