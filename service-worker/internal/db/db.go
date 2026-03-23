package db

import (
	"database/sql"
	"log/slog"
	"os"
	"time"

	_ "github.com/lib/pq"
)

type DB struct {
	*sql.DB
}

func ConnectAndInit(dbURL string) *DB {
	var db *sql.DB
	var err error
	for i := 0; i < 5; i++ {
		db, err = sql.Open("postgres", dbURL)
		if err == nil {
			err = db.Ping()
			if err == nil {
				break
			}
		}
		time.Sleep(2 * time.Second)
	}
	if err != nil {
		slog.Error("failed to connect to db", "error", err)
		os.Exit(1)
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS messages (
		id SERIAL PRIMARY KEY,
		text TEXT NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	)`)
	if err != nil {
		slog.Error("failed to create table", "error", err)
		os.Exit(1)
	}

	slog.Info("db initialized successfully")
	return &DB{DB: db}
}

func (db *DB) SaveMessage(text string) error {
	_, err := db.Exec("INSERT INTO messages (text) VALUES ($1)", text)
	return err
}
