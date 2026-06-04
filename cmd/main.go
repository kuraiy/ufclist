package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"server/internal/handler"
	"server/internal/repository"
	"server/internal/service"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	loadEnvVars()

	db := connectToDb()

	frepo := repository.New(db)
	fsvc := service.New(frepo)
	v := validator.New()
	fh := handler.New(fsvc, v)

	startServer(fh)
}

func loadEnvVars() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
}

func connectToDb() *sql.DB {
	dbPath := os.Getenv("DB_PATH")

	db, err := sql.Open("sqlite3", dbPath)

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	return db
}

func startServer(fh *handler.FighterHandler) {
	port := os.Getenv("PORT")

	mux := http.NewServeMux()
	fh.RegisterRoutes(mux)

	if err := http.ListenAndServe(port, mux); err != nil {
		log.Fatal(err)
	}

	log.Printf("Listening on %s\n", port)
}
