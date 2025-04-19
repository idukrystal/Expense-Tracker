package cmd

import (
	"fmt"
	"github.com/idukrystal/Expense-Tracker/expense-tracker/util"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	amt uint64
	desc string
)

var addCmd = &cobra.Command {
	Use: "add {--description <description> --amount <amount> | --help}",
	Short: "Adds a new expense",
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		var id uint64
		
		// Ensure proper use of the --day, --month and --year flags
		if err = validateDateFlags(cmd); err == nil {
			id, err = util.AddExpense(viper.GetString("file"), desc, amt, year, month, day)
			if err == nil {
				fmt.Printf("Expense added successfully (ID: %d)\n", id)
				return
			}
		}	

		// This only runs if err ! nill in one of the if statements
		fmt.Printf("%s: %v\n", cmd.CalledAs(),  err)
	},
}

func validateDateFlags(cmd *cobra.Command) error {
	if err := util.ValidateDate(year, month, day); err != nil {
		return err
	}


	if cmd.Flags().Changed("year") {
		if !cmd.Flags().Changed("month") {
			return fmt.Errorf("--month and --day flags  must be provided if --year is provided")
		}
	}
	if cmd.Flags().Changed("month") {
		if !cmd.Flags().Changed("day") {
			return fmt.Errorf("--day flag must be provided if --month is provided")
		}
	}
	return nil
}

func init() {
	addCmd.Flags().StringVarP(&desc, "description", "d", "", "Short description of expense")
	addCmd.MarkFlagRequired("description")
	
	addCmd.Flags().Uint64VarP(&amt, "amount", "a", 0, "Cost of the expense")
	addCmd.MarkFlagRequired("amount")
	
	rootCmd.AddCommand(addCmd)
}

