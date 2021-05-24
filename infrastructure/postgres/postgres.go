package postgres

import (
	"database/sql"
	"log"
	"math/rand"
	"sync"
	"time"

	_ "github.com/lib/pq"
)

var (
	db   *sql.DB
	once sync.Once
)

// Storage storage
type Storage struct{}

// New returns a new Postgres
func New() Storage {
	getInstance()
	return Storage{}
}

// GetAPoem returns a random poem
func (s Storage) GetAPoem() string {
	stmt, err := db.Prepare("SELECT poem FROM library WHERE id = $1")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	var poem string
	if err := stmt.QueryRow(random()).Scan(&poem); err != nil {
		log.Fatal(err)
	}

	return poem
}

func random() int {
	rand.Seed(time.Now().UnixNano())
	min, max := 1, 4
	return rand.Intn(max-min) + min
}

func getInstance() {
	once.Do(func() {
		var err error
		db, err = sql.Open(
			"postgres",
			"postgres://alejo:123@localhost:5432/hexagonal?sslmode=disable",
		)
		if err != nil {
			log.Fatalf("can't open db: %v", err)
		}

		if err = db.Ping(); err != nil {
			log.Fatalf("can't do ping: %v", err)
		}
	})
}
