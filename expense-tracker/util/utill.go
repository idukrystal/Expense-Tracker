package util

import (
	"fmt"
	"strconv"
	"time"

	"github.com/idukrystal/Expense-Tracker/expense-tracker/file"
)

type Filter struct {
	Name string
	Value any
}

func (f Filter) String() string {
	return fmt.Sprintf("%s: %v", f.Name, f.Value.(int))
}

var timeLayout string = "2006-01-02"

func AddExpense(csvFilePath string, desc string, amt uint64, year , month, day int) (uint64, error) {
	expensesCsv, err := file.ReadCsv(csvFilePath)
	if err != nil {
		return 0, err
	}
	id, err := getNextId(expensesCsv)
	if err != nil {
		return 0, err
	}
	date := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.Local).Format(timeLayout)
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

func DeleteExpense(csvFilePath string, id uint64) error {
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

func GetExpenses(csvFilePath string, filters ...Filter) ([][]string, error) {
	fmt.Println("filters: ", filters)
	allExpensesCsv, err := file.ReadCsv(csvFilePath)
	if err != nil {
		return nil, err
	}
	if len(filters) < 1 {
		return allExpensesCsv, nil
	}
	filteredExpensesCsv := make([][]string, 0, len(allExpensesCsv))
	filteredExpensesCsv = append(filteredExpensesCsv, allExpensesCsv[0])
	for _, line := range allExpensesCsv[1:] {
		if matchesFilters(line, filters...) {
			filteredExpensesCsv = append(filteredExpensesCsv, line)
		}
	}
	return filteredExpensesCsv, nil
}

func SumExpenses(csvFilePath string, filters ...Filter) (uint64, error) {
	expensesCsv, err := GetExpenses(csvFilePath, filters...)
	if err != nil {
		return 0, err
	}
	var sum uint64
	for _, line := range expensesCsv[1:] {
		curAmt, err := strconv.ParseUint(line[3], 10, 64)
		if err != nil {
			return 0, err
		}
		sum += curAmt
	}
	return sum, nil
}

func UpdateExpense(csvFilePath string, id uint64, filters ...Filter) error {
	expensesCsv, err := file.ReadCsv(csvFilePath)
	if err != nil {
		return err
	}
	for i,  line := range expensesCsv {
		if i == 0 { continue }
		lineId, err := strconv.ParseUint(line[0], 10, 64)
		if err != nil {
			return fmt.Errorf("Error with CSV file: %w", err)
		}
		if id == lineId {
			newLine, err := updateUsingFilters(line, filters...)
			if err != nil {
				return err
			}
			expensesCsv = append(append(expensesCsv[:i], newLine), expensesCsv[i+1:]...)
			err = file.WriteCsv(csvFilePath, expensesCsv)
			if err != nil {
				return err
			}
			return nil 
		}
	}
	return fmt.Errorf("(Id: %d) Not Found", id)
}

func ExportCsv(filePath string, data [][]string) error {
	err := file.WriteCsv(filePath, data)
	if err != nil {
		return err
	}
	return nil
}
