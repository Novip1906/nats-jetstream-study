package broker

import (
	"log/slog"
	"os"
	"time"

	"github.com/nats-io/nats.go"
)

func ConnectAndInit() (*nats.Conn, nats.JetStreamContext) {
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

	_, err = js.AddStream(&nats.StreamConfig{
		Name:     "STREAM_SLOW",
		Subjects: []string{"SLOW.*"},
	})
	if err != nil {
		slog.Error("failed to add stream STREAM_SLOW", "error", err)
	} else {
		slog.Info("stream STREAM_SLOW ready")
	}

	_, err = js.AddStream(&nats.StreamConfig{
		Name:     "STREAM_FAST",
		Subjects: []string{"FAST.*"},
	})
	if err != nil {
		slog.Error("failed to add stream STREAM_FAST", "error", err)
	} else {
		slog.Info("stream STREAM_FAST ready")
	}

	return nc, js
}
