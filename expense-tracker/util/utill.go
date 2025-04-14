package util

import (
	"github.com/idukrystal/Expense-Tracker/expense-tracker/file"
	"fmt"
	"strconv"
	"time"
)

var csvFilePath string

func init() {
	_csvFilePath, err := file.GetCsvFilePath()
	if err != nil {
		print("error filename \n")
	}
	csvFilePath = _csvFilePath
}

func AddExpense(desc string, amt uint64) (uint64, error) {
	expensesCsv, err := file.ReadCsv(csvFilePath)
	if err != nil {
		return 0, err
	}
	id, err := getNextId(expensesCsv)
	if err != nil {
		return 0, err
	}
	date := time.Now().Format("Mon, 2 Jan 2006")
	expensesCsv = append(expensesCsv, []string{
		strconv.FormatUint(id, 10),
		date,
		desc,
		strconv.FormatUint(amt, 10),
	})
	err = file.WriteCsv(csvFilePath, expensesCsv)
	if err != nil {
		return 0, err
	}
	return id, nil
	
}

func DeleteExpense(id uint64) error {
	expensesCsv, err := file.ReadCsv(csvFilePath)
	if err != nil {
		return err
	}
	for i, line := range expensesCsv[1:] {
		lineId, err := strconv.ParseUint(line[0], 10, 64)
		if err != nil {
			return fmt.Errorf("Wrong Csv Format: %w", err)
		}
		
		if lineId == id {
			expensesCsv = append(expensesCsv[:i+1], expensesCsv[i+2:]...)
			return file.WriteCsv(csvFilePath, expensesCsv)
		}
	}
	return fmt.Errorf("(id: %d) Not Found", id)
}


func GetExpenses() ([][]string, error) {
	return file.ReadCsv(csvFilePath)
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
