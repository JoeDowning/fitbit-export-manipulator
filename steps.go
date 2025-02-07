package main

import (
	"fmt"
	"strconv"
)

func totalStepsByDate(files FilesMap) (StepsCalendar, error) {
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
