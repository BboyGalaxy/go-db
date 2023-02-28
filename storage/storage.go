package storage

import (
	"database/sql"
	"fmt"
	"log"
	"sync"

	_ "github.com/lib/pq"
)

var (
	db   *sql.DB
	once sync.Once
)

func NewPostgresDB() {
	once.Do(func() {
		var err error
		db, err = sql.Open("postgres", "postgres://ingadg:Atiradeonhd23@localhost:5432/godb?sslmode=disable")
		if err != nil {
			log.Fatalf("Can't open DB: %v", err)
		}

		if err = db.Ping(); err != nil {
			log.Fatalf("Can't do ping: %v", err)
		}
		fmt.Println("Conectado a postgres")
	})
}

// Pool return a unique instance of db
func Pool() *sql.DB {
	return db
}
