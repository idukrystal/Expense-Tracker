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
	if _, err = os.Stat(filePath); err != nil {
		if os.IsNotExist(err) {
			err = os.MkdirAll(filepath.Dir(filePath), 0775)
			if err != nil {
				return
			}
			err = createCsvFile(filePath)
			if err != nil {
				return
			}
		} else {
			return
		}
	}
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



func createCsvFile(filePath string) error {
	var headers = []string{"ID", "Date", "Description", "Amount"}
	
	// Create and write headers
	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("failed to create file: %v", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	if err := writer.Write(headers); err != nil {
		return fmt.Errorf("failed to write headers: %v", err)
	}
	writer.Flush()
	if err := writer.Error(); err != nil {
		return fmt.Errorf("error flushing CSV writer: %v", err)
	}
	return nil
}
