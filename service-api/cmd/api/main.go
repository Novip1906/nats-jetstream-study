package main

import (
	"log/slog"
	"os"
	"service-api/internal/broker"
	"service-api/internal/db"
	"service-api/internal/handlers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	database := db.Connect()
	defer database.Close()

	nc, js := broker.Connect()
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
