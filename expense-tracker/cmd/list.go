package cmd

import (
	"fmt"
	"github.com/idukrystal/Expense-Tracker/expense-tracker/util"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
)

var exportFile string

var listCmd = &cobra.Command {
	Use: "list [--help | -h]",
	Short: "List expenses",
	Run: func(cmd *cobra.Command, args []string) {
		filters := getFilters(cmd)
		expensesCsv, err := util.GetExpenses(viper.GetString("file"), filters...)
		if err != nil {
			log.Fatal(err)
		}
		if cmd.Flags().Changed("export") {
			err := util.ExportCsv(exportFile, expensesCsv)
			if err != nil {
				fmt.Printf("Could not export to %s: %v\n", exportFile, err)
			}
			fmt.Printf("Exported succesfully to %s\n", exportFile)
		} else {
			for _, line := range expensesCsv {
				fmt.Printf("# %s %s %s $%s\n",
					line[0],
					line[1],
					line[2],
					line[3],
				)
			}
		}
		
	},
}

func getFilters(cmd *cobra.Command) []util.Filter{
	filters := make([]util.Filter, 0 ,4)
	if cmd.Flags().Changed("month") || cmd.Flags().Changed("day") {
		filters = append(filters, util.Filter{Name: "month", Value: month})
	}
	if cmd.Flags().Changed("year") || cmd.Flags().Changed("month") || cmd.Flags().Changed("day") {
		filters = append(filters, util.Filter{Name: "year", Value: year})
	}
	if cmd.Flags().Changed("day") {
		filters = append(filters, util.Filter{Name: "day", Value: day})
	}
	return filters
}

func init() {
	listCmd.Flags().StringVarP(&exportFile, "export", "e", "", "File to export to")
	rootCmd.AddCommand(listCmd)
}
