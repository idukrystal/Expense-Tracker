package cmd

import (
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


func Execute() {
	rootCmd.Execute()
}
