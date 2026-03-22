package main

import (
	"fmt"
	"log/slog"
	"os"
	"service-api/internal/broker"
	"service-api/internal/config"
	"service-api/internal/db"
	"service-api/internal/handlers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	cfg, err := config.LoadConfig("config.yaml")
	if err != nil {
		slog.Error("failed to load config", "error", err)
		os.Exit(1)
	}

	dbURL := fmt.Sprintf("postgres://postgres:postgres@%s/messages?sslmode=disable", cfg.DBAddress)
	database := db.Connect(dbURL)
	defer database.Close()

	nc, js := broker.Connect(cfg.NATSAddress)
	defer nc.Close()

	h := &handlers.Handler{
		DB: database,
		JS: js,
	}

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"}
	corsConfig.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization"}
	r.Use(cors.New(corsConfig))

	r.POST("/api/register", h.Register)
	r.POST("/api/login", h.Login)

	protected := r.Group("/")
	protected.Use(handlers.AuthMiddleware())
	protected.POST("/api/messages", h.PublishMessage)
	protected.GET("/api/messages", h.GetMessages)
	protected.DELETE("/api/messages/:id", h.DeleteMessage)

	r.Run(":8080")
}
