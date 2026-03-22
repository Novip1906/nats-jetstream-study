package main

import (
	"fmt"
	"log/slog"
	"os"
	"sync"
	"time"

	"service-worker/internal/broker"
	"service-worker/internal/config"
	"service-worker/internal/db"
	"service-worker/internal/worker"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	time.Sleep(5 * time.Second)

	cfg, err := config.LoadConfig("config.yaml")
	if err != nil {
		slog.Error("failed to load config", "error", err)
		os.Exit(1)
	}

	dbURL := fmt.Sprintf("postgres://postgres:postgres@%s/messages?sslmode=disable", cfg.DBAddress)
	database := db.ConnectAndInit(dbURL)
	defer database.Close()

	nc, js := broker.ConnectAndInit(cfg.NATSAddress)
	defer nc.Close()

	w := &worker.Worker{
		DB: database,
		JS: js,
	}

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		w.Start("SLOW.new", 5*time.Second, "consumer_slow")
	}()

	go func() {
		defer wg.Done()
		w.Start("FAST.new", 2*time.Second, "consumer_fast")
	}()

	wg.Wait()
}
