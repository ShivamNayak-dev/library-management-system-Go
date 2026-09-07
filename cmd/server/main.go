package main

import (
	"log"
	"net/http"

	"github.com/ShivamNayak-dev/library-management-system/internal/book"
	"github.com/ShivamNayak-dev/library-management-system/internal/config"
	"github.com/ShivamNayak-dev/library-management-system/internal/database"
	"github.com/ShivamNayak-dev/library-management-system/internal/author"
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


	authorRepository := author.NewRepository(db)
    authorService := author.NewService(authorRepository)
    authorHandler := author.NewHandler(authorService)
	

	mux := http.NewServeMux()

	mux.HandleFunc("POST /api/v1/books", bookHandler.CreateBook)
mux.HandleFunc("GET /api/v1/books", bookHandler.GetBooks)
mux.HandleFunc("GET /api/v1/books/{id}", bookHandler.GetBookByID)
mux.HandleFunc("PUT /api/v1/books/{id}", bookHandler.UpdateBook)
mux.HandleFunc("DELETE /api/v1/books/{id}", bookHandler.DeleteBook)

mux.HandleFunc("POST /api/v1/authors", authorHandler.CreateAuthor)
mux.HandleFunc("GET /api/v1/authors", authorHandler.GetAuthors)
mux.HandleFunc("GET /api/v1/authors/{id}", authorHandler.GetAuthorByID)
mux.HandleFunc("PUT /api/v1/authors/{id}", authorHandler.UpdateAuthor)
mux.HandleFunc("DELETE /api/v1/authors/{id}", authorHandler.DeleteAuthor)

mux.HandleFunc("POST /api/v1/books/{id}/authors", bookHandler.AddAuthor)



	log.Println("Server running on http://localhost:8080")

	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}