package main

import (
	"log"

	"wokey/app"

	// _ "github.com/Rambatino/gooff"/
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Print("Error loading .env file")
	}

	app.Init()
	StartServer()
}
