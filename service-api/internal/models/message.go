package models

import "time"

type Message struct {
	ID        int       `json:"id"`
	Text      string    `json:"text"`
	CreatedAt time.Time `json:"created_at"`
}

type PublishRequest struct {
	Text string `json:"text" binding:"required"`
	Type string `json:"type" binding:"required"`
}
