package cmd

import (
	"fmt"
	"github.com/idukrystal/Expense-Tracker/expense-tracker/util"
	"github.com/spf13/cobra"
	"log"
)


var summaryCmd = &cobra.Command {
	Use: "summary [--month <month> | --help]",
	Short: "Shows sum of expenses",
	Run: func(cmd *cobra.Command, args []string) {
		sum, err := util.SumExpenses(month)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Total expenses: $%d\n", sum)
	},
}

func init() {
	rootCmd.AddCommand(summaryCmd)
}
