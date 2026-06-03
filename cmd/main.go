package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"server/internal/handler"
	"server/internal/repository/sqlite"
	"server/internal/service"

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
	handler := handler.New(svc)

	mux := http.NewServeMux()
	handler.RegisterRoutes(mux)

	log.Printf("Listening on %s\n", port)
	if err := http.ListenAndServe(port, mux); err != nil {
		log.Fatal(err)
	}

}
