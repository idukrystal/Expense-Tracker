package cmd

import (
	"github.com/spf13/cobra"
)

var summaryCmd = &cobra.Command {
	Use: "summary [--month <month> | --help]",
	Short: "Shows sum of expenses",
	Run: func(cmd *cobra.Command, args []string) {
		print("summary\n")
	},
}

func init() {
	rootCmd.AddCommand(summaryCmd)
}
