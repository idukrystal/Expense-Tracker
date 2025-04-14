package util

import (
	"encoding/csv"
	"path/filepath"
	"fmt"
	"os"
	"strconv"
	"time"
)

var headers = []string{"ID", "Date", "Description", "Amount"}

func getOpenCsvFile() (*os.File, error) {
	csvFilePath, err := getCsvFilePath()
	if err != nil {
		return nil, err
	}
	csvFile, err := os.OpenFile(csvFilePath, os.O_RDWR, 0644)
	if err != nil {
		return nil, err
	}
	return csvFile, nil
}

func GetExpensesCsv() ([][]string, error) {
	csvFile, err := getOpenCsvFile()
	if err  != nil {
		return nil, err
	}
	defer csvFile.Close()
	reader := csv.NewReader(csvFile)
	return reader.ReadAll()
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

func AddExpense(desc string, amt uint64) (uint64, error) {
	expenses, err := GetExpensesCsv()
	if err != nil {
		return 0, err
	}
	id, err := getNextId(expenses)
	if err != nil {
		return 0, err
	}
	date := time.Now().Format("Mon, 2 Jan 2006")
	expenses = append(expenses, []string{
		strconv.FormatUint(id, 10),
		date,
		desc,
		strconv.FormatUint(amt, 10)})
	csvFile, err := getOpenCsvFile()
	if err != nil {
		return 0, err
	}
	defer csvFile.Close()
	csvWritter := csv.NewWriter(csvFile)
	err = csvWritter.WriteAll(expenses)
	if err != nil {
		return 0, err
	}
	if err = csvWritter.Error(); err!= nil {
		return 0, err
	}
	fmt.Println(expenses)
	return id, nil
	
}

func getNextId(expensesCsv [][]string) (uint64, error) {
	var higestId uint64 = 0
	for _, line := range expensesCsv[1:] {
		id, err := strconv.ParseUint(line[0], 10, 64)
		if err != nil{
			return 0, fmt.Errorf("Error with csv file: %w", err)
		}
		if id > higestId {
			higestId = id
		}
	}
	return (higestId + 1), nil
}
