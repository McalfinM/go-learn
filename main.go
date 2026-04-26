package main

import (
	"context"
	"go-api/internal/config"
	"go-api/internal/server"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/joho/godotenv"
)

func main() {
	// build app (tidak async)

	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found")
	}

	db := config.InitDB()
	app := server.New(db)

	app.Listen(":3001")

	// run server di goroutine
	go func() {
		log.Println("Server running on 0.0.0.0:3001")
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen error: %s\n", err)
		}
	}()

	// graceful shutdown (mirip terminate di TS kamu)
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit
	log.Println("Shutting down server...")

	// timeout shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Println("Server exiting")
}
