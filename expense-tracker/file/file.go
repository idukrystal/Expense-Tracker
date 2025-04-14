package file

import (
	"encoding/csv"
	"fmt"
	"os"
	"path/filepath"
)

type Usage string
const (
	READ Usage = "READ" 
	WRITE Usage = "write"
)

func getOpenFile(filePath string, use Usage) (file *os.File, err error) {
	switch use {
	case READ:
		file, err = os.Open(filePath)
	case WRITE:
		file, err = os.OpenFile(filePath, os.O_WRONLY | os.O_TRUNC, 0644)
	default:
		file, err = nil, fmt.Errorf("Unknow use case for file")
	}
	return 
}

func ReadCsv(csvFilePath string) ([][]string, error) {
	csvFile, err := getOpenFile(csvFilePath, READ)
	if err  != nil {
		return nil, err
	}
	defer csvFile.Close()
	reader := csv.NewReader(csvFile)
	return reader.ReadAll()
}

func WriteCsv(csvFilePath string, csvString [][]string) error {
	csvFile, err := getOpenFile(csvFilePath, WRITE)
	if err != nil {
		return err
	}
	defer csvFile.Close()
	csvWritter := csv.NewWriter(csvFile)
	err = csvWritter.WriteAll(csvString)
	if err != nil {
		return err
	}
	if err = csvWritter.Error(); err!= nil {
		return err
	}
	return nil
}

func GetCsvFilePath() (string, error) {
	var headers = []string{"ID", "Date", "Description", "Amount"}
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
	if err := os.MkdirAll(appDir, 0755); err != nil {
		return "", fmt.Errorf("failed to create directory: %v", err)
	}

	// Full file path
	filePath := filepath.Join(appDir, "expenses.csv")

	// Check if file exists
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		// Create and write headers
		file, err := os.Create(filePath)
		if err != nil {
			fmt.Printf("create file for first time: %w ", err)
			return "", fmt.Errorf("failed to create file: %v", err)
		}
		defer file.Close()

		writer := csv.NewWriter(file)
		if err := writer.Write(headers); err != nil {
			return "", fmt.Errorf("failed to write headers: %v", err)
		}
		writer.Flush()
		if err := writer.Error(); err != nil {
			return "", fmt.Errorf("error flushing CSV writer: %v", err)
		}
	}

	return filePath, nil
}
