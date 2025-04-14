package cmd


import (
	"github.com/spf13/cobra"
)

var id uint64

var deleteCmd = &cobra.Command {
	Use: "delete {--id <id> | --help}",
	Short: "Delete expense(s)",
	Run: func(cmd *cobra.Command, args []string) {
	
	},
}

func init() {
	deleteCmd.Flags().Uint64Var(&id, "id", 0, "Unique id of expense to delete use 'list' command to view all expenses")
	rootCmd.AddCommand(deleteCmd)
}
