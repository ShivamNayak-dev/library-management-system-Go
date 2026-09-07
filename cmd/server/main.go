package main

import (
	"log"
	"net/http"

	"github.com/ShivamNayak-dev/library-management-system/internal/author"
	"github.com/ShivamNayak-dev/library-management-system/internal/book"
	"github.com/ShivamNayak-dev/library-management-system/internal/category"
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

	authorRepository := author.NewRepository(db)
	authorService := author.NewService(authorRepository)
	authorHandler := author.NewHandler(authorService)

	categoryRepository := category.NewRepository(db)
	categoryService := category.NewService(categoryRepository)
	categoryHandler := category.NewHandler(categoryService)

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

	mux.HandleFunc("POST /api/v1/categories", categoryHandler.CreateCategory)
	mux.HandleFunc("GET /api/v1/categories", categoryHandler.GetCategories)
	mux.HandleFunc("GET /api/v1/categories/{id}", categoryHandler.GetCategoryByID)
	mux.HandleFunc("PUT /api/v1/categories/{id}", categoryHandler.UpdateCategory)
	mux.HandleFunc("DELETE /api/v1/categories/{id}", categoryHandler.DeleteCategory)

	log.Println("Server running on http://localhost:8080")

	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}
