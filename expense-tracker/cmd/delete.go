package cmd


import (
	"fmt"
	"github.com/idukrystal/Expense-Tracker/expense-tracker/util"
	"github.com/spf13/cobra"
	"log"
)

var id uint64

var deleteCmd = &cobra.Command {
	Use: "delete {--id <id> | --help}",
	Short: "Delete expense(s)",
	Run: func(cmd *cobra.Command, args []string) {
		err := util.DeleteExpense(id)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Expense deleted successfully")
	},
}

func init() {
	deleteCmd.Flags().Uint64Var(&id, "id", 0, "Unique id of expense to delete use 'list' command to view all expenses")
	rootCmd.AddCommand(deleteCmd)
}
