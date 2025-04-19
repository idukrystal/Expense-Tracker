package cmd

import (
	"fmt"
	"github.com/idukrystal/Expense-Tracker/expense-tracker/util"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)


var updateCmd = &cobra.Command {
	Use: "update {--id  <id> | --help}",
	Short: "Updates an expense",
	Run: func(cmd *cobra.Command, args []string) {
		filters := getFiltersForUpdate(cmd)
		if len(filters) == 0 {
			fmt.Printf("%s: No flags provided\n", cmd.CalledAs())
			return
		}
		err := util.UpdateExpense(viper.GetString("file"), id, filters...)
		if err != nil {
			fmt.Printf("%s: %s\n", cmd.CalledAs(), err)
			return
		}
		fmt.Printf("Update successfull\n")
	},
}

func init() {
	updateCmd.Flags().StringVarP(&desc, "description", "d", "", "Short description of expense")
	
	updateCmd.Flags().Uint64VarP(&amt, "amount", "a", 0, "Cost of the expense")
	updateCmd.Flags().Uint64Var(&id, "id", 0, "Id of expense to update")
	updateCmd.MarkFlagRequired("id")
	
	rootCmd.AddCommand(updateCmd)
}

func getFiltersForUpdate(cmd *cobra.Command) []util.Filter{
	filters := make([]util.Filter, 0, 6)
	if cmd.Flags().Changed("day") {
		filters = append(filters, util.Filter{Name: "day", Value: day})
	}
	if cmd.Flags().Changed("month") {
		filters = append(filters, util.Filter{Name: "month", Value: month})
	}
	if cmd.Flags().Changed("year") {
		filters = append(filters, util.Filter{Name: "year", Value: year})
	}
	if cmd.Flags().Changed("amount") {
		filters = append(filters, util.Filter{Name: "amount", Value: amt})
	}
	if cmd.Flags().Changed("description") {
		filters = append(filters, util.Filter{Name: "description", Value:desc})
	}
	return filters
}
