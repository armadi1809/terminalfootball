/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package terminalfootballcmd

import (
	"fmt"
	"log"
	"os"

	"github.com/armadi1809/terminalfootball/footballApiClient"
	"github.com/spf13/cobra"
)

// leaguesCmd represents the leagues command
var leaguesCmd = &cobra.Command{
	Use:   "leagues",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		apiClient := footballApiClient.New(os.Getenv("AUTH_KEY"))
		matches, err := apiClient.GetTodayMatchesForLeagues(args)

		if err != nil {
			log.Fatal("Could not get fixtures")
		}
		for _, match := range matches {
			fmt.Printf("%s %d - %d %s\n", match.HomeTeam.Name, match.Score.FullTime.Home, match.Score.FullTime.Away, match.AwayTeam.Name)
		}
	},
}

func init() {
	RootCmd.AddCommand(leaguesCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// leaguesCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// leaguesCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	leaguesCmd.Flags().String("leagues", "", "Define the leagues you want to see matches of")
}
