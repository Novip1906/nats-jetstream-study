package db

import (
	"database/sql"
	"log/slog"
	"os"
	"time"

	"service-api/internal/models"

	_ "github.com/lib/pq"
)

type DB struct {
	*sql.DB
}

func Connect(dbURL string) *DB {
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

	res := &DB{DB: db}
	res.init()

	slog.Info("connected to db")
	return res
}

func (db *DB) init() {
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		username VARCHAR(255) UNIQUE NOT NULL,
		password_hash VARCHAR(255) NOT NULL,
		created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
	)`)
	if err != nil {
		slog.Error("failed to create users table", "error", err)
		os.Exit(1)
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS messages (
		id SERIAL PRIMARY KEY,
		text TEXT NOT NULL,
		created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
	)`)
	if err != nil {
		slog.Error("failed to create messages table", "error", err)
		os.Exit(1)
	}
}

func (db *DB) GetMessages() ([]models.Message, error) {
	rows, err := db.Query("SELECT id, text, created_at FROM messages ORDER BY created_at DESC")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var messages []models.Message
	for rows.Next() {
		var msg models.Message
		if err := rows.Scan(&msg.ID, &msg.Text, &msg.CreatedAt); err != nil {
			return nil, err
		}
		messages = append(messages, msg)
	}

	if messages == nil {
		messages = []models.Message{}
	}

	return messages, nil
}

func (db *DB) DeleteMessage(id string) error {
	_, err := db.Exec("DELETE FROM messages WHERE id = $1", id)
	return err
}

func (db *DB) CreateUser(username, passwordHash string) error {
	_, err := db.Exec("INSERT INTO users (username, password_hash) VALUES ($1, $2)", username, passwordHash)
	return err
}

func (db *DB) GetUserByUsername(username string) (models.User, error) {
	var user models.User
	err := db.QueryRow("SELECT id, username, password_hash FROM users WHERE username = $1", username).
		Scan(&user.ID, &user.Username, &user.PasswordHash)
	return user, err
}
