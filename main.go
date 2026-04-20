package main

import (
	"os"

	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()

	addr := os.Getenv("ADDR")
	if addr == "" {
		addr = "8080"
	}

	api := NewAPI()
	err := api.Start(addr)
	if err != nil {
		return
	}
}
