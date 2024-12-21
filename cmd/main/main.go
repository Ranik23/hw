package main

import (
	"database/sql"
	"hw/internal/seed"
	"log"

)


func connectDB() *sql.DB {
	dsn := "host=localhost user=postgres password=postgres dbname=library port=5432 sslmode=disable search_path=library"
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	return db
}

func main() {
	db := connectDB()
	defer db.Close()
	//seed.GenerateAndInsertAuthors(db)
	//seed.GenerateAndInsertBooks(db)
	//seed.GenerateAndInsertGenres(db)
	//seed.GenerateAndInsertSubscribers(db)
	//seed.GenerateAndInsertM2MBooksAuthors(db)
	seed.GenerateAndInsertM2MBooksGenres(db)
	// seed.GenerateAndInsertSubscriptions(db)

}