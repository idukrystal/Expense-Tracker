package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"time"
)

var (
	now = time.Now()
	month int
	year int
	day int
)

var rootCmd = &cobra.Command {
	Use: "expense {<operation> [--<key> <value> | --help] ... | --help}",
	Short: "A CLI based expense tracker",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	rootCmd.PersistentFlags().IntVarP(&year, "year", "y", now.Year(), "Year expense was made")
	rootCmd.PersistentFlags().IntVarP(&month, "month", "m", int(now.Month()), "Month  expense was made")
	rootCmd.PersistentFlags().IntVarP(&day, "day", "n", now.Day(), "Day of month expense was made")
}

func validateDate(year, month, day int) error {
	if month < 1 || month > 12 {
		return fmt.Errorf("month should be between 1 to 12")
	}
	if day < 0 || day > 31 {
		return fmt.Errorf("day should be between 1 to 31")
	}

	// if a day arg is more than no of days in month time would count up the next month by exceess no of days
	//eg sep 31 results in oct 1
	temp := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.Local)
	if !(year == temp.Year() && time.Month(month) == temp.Month() && day == temp.Day()) {
		return fmt.Errorf("Month %d doesn't have upto %d days in year %d", month, day, year)

	}
	return nil
}


func Execute() {
	rootCmd.Execute()
}
