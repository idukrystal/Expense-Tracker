package cmd

import (
	"fmt"
	"github.com/idukrystal/Expense-Tracker/expense-tracker/util"
	"github.com/spf13/cobra"
	"log"
)


var updateCmd = &cobra.Command {
	Use: "update {--id  <id> | --help}",
	Short: "Updates an expense",
	Run: func(cmd *cobra.Command, args []string) {
		err := util.UpdateExpense(id, desc, amt, year, month, day)
		if err != nil {
			log.Fatal(err)
			fmt.Printf("update: %s\n", err)
		}
		fmt.Printf("Update successfull\n")
	},
}

func init() {
	updateCmd.Flags().StringVarP(&desc, "description", "d", "", "Short description of expense")
	
	updateCmd.Flags().Uint64VarP(&amt, "amount", "a", 0, "Cost of the expense")
	updateCmd.Flags().Uint64Var(&id, "id", 0, "Id of expense to update")
	updateCmd.MarkFlagRequired("id")
	updateCmd.Flags().IntVarP(&year, "year", "y", now.Year(), "Year expense was made")
	updateCmd.Flags().IntVarP(&month, "month", "m", int(now.Month()), "Month  expense was made")
	updateCmd.Flags().IntVarP(&day, "day", "n", now.Day(), "Day of month expense was made")
	rootCmd.AddCommand(updateCmd)
}
