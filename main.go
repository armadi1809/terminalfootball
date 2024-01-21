package main

import (
	"fmt"
	"log"
	"os"

	"github.com/armadi1809/terminalfootball/footballApiClient"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Could not get api key for the football api client")
	}
	authKey := os.Getenv("AUTH_KEY")
	client := footballApiClient.New(authKey)

	matches, err := client.GetAllTodaysMatches()

	if err != nil {
		log.Fatal("Could not get fixtures")
	}

	for _, match := range matches {
		fmt.Printf("%s %d - %d %s\n", match.HomeTeam.Name, match.Score.FullTime.Home, match.Score.FullTime.Away, match.AwayTeam.Name)
	}
}
