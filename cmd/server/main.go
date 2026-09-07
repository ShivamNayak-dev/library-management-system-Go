package main

import (
	"fmt"
	"log"

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

	fmt.Println("Library Management System")
	fmt.Println("Database connected successfully!")
}