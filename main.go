package main

import (
	"github.com/joho/godotenv"
	"log"
	"school-grades-calc/inputs"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	inputs.Menu()
}
