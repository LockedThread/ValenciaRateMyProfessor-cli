package main

import (
	"cli/commands"
	"log"
)

func main() {
	commands.RootInit()
	err := commands.RootExecute()
	if err != nil {
		log.Fatalln(err)
	}
}
