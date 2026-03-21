package broker

import (
	"log/slog"
	"os"
	"time"

	"github.com/nats-io/nats.go"
)

func Connect() (*nats.Conn, nats.JetStreamContext) {
	var nc *nats.Conn
	var err error
	for i := 0; i < 5; i++ {
		nc, err = nats.Connect("nats:4222")
		if err == nil {
			break
		}
		time.Sleep(2 * time.Second)
	}
	if err != nil {
		slog.Error("failed to connect to nats", "error", err)
		os.Exit(1)
	}

	js, err := nc.JetStream()
	if err != nil {
		slog.Error("failed to get jetstream context", "error", err)
		os.Exit(1)
	}

	slog.Info("connected to nats jetstream")
	return nc, js
}
