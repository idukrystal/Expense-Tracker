package util

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/idukrystal/Expense-Tracker/expense-tracker/file"
)

func init() {
	_csvFilePath, err := file.GetCsvFilePath()
	if err != nil {
		print("error filename \n")
	}
	csvFilePath = _csvFilePath
}

func matchesFilters(csvLine []string, filters ...Filter) bool {
	temp, err := time.Parse(timeLayout, csvLine[1])
	if err != nil {
		log.Print(err)
		return false
	}
	for _, filter := range filters {
		switch filter.Name {
		case "day":
			if !(temp.Day() == filter.Value) { return false }
		case "month":
			if !(temp.Month() == time.Month(filter.Value.(int))) { return false }
		case "year":
			if !(temp.Year() == filter.Value) { return false }
		default:
			return false
		}
	}
	return true
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

