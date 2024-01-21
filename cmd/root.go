/*
Copyright © 2024 Aziz Rmadi azizrmadi@gmail.com
*/
package terminalfootballcmd

import (
	"fmt"
	"log"
	"os"

	"github.com/armadi1809/terminalfootball/footballApiClient"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "terminalfootball",
	Short: "A CLI to get the latest scores in soccer world",
	Run: func(cmd *cobra.Command, args []string) {
		apiClient := footballApiClient.New(os.Getenv("AUTH_KEY"))
		matches, err := apiClient.GetAllTodaysMatches()

		if err != nil {
			log.Fatal("Could not get fixtures")
		}

		for _, match := range matches {
			fmt.Printf("%s %d - %d %s\n", match.HomeTeam.Name, match.Score.FullTime.Home, match.Score.FullTime.Away, match.AwayTeam.Name)
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.terminalfootball.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
}
