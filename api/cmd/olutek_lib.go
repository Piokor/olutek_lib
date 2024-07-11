package main

import (
	"os"

	"github.com/Piokor/olutek_lib/internal/api"
	"github.com/Piokor/olutek_lib/internal/database"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	dbService, err := database.NewDbService()
	if err != nil {
		panic(err)
	}

	router := api.GetRouter(dbService)
	router.Run(os.Getenv("DOMAIN"))
}
