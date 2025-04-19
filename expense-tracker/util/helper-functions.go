package util

import (
	"fmt"
	"log"
	"strconv"
	"time"
)


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

func updateUsingFilters(csvLine []string, filters ...Filter) ([]string, error) {
	if len(filters) == 0 {
		return nil, fmt.Errorf("No values speciefied")
	}
	newLine := make([]string, len(csvLine))
	temp, err := time.Parse(timeLayout, csvLine[1])
	if err != nil {
		log.Print(err)
		return nil, err
	}
	for index, value := range csvLine {
		newLine[index] = value
	}
	for _, filter := range filters {
		switch filter.Name {
		case "description":
			newLine[2] = filter.Value.(string)
		case "amount":
			newLine[3] = strconv.FormatUint(filter.Value.(uint64), 10)
		case "day":
			err = ValidateDate(temp.Year(), int(temp.Month()), filter.Value.(int))
			temp = time.Date(temp.Year(), temp.Month(), filter.Value.(int), 0, 0, 0, 0, time.Local)
		case "month":
			err = ValidateDate(temp.Year(), int(time.Month(filter.Value.(int))), temp.Day())
			temp = time.Date(temp.Year(), time.Month(filter.Value.(int)), temp.Day(), 0, 0, 0, 0, time.Local)
		case "year":
			err = ValidateDate(filter.Value.(int), int(temp.Month()), temp.Day())
			temp = time.Date(filter.Value.(int), temp.Month(), temp.Day(), 0, 0, 0, 0, time.Local)
			
		default:
			return nil, fmt.Errorf("Unknown Value(%s):", filter.Name)
		}
		if err != nil {
			return nil, err
		}
	}
	newLine[1] = temp.Format(timeLayout)
	return newLine, nil
}

func getNextId(expensesCsv [][]string) (uint64, error) {
	var higestId uint64 = 0
	for index, line := range expensesCsv {
		if index == 0 {
			continue
		}
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

func ValidateDate(year, month, day int) error {
	if month < 1 || month > 12 {
		return fmt.Errorf("month should be between 1 to 12")
	}
	if day < 0 || day > 31 {
		return fmt.Errorf("day should be between 1 to 31")
	}

	// if a day arg is more than no of days in month time would count up the next month by exceess no of days
	//eg sep 31 results in oct 1
	temp := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.Local)
	if !(year == temp.Year() && time.Month(month) == temp.Month() && day == temp.Day()) {
		return fmt.Errorf("Month %d doesn't have upto %d days in year %d", month, day, year)

	}
	return nil
}
