package cmd

import (
	"fmt"
	"github.com/idukrystal/Expense-Tracker/expense-tracker/util"
	"github.com/spf13/cobra"
	"log"
)

var listCmd = &cobra.Command {
	Use: "list [--help | -h]",
	Short: "List expenses",
	Run: func(cmd *cobra.Command, args []string) {
		expensesCsv, err := util.GetExpenses()
		if err != nil {
			log.Fatal(err)
		}
		for _, line := range expensesCsv {
			fmt.Printf("# %s %s %s %s\n",
				line[0],
				line[1],
				line[2],
				line[3],
			)
		}
		
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
