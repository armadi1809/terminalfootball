/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package terminalfootballcmd

import (
	"log"
	"os"
	"time"

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
		dateFlag, err := cmd.Flags().GetString("date")
		if err != nil {
			log.Fatalf("ubanle to parse the provided date %v", err)
		}
		dayAfterDateFlag := ""
		if dateFlag != "" {
			date, err := time.Parse(dateLayout, dateFlag)
			if err != nil {
				log.Fatalf("error parsing date: %v", err)
				return
			}

			date = date.AddDate(0, 0, 1)
			dayAfterDateFlag = date.Format(dateLayout)
		}
		matches, err := apiClient.GetTodayMatchesForLeagues(args, dateFlag, dayAfterDateFlag)

		if err != nil {
			log.Fatal("Could not get fixtures")
		}

		err = renderMatches(cmd, matches)
		if err != nil {
			log.Fatalf("An Error Occurred When Rendering the Games Table: %s", err.Error())
		}
	},
}

func init() {
	rootCmd.AddCommand(leaguesCmd)
}
