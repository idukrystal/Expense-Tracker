package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command {
	Use: "expense {<operation> [--<key> <value> | --help] ... | --help}",
	Short: "A CLI based expense tracker",
	Run: func(cmd *cobra.Command, args []string) {
		print("root\n")
	},
}

func Execute() {
	rootCmd.Execute()
}
