package main

import (
	"context"
	"database/sql"
	"log"
	"net/http"
	"os"
	"os/signal"
	"server/internal/handler"
	"server/internal/repository/sqlite"
	"server/internal/service"
	"syscall"
	"time"

	"github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	godotenv.Load()
	dbPath := os.Getenv("DB_PATH")
	port := os.Getenv("PORT")

	db, err := sql.Open("sqlite3", dbPath)

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	repo := sqlite.New(db)
	svc := service.New(repo)
	fighterHandler := handler.New(svc)

	mux := http.NewServeMux()
	fighterHandler.RegisterRoutes(mux)

	rl := handler.NewRateLimiter(5, 10)

	server := &http.Server{
		Addr:    port,
		Handler: rl.Middleware(mux),
	}

	go func() {
		log.Printf("Listening on %s\n", port)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %v", err)
		}
	}()

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	<-ctx.Done()
	log.Println("Shutting down")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(shutdownCtx); err != nil {
		log.Fatalf("Shutdown error: %v", err)
	}

	log.Println("Server stopped.")
}
