package cmd

import (
	"fmt"
	"github.com/idukrystal/Expense-Tracker/expense-tracker/util"
	"github.com/spf13/cobra"
	"log"
)

var amt uint64
var desc string

var addCmd = &cobra.Command {
	Use: "add {--description <description> --amount <amount> | --help}",
	Short: "Adds a new expense",
	Run: func(cmd *cobra.Command, args []string) {
		id, err := util.AddExpense(desc, amt)
		if err != nil {
			log.Fatal(err)
			fmt.Printf("add: %s\n", err)
		}
		fmt.Printf("Expense added successfully (ID: %d)\n", id)
	},
}

func init() {
	addCmd.Flags().StringVarP(&desc, "description", "d", "", "Short description of expense")
	addCmd.Flags().Uint64VarP(&amt, "amount", "a", 0, "Cost of the expense")
	rootCmd.AddCommand(addCmd)
}
