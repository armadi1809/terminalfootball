/*
Copyright © 2024 Aziz Rmadi azizrmadi@gmail.com
*/
package terminalfootballcmd

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/armadi1809/terminalfootball/cmd/ui"
	"github.com/armadi1809/terminalfootball/footballApiClient"
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
)

var dateLayout string = "2006-01-02"

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "terminalfootball",
	Short: "A CLI to get the latest scores in soccer world",
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

		matches, err := apiClient.GetAllTodaysMatches(dateFlag, dayAfterDateFlag)

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
	rootCmd.PersistentFlags().StringP("date", "d", "", "The date matches are retrieved for")
}
func renderMatches(cmd *cobra.Command, matches []footballApiClient.Match) error {
	rows := []table.Row{}
	for _, match := range matches {
		row := table.Row{match.HomeTeam.Name, strconv.Itoa(match.Score.FullTime.Home) + " - " + strconv.Itoa(match.Score.FullTime.Away), match.AwayTeam.Name}
		rows = append(rows, row)
	}
	table := ui.NewTable(rows)
	if _, err := tea.NewProgram(table, tea.WithOutput(cmd.Root().OutOrStdout()), tea.WithInput(cmd.Root().InOrStdin())).Run(); err != nil {
		return err
	}

	return nil
}

func GetRoot() *cobra.Command {
	return rootCmd
}

func createMatchLeagueMap(matches []footballApiClient.Match) (map[string][]footballApiClient.Match, error) {
	res := make(map[string][]footballApiClient.Match)

	for _, match := range matches {
		res[match.Competition.Name] = append(res[match.Competition.Name], match)
	}
	return res, nil
}
