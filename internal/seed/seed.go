package seed

import (
	"database/sql"
	"fmt"
	"hw/internal/api"
	"log"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/brianvoe/gofakeit"
	"github.com/bxcodec/faker/v3"
	_ "github.com/lib/pq"
)

func GenerateAndInsertAuthors(db *sql.DB) {

	count := 10000

	response := api.GetData(count)

	actualCount := min(count, response.Total)

	query := "INSERT INTO library.Authors (a_id, a_name) VALUES ($1, $2)"

	current := "SELECT COUNT(*) FROM library.Authors"
	var c int
	err := db.QueryRow(current).Scan(&c)
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < actualCount; i++ {
		_, err := db.Exec(query, i+c, response.Data[i].Author)
		if err != nil {
			log.Fatalf("Failed to insert data into Authors: %v", err)
		}
	}

	fmt.Printf("Inserted %d authors\n", actualCount)
}

func GenerateAndInsertBooks(db *sql.DB) {

	count := 100000

	response := api.GetData(count)

	actualCount := min(count, response.Total)

	query := "INSERT INTO library.books (b_id, b_name, b_year, b_quantity) VALUES ($1, $2, $3, $4)"

	current := "SELECT COUNT(*) FROM library.books"
	var c int
	err := db.QueryRow(current).Scan(&c)
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < actualCount; i++ {
		date := strings.Split(response.Data[i].Published, "-")
		year, _ := strconv.Atoi(date[0])

		_, err := db.Exec(query, i+c+1, response.Data[i].Title, year, gofakeit.Number(1, 1000))
		if err != nil {
			log.Fatalf("Failed to insert data into books: %v", err)
		}
	}

	fmt.Printf("Inserted %d books\n", count)
}

// Генерация и вставка данных для таблицы genres
func GenerateAndInsertGenres(db *sql.DB) {

	count := 100
	response := api.GetData(count)

	actualCount := min(count, response.Total)

	current := "SELECT COUNT(*) FROM library.genres"
	var c int
	err := db.QueryRow(current).Scan(&c)
	if err != nil {
		log.Fatal(err)
	}

	query := "INSERT INTO library.genres (g_id, g_name) VALUES ($1, $2)"

	hmap := make(map[string]bool)

	for i := 0; i < actualCount; i++ {
		genre := response.Data[i].Genre
		_, ok := hmap[genre]
		if !ok {
			_, err := db.Exec(query, i+c+1, genre)
			if err != nil {
				log.Fatalf("Failed to insert data into genres: %v", err)
			}
			hmap[genre] = true
		}
	}

	fmt.Printf("Inserted %d unique genres\n", len(hmap))
}

// Генерация и вставка данных для таблицы subscribers
func GenerateAndInsertSubscribers(db *sql.DB) {

	count := 1000000
	query := "INSERT INTO library.subscribers (s_id, s_name) VALUES ($1, $2)"

	current := "SELECT COUNT(*) FROM library.subscribers"
	var c int
	err := db.QueryRow(current).Scan(&c)
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < count; i++ {
		_, err := db.Exec(query, i+c+1, faker.Name())
		if err != nil {
			log.Fatalf("Failed to insert data into subscribers: %v", err)
		}
	}
	fmt.Printf("Inserted %d subscribers\n", count)
}

// Генерация и вставка данных для таблицы m2m_books_authors
func GenerateAndInsertM2MBooksAuthors(db *sql.DB) {
	count := 1000000
	query := "INSERT INTO library.m2m_books_authors (b_id, a_id) VALUES ($1, $2)"

	maxBookId := "SELECT COUNT(*) FROM library.books"
	maxAuthorId := "SELECT COUNT(*) FROM library.authors"
	var maxBook, maxAuthor int
	err := db.QueryRow(maxBookId).Scan(&maxBook)
	if err != nil {
		log.Fatalf("Failed to get max book id: %v", err)
	}
	err = db.QueryRow(maxAuthorId).Scan(&maxAuthor)
	if err != nil {
		log.Fatalf("Failed to get the max author id: %v", err)
	}

	for i := 0; i < count; i++ {
		bookID := rand.Intn(maxBook) + 1
		authorID := rand.Intn(maxAuthor) + 1
		_, err := db.Exec(query, bookID, authorID)
		if err != nil {
			log.Fatalf("Failed to insert data into m2m_books_authors: %v", err)
		}
	}

	fmt.Printf("Inserted %d m2m_books_authors\n", count)
}

// Генерация и вставка данных для таблицы m2m_books_genres
func GenerateAndInsertM2MBooksGenres(db *sql.DB) {

	count := 1000000

	query := "INSERT INTO library.m2m_books_genres (b_id, g_id) VALUES ($1, $2)"

	maxBookId := "SELECT COUNT(*) FROM library.books"
	maxGenreId := "SELECT COUNT(*) FROM library.genres"
	var maxBook, maxGenre int
	err := db.QueryRow(maxBookId).Scan(&maxBook)
	if err != nil {
		log.Fatalf("Failed to get max book id: %v", err)
	}
	err = db.QueryRow(maxGenreId).Scan(&maxGenre)
	if err != nil {
		log.Fatalf("Failed to get the max author id: %v", err)
	}

	for i := 0; i < count; i++ {
		bookID := rand.Intn(maxBook) + 1
		genreID := rand.Intn(maxGenre) + 1
		_, err := db.Exec(query, bookID, genreID)
		if err != nil {
			log.Fatalf("Failed to insert data into m2m_books_genres: %v", err)
		}
	}

	fmt.Printf("Inserted %d m2m_books_genres\n", count)
}

// Генерация и вставка данных для таблицы subscriptions
func GenerateAndInsertSubscriptions(db *sql.DB) {

	count := 10000000

	query := "INSERT INTO library.subscriptions (sb_id, sb_subscriber, sb_book, sb_start, sb_finish, sb_is_active) VALUES ($1, $2, $3, $4, $5, $6)"

	current := "SELECT COUNT(*) FROM library.subscriptions"
	var c int
	err := db.QueryRow(current).Scan(&c)
	if err != nil {
		log.Fatal(err)
	}

	maxSubscriberID := "SELECT COUNT(*) FROM library.subscribers"
	maxBookID := "SELECT COUNT(*) FROM library.books"
	var maxSubscriber, maxBook int
	err = db.QueryRow(maxSubscriberID).Scan(&maxSubscriber)
	if err != nil {
		log.Fatalf("Failed to get max subscriber id: %v", err)
	}

	err = db.QueryRow(maxBookID).Scan(&maxBook)
	if err != nil {
		log.Fatalf("Failed to get the max book id: %v", err)
	}

	for i := 0; i < count; i++ {
		startDate := gofakeit.Date()

		endDate := gofakeit.DateRange(startDate, startDate.AddDate(2, 0, 0))

		isActive := time.Now().Before(endDate)

		_, err = db.Exec(query,
			i+c+1,
			rand.Intn(maxSubscriber)+1, // Random subscriber ID
			rand.Intn(maxBook)+1,       // Random book ID
			startDate,                  // Start date
			endDate,                    // End date
			isActive,                   // Is active
		)

		if err != nil {
			log.Fatalf("Failed to insert data into subscriptions: %v", err)
		}
	}

	fmt.Printf("Inserted %d subscriptions\n", count)
}
