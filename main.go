package main

import (
	"cli/commands"
	"cli/database"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	database.Client = *database.Connect()
	commands.RootInit()
	err = commands.RootExecute()
	if err != nil {
		log.Fatalln(err)
	}
}
