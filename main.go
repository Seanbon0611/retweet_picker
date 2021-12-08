package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error Loading .env file")
	}
	ApiKey := os.Getenv("API_KEY")
	ApiSecret := os.Getenv("API_KEY_SECRET")

	fmt.Println(ApiKey)
	fmt.Println(ApiSecret)

}
