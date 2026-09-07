package main

import (
	"log"
	"net/http"

	"github.com/ShivamNayak-dev/library-management-system/internal/book"
	"github.com/ShivamNayak-dev/library-management-system/internal/config"
	"github.com/ShivamNayak-dev/library-management-system/internal/database"
)

func main() {
	cfg := config.Load()

	db, err := database.Connect(cfg)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer db.Close()

	bookRepository := book.NewRepository(db)
	bookService := book.NewService(bookRepository)
	bookHandler := book.NewHandler(bookService)

	mux := http.NewServeMux()

	mux.HandleFunc("POST /api/v1/books", bookHandler.CreateBook)
	mux.HandleFunc("GET /api/v1/books", bookHandler.GetBooks)

	log.Println("Server running on http://localhost:8080")

	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}