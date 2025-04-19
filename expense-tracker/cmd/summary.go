package cmd

import (
	"fmt"
	"github.com/idukrystal/Expense-Tracker/expense-tracker/util"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
)


var summaryCmd = &cobra.Command {
	Use: "summary [--month <month> | --help]",
	Short: "Shows sum of expenses",
	Run: func(cmd *cobra.Command, args []string) {
		filters := getFilters(cmd)
		sum, err := util.SumExpenses(viper.GetString("file"), filters...)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Total expenses: $%d\n", sum)
	},
}

func init() {
	rootCmd.AddCommand(summaryCmd)
}
