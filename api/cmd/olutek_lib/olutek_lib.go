package main

import (
	"fmt"
	"os"

	"github.com/Piokor/olutek_lib/internal/bookapi/googleapi"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	bookTitle := os.Args[1]
	volumes, err := googleapi.GetVolumes(bookTitle)
	if err != nil {
		fmt.Print(err)
		return
	}
	fmt.Print(volumes)
}
