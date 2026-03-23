package worker

import (
	"log/slog"
	"service-worker/internal/db"
	"time"

	"github.com/nats-io/nats.go"
)

type Worker struct {
	DB *db.DB
	JS nats.JetStreamContext
}

func (w *Worker) Start(subject string, delay time.Duration, durableName string) {
	sub, err := w.JS.SubscribeSync(subject, nats.Durable(durableName))
	if err != nil {
		slog.Error("failed to subscribe", "subject", subject, "error", err)
		return
	}
	
	slog.Info("worker started", "subject", subject)

	for {
		m, err := sub.NextMsg(10 * time.Minute)
		if err != nil {
			if err != nats.ErrTimeout {
				slog.Error("error fetching next msg", "subject", subject, "error", err)
			}
			continue
		}

		text := string(m.Data)
		slog.Info("received msg", "subject", subject, "text", text)
		
		time.Sleep(delay)

		err = w.DB.SaveMessage(text)
		if err != nil {
			slog.Error("failed to insert msg", "text", text, "error", err)
			continue
		}

		m.Ack()
		slog.Info("processed msg", "subject", subject, "text", text)
	}
}
