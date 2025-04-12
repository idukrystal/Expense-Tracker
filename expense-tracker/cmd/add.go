package cmd

import (
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command {
	Use: "add {--description <description> --amount <amount> | --help}",
	Short: "Adds a new expense",
	Run: func(cmd *cobra.Command, args []string) {
		print("add\n")
	},
}

func init() {
	addCmd.Flags().StringP("description", "d", "", "Short description of expense")
	addCmd.Flags().Int64P("amount", "a", 0, "Cost of the expense")
	rootCmd.AddCommand(addCmd)
}
