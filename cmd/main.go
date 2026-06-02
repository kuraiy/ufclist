package main

import (
	"database/sql"
	"log"
	"server/internal/handler"
	"server/internal/repository/sqlite"
	"server/internal/service"
)

func main() {
	db, err := sql.Open("sqlite3", "fighters")

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	repo := sqlite.New(db)
	svc := service.New(repo)
	handler := handler.New(svc)
}
