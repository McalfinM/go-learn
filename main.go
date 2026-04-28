package main

import (
	"context"
	"go-api/internal/config"
	"go-api/internal/server"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)

func main() {
	// load env
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	// init db
	db := config.InitDB()

	// init app
	app := server.New(db)
	app.Use(logger.New())

	// run server di goroutine
	go func() {
		if err := app.Listen(":3001"); err != nil {
			log.Println("Server stopped:", err)
		}
	}()

	log.Println("Server running on :3001")

	// graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit
	log.Println("Shutting down server...")

	// timeout shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := app.ShutdownWithContext(ctx); err != nil {
		log.Println("Server forced shutdown:", err)
	}

	// close DB (PENTING)
	if err := db.Close(); err != nil {
		log.Println("Error closing DB:", err)
	}

	log.Println("Server exiting")
}
