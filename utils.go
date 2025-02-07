package main

import (
	"fmt"
	"sort"
	"strings"
)

const dateTimeFormat = "2024/01/13"

func getEntryDate(dateTime string) (string, error) {
	dateSplit := strings.Split(dateTime, " ")
	if len(dateSplit) != 2 {
		return "", fmt.Errorf("invalid date time format")
	}

	date := dateSplit[0]
	monthDayYear := strings.Split(date, "/")
	if monthDayYear == nil || len(monthDayYear) != 3 {
		return "", fmt.Errorf("invalid date format")
	}

	month := monthDayYear[0]
	if len(month) < 2 {
		month = "0" + month
	}

	day := monthDayYear[1]
	if len(day) < 2 {
		day = "0" + day
	}

	year := monthDayYear[2]
	if len(year) == 2 {
		year = "20" + year
	}

	return fmt.Sprintf("%s/%s/%s", year, month, day), nil
}

func printInOrder(m map[string]int) {
	var dates []string
	for date := range m {
		dates = append(dates, date)
	}

	sort.Strings(dates)

	for _, date := range dates {
		fmt.Printf("%s: %d\n", date, m[date])
	}
}
