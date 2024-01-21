/*
Copyright © 2024 Aziz Rmadi azizrmadi@gmail.com
*/
package main

import (
	"log"

	terminalfootballcmd "github.com/armadi1809/terminalfootball/cmd"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	terminalfootballcmd.Execute()
}
