package utils

import (
	"github.com/joho/godotenv"
	"log"
)

func Loadenv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("---failed to load .env file---")
	}
}
