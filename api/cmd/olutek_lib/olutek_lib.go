package main

import (
	"fmt"
	"os"

	"github.com/Piokor/olutek_lib/internal/database"
	"github.com/Piokor/olutek_lib/internal/googleapi"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	dbService, err := database.NewDbService()
	if err != nil {
		panic(err)
	}

	bookTitle := os.Args[1]
	books, err := googleapi.GetBooks(bookTitle)
	if err != nil {
		fmt.Print(err)
		return
	}
	dbService.InsertBooks(books)
}
