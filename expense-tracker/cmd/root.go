package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
	"time"
)

var (
	cfgFilePath string
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
	rootCmd.PersistentFlags().StringVar(&cfgFilePath, "config", "", "config file (Default is $XDG_CONFIG_HOME/expense or $HOME/.config/expense)")
	rootCmd.PersistentFlags().String("file", "", "csv file to use (Default is $XDG_DATA_HOME/expense or $HOME/.local/share/expense)")
	rootCmd.PersistentFlags().IntVarP(&year, "year", "y", now.Year(), "Yeaur expense was made")
	rootCmd.PersistentFlags().IntVarP(&month, "month", "m", int(now.Month()), "Month  expense was made")
	rootCmd.PersistentFlags().IntVarP(&day, "day", "n", now.Day(), "Day of month expense was made")
	viper.BindPFlag("file", rootCmd.PersistentFlags().Lookup("file"))
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
			configHome = filepath.Join(homeDir, ".config")
		}
		appConfigHome = filepath.Join(configHome, "expense")
		
		configFileName = "expense-config.yaml"

		cfgFilePath = filepath.Join(appConfigHome, configFileName)
		
		viper.AddConfigPath(appConfigHome)
		viper.SetConfigName(configFileName)
		viper.SetConfigType("yaml")
	}
	if err := os.MkdirAll(filepath.Dir(cfgFilePath), 0775); err != nil {
		fmt.Println(err)
	}
	if err := viper.ReadInConfig(); err != nil {
		viper.Set("AppName", rootCmd.CalledAs())
		err = viper.SafeWriteConfigAs(cfgFilePath)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
	if viper.GetString("file") == "" {
		csvFilePath, err := getCsvFilePath()
		if err != nil {
			fmt.Println(err)
		}
		viper.Set("file", csvFilePath)
	}
}

func getCsvFilePath() (string, error) {
	// getFile returns the path to the CSV file, creating it with headers if it doesn't exist
	// Check for XDG_DATA_HOME
	dataHome := os.Getenv("XDG_DATA_HOME")
	if dataHome == "" {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			return "", fmt.Errorf("cannot get home directory: %v", err)
		}
		dataHome = filepath.Join(homeDir, ".local", "share")
	}


	// Construct app-specific path
	appDir := filepath.Join(dataHome, "expense-tracker")

	// Full file path
	filePath := filepath.Join(appDir, "expenses.csv")
	return filePath, nil
}


func Execute() {
	rootCmd.Execute()
}
