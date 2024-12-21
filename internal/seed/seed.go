package seed

import (
	"database/sql"
	"fmt"
	"hw/internal/api"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/brianvoe/gofakeit"
	"github.com/bxcodec/faker/v3"
	_ "github.com/lib/pq"
	"golang.org/x/exp/rand"
)

func GenerateAndInsertAuthors(db *sql.DB) {
	count := 10000
	query := "INSERT INTO library.authors (a_id, a_name) VALUES ($1, $2)"

	// Получаем текущее количество записей в таблице authors
	current := "SELECT COUNT(*) FROM library.authors"
	var c int
	err := db.QueryRow(current).Scan(&c)
	if err != nil {
		log.Fatal(err)
	}

	// Количество записей, которые можно получить за один запрос
	batchSize := 1000
	insertedCount := 0

	for insertedCount < count {
		// Вызываем API для получения данных
		response := api.GetData(batchSize)

		// Проверяем, что API вернул данные
		if len(response.Data) == 0 {
			log.Fatalf("API returned no data, but we expected %d records", batchSize)
		}

		for i := 0; i < batchSize; i++ {
			_, err := db.Exec(query, insertedCount+c+1, response.Data[i].Author)
			if err != nil {
				log.Fatalf("Failed to insert data into authors: %v", err)
			}
			insertedCount++
		}
	}

	fmt.Printf("Inserted %d authors\n", count)
}


func GenerateAndInsertBooks(db *sql.DB) {
	count := 100000
	query := "INSERT INTO library.books (b_id, b_name, b_year, b_quantity) VALUES ($1, $2, $3, $4)"

	current := "SELECT COUNT(*) FROM library.books"
	var c int
	err := db.QueryRow(current).Scan(&c)
	if err != nil {
		log.Fatal(err)
	}

	batchSize := 1000
	insertedCount := 0

	for insertedCount < count {
		response := api.GetData(batchSize)

		if len(response.Data) == 0 {
			log.Fatalf("API returned no data, but we expected %d records", batchSize)
		}

		for i := 0; i < batchSize; i++ {
			date := strings.Split(response.Data[i].Published, "-")
			year, _ := strconv.Atoi(date[0])

			_, err := db.Exec(query, insertedCount+c+1, response.Data[i].Title, year, gofakeit.Number(1, 1000))
			if err != nil {
				log.Fatalf("Failed to insert data into books: %v", err)
			}
			insertedCount++
		}
	}

	fmt.Printf("Inserted %d books\n", count)
}

func GenerateAndInsertGenres(db *sql.DB) {
	count := 1000
	response := api.GetData(count)

	log.Println("Total genres from API:", response.Total)

	// Получаем текущее количество записей в таблице genres
	current := "SELECT COUNT(*) FROM library.genres"
	var c int
	err := db.QueryRow(current).Scan(&c)
	if err != nil {
		log.Fatal(err)
	}

	query := "INSERT INTO library.genres (g_name) VALUES ($1)"

	hmap := make(map[string]bool)
	inserted := 0
	i := 0

	for inserted < count && i < response.Total {
		genre := response.Data[i].Genre
		var exists bool
		err := db.QueryRow("SELECT EXISTS(SELECT 1 FROM library.genres WHERE g_name = $1)", genre).Scan(&exists)
		if err != nil {
			log.Fatalf("Failed to check if genre exists: %v", err)
		}

		if !exists {
			_, err := db.Exec(query, genre)
			if err != nil {
				log.Fatalf("Failed to insert data into genres: %v", err)
			}
			hmap[genre] = true
			inserted++
		}
		i++
	}

	fmt.Printf("Inserted %d unique genres\n", len(hmap))
}

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

	// Получаем количество книг и авторов
	var bookCount, authorCount int
	err := db.QueryRow("SELECT COUNT(*) FROM library.books").Scan(&bookCount)
	if err != nil {
		log.Fatalf("Failed to get book count: %v", err)
	}
	err = db.QueryRow("SELECT COUNT(*) FROM library.authors").Scan(&authorCount)
	if err != nil {
		log.Fatalf("Failed to get author count: %v", err)
	}

	// Проверяем, достаточно ли книг и авторов для создания 1 000 000 комбинаций
	if bookCount*authorCount < count {
		log.Fatalf("Not enough books or authors to create %d unique combinations", count)
	}

	// Генерируем уникальные комбинации
	insertedCount := 0
	for insertedCount < count {
		// Генерируем случайные b_id и a_id
		bookID := rand.Intn(bookCount) + 1
		authorID := rand.Intn(authorCount) + 1

		// Проверяем, существует ли уже такая запись
		existsQuery := "SELECT EXISTS(SELECT 1 FROM library.m2m_books_authors WHERE b_id = $1 AND a_id = $2)"
		var exists bool
		err = db.QueryRow(existsQuery, bookID, authorID).Scan(&exists)
		if err != nil {
			log.Printf("Failed to check if record exists: %v", err)
			continue
		}

		// Если запись не существует, выполняем вставку
		if !exists {
			_, err := db.Exec(query, bookID, authorID)
			if err != nil {
				log.Fatalf("Failed to insert data into m2m_books_authors: %v", err)
			}
			insertedCount++
		}
	}

	fmt.Printf("Inserted %d unique m2m_books_authors\n", insertedCount)
}

// Генерация и вставка данных для таблицы m2m_books_genres
func GenerateAndInsertM2MBooksGenres(db *sql.DB) {
	count := 1000000
	query := "INSERT INTO library.m2m_books_genres (b_id, g_id) VALUES ($1, $2)"

	// Получаем количество книг и жанров
	var bookCount, genreCount int
	err := db.QueryRow("SELECT COUNT(*) FROM library.books").Scan(&bookCount)
	if err != nil {
		log.Fatalf("Failed to get book count: %v", err)
	}
	err = db.QueryRow("SELECT COUNT(*) FROM library.genres").Scan(&genreCount)
	if err != nil {
		log.Fatalf("Failed to get genre count: %v", err)
	}

	// Проверяем, достаточно ли книг и жанров для создания 1 000 000 комбинаций
	if bookCount*genreCount < count {
		log.Fatalf("Not enough books or genres to create %d unique combinations", count)
	}

	// Генерируем уникальные комбинации
	insertedCount := 0
	for insertedCount < count {
		// Генерируем случайные b_id и g_id
		bookID := rand.Intn(bookCount) + 1
		genreID := rand.Intn(genreCount) + 1

		// Проверяем, существует ли уже такая запись
		existsQuery := "SELECT EXISTS(SELECT 1 FROM library.m2m_books_genres WHERE b_id = $1 AND g_id = $2)"
		var exists bool
		err = db.QueryRow(existsQuery, bookID, genreID).Scan(&exists)
		if err != nil {
			log.Printf("Failed to check if record exists: %v", err)
			continue
		}

		// Если запись не существует, выполняем вставку
		if !exists {
			_, err := db.Exec(query, bookID, genreID)
			if err != nil {
				log.Fatalf("Failed to insert data into m2m_books_genres: %v", err)
			}
			insertedCount++
		}
	}

	fmt.Printf("Inserted %d unique m2m_books_genres\n", insertedCount)
}

func GenerateAndInsertSubscriptions(db *sql.DB) {
	count := 10000000
	query := "INSERT INTO library.subscriptions (sb_id, sb_subscriber, sb_book, sb_start, sb_finish, sb_is_active) VALUES ($1, $2, $3, $4, $5, $6)"

	// Получаем текущее количество записей в таблице subscriptions
	current := "SELECT COUNT(*) FROM library.subscriptions"
	var c int
	err := db.QueryRow(current).Scan(&c)
	if err != nil {
		log.Fatal(err)
	}

	// Количество горутин (параллельных вставок)
	numWorkers := 100

	// Канал для передачи задач
	taskChan := make(chan int, count)

	// Канал для сбора результатов
	resultChan := make(chan int, numWorkers)

	// Запускаем горутины
	for w := 0; w < numWorkers; w++ {
		go func() {
			inserted := 0
			for i := range taskChan {
				var subscriberID int
				err := db.QueryRow("SELECT id FROM library.subscribers ORDER BY RANDOM() LIMIT 1").Scan(&subscriberID)
				if err != nil {
					log.Printf("Failed to get random subscriber id: %v", err)
					continue
				}

				var bookID int
				err = db.QueryRow("SELECT id FROM library.books ORDER BY RANDOM() LIMIT 1").Scan(&bookID)
				if err != nil {
					log.Printf("Failed to get random book id: %v", err)
					continue
				}

				startDate := gofakeit.Date()
				endDate := gofakeit.DateRange(startDate, startDate.AddDate(2, 0, 0))
				isActive := time.Now().Before(endDate)

				var active rune
				if isActive {
					active = 'Y'
				} else {
					active = 'N'
				}

				// Выполняем вставку
				_, err = db.Exec(query,
					i+c+1,       // Уникальный sb_id
					subscriberID, // Случайный subscriber ID
					bookID,       // Случайный book ID
					startDate,    // Дата начала
					endDate,      // Дата окончания
					active,       // Активна ли подписка
				)

				if err != nil {
					log.Printf("Failed to insert data into subscriptions: %v", err)
					continue
				}

				inserted++
			}
			resultChan <- inserted
		}()
	}

	// Отправляем задачи в канал
	for i := 0; i < count; i++ {
		taskChan <- i
	}
	close(taskChan)

	// Собираем результаты
	totalInserted := 0
	for w := 0; w < numWorkers; w++ {
		totalInserted += <-resultChan
	}

	fmt.Printf("Inserted %d subscriptions\n", totalInserted)
}