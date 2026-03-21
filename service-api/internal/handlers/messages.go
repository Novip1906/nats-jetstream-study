package handlers

import (
	"database/sql"
	"log/slog"
	"net/http"
	"service-api/internal/models"

	"github.com/gin-gonic/gin"
	"github.com/nats-io/nats.go"
)

type Handler struct {
	DB *sql.DB
	JS nats.JetStreamContext
}

func (h *Handler) PublishMessage(c *gin.Context) {
	var req models.PublishRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		slog.Error("invalid request", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	subject := "SLOW.new"
	if req.Type == "fast" {
		subject = "FAST.new"
	}

	_, err := h.JS.Publish(subject, []byte(req.Text))
	if err != nil {
		slog.Error("failed to publish", "subject", subject, "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	slog.Info("message published", "subject", subject, "text", req.Text)
	c.JSON(http.StatusOK, gin.H{"status": "published"})
}

func (h *Handler) GetMessages(c *gin.Context) {
	rows, err := h.DB.Query("SELECT id, text, created_at FROM messages ORDER BY created_at DESC")
	if err != nil {
		slog.Error("failed to query messages", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var messages []models.Message
	for rows.Next() {
		var msg models.Message
		if err := rows.Scan(&msg.ID, &msg.Text, &msg.CreatedAt); err != nil {
			slog.Error("failed to scan row", "error", err)
			continue
		}
		messages = append(messages, msg)
	}

	if messages == nil {
		messages = []models.Message{}
	}

	c.JSON(http.StatusOK, messages)
}

func (h *Handler) DeleteMessage(c *gin.Context) {
	id := c.Param("id")
	_, err := h.DB.Exec("DELETE FROM messages WHERE id = $1", id)
	if err != nil {
		slog.Error("failed to delete message", "id", id, "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	slog.Info("message deleted", "id", id)
	c.JSON(http.StatusOK, gin.H{"status": "deleted"})
}
