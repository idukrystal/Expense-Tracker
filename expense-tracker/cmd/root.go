package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"time"
)

var (
	cfgFilePath string
	csvFilePath string
)

var (
	now = time.Now()
	month int
	year int
	day int
)

var rootCmd = &cobra.Command {
	Use: "expense {<operation> [--<key> <value> | --help] ... | --help}",
	Short: "A CLI based expense tracker",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.Flags().StringVar(&cfgFilePath, "config", "", "config file (Default is $XDG_CONFIG_HOME/expense or $HOME/.config/expense)")
	rootCmd.Flags().StringVar(&csvFilePath, "file", "", "csv file to use (Default is $XDG_DATA_HOME/expense or $HOME/.local/share/expense)")
	rootCmd.PersistentFlags().IntVarP(&year, "year", "y", now.Year(), "Yeaur expense was made")
	rootCmd.PersistentFlags().IntVarP(&month, "month", "m", int(now.Month()), "Month  expense was made")
	rootCmd.PersistentFlags().IntVarP(&day, "day", "n", now.Day(), "Day of month expense was made")
	viper.BindPFlag("file", rootCmd.Flags().Lookup("file"))
}

func initConfig() {
	var appConfigHome string
	var configFileName string
	var err error
	
	if cfgFilePath != "" {
		viper.SetConfigFile(cfgFilePath)
	} else {
		configHome := os.Getenv("XDG_CONFIG_HOME")
		if configHome == "" {
			homeDir := os.Getenv("$HOME")
			if homeDir == "" {
				homeDir, err = os.UserHomeDir()
				if err != nil {
					fmt.Println("cant get home dir automatically: %w, and $XDG_DATA_HOME or $HOME not set\n", err)
					os.Exit(1)
				}
			}
			configHome = homeDir+"/.config"
		}
		appConfigHome = configHome+"/expense"
		if err = os.MkdirAll(appConfigHome, 0755); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		
		configFileName = "expense-config.yaml"

		cfgFilePath = appConfigHome+"/"+configFileName
		
		viper.AddConfigPath(appConfigHome)
		viper.SetConfigName(configFileName)
		viper.SetConfigType("yaml")
	}
	if err := viper.ReadInConfig(); err != nil {
		viper.Set("AppName", rootCmd.CalledAs())
		err = viper.SafeWriteConfigAs(cfgFilePath)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
}

func Execute() {
	rootCmd.Execute()
}
