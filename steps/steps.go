package steps

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/private/fitbit-export-manipulator/files"
)

func TotalStepsByDate(files files.FilesMap) (StepsCalendar, error) {
	stepsCalendar := StepsCalendar{DateToSteps: make(map[string]int)}

	for _, entries := range files {
		for _, entry := range entries {
			date, err := getEntryDate(entry.DateTime)
			if err != nil {
				return StepsCalendar{}, fmt.Errorf("failed to parse entry date time: %w", err)
			}

			steps, err := strconv.Atoi(entry.Value)
			if err != nil {
				return StepsCalendar{}, fmt.Errorf("failed to convert steps to int: %w", err)
			}

			if _, ok := stepsCalendar.DateToSteps[date]; !ok {
				stepsCalendar.DateToSteps[date] = steps
			} else {
				stepsCalendar.DateToSteps[date] += steps
			}
		}
	}

	return stepsCalendar, nil
}

func AdditionalStepsStats(stepsCal StepsCalendar) {
	var totalSteps int
	extremes := map[string]extremeStat{}

	for date, steps := range stepsCal.DateToSteps {
		if totalSteps == 0 {
			extremes["max"] = extremeStat{
				date:  date,
				steps: steps,
			}
			extremes["min"] = extremeStat{
				date:  date,
				steps: steps,
			}
			totalSteps += steps
			continue
		}

		if extremes["max"].steps < steps {
			extremes["max"] = extremeStat{
				date:  date,
				steps: steps,
			}
		}

		if extremes["min"].steps > steps {
			extremes["min"] = extremeStat{
				date:  date,
				steps: steps,
			}
		}

		totalSteps += steps
	}

	fmt.Printf("-- Additional Stats --\n")
	fmt.Printf("Min steps on %s: %d\n", extremes["min"].date, extremes["min"].steps)
	fmt.Printf("Max steps on %s: %d\n", extremes["max"].date, extremes["max"].steps)
	fmt.Printf("Average steps: %+v per day\n", totalSteps/len(stepsCal.DateToSteps))
	fmt.Printf("Total dates: %d\n", len(stepsCal.DateToSteps))
	fmt.Printf("Total steps: %d\n", totalSteps)
	fmt.Printf("----------------\n")
}

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
