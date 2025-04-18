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

var csvFilePath string
var timeLayout string = "2006-01-02"

func AddExpense(desc string, amt uint64, year , month, day int) (uint64, error) {
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

func GetExpenses(filters ...Filter) ([][]string, error) {
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

func SumExpenses(filters ...Filter) (uint64, error) {
	expensesCsv, err := GetExpenses(filters...)
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

func UpdateExpense(id uint64, desc string, amt uint64, year, month, day int) error {
	return nil
}
