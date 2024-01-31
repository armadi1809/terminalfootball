/*
Copyright © 2024 Aziz Rmadi azizrmadi@gmail.com
*/
package terminalfootballcmd

import (
	"log"
	"os"
	"strconv"

	"github.com/armadi1809/terminalfootball/cmd/ui"
	"github.com/armadi1809/terminalfootball/footballApiClient"
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
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

		err = renderMatches(cmd, matches)
		if err != nil {
			log.Fatalf("An Error Occurred When Rendering the Games Table: %s", err.Error())
		}
	},
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
