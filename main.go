package main

import (
	"fmt"

	"github.com/private/fitbit-export-manipulator/files"
	"github.com/private/fitbit-export-manipulator/output"
	"github.com/private/fitbit-export-manipulator/steps"
)

const (
	// change for other locations, currently pointing inside the project
	// e.g. "path/" -> "fitbit-export-manipulator/data/"
	path = "data/"

	outputMode = "print"
	printOrder = "newFirst"
)

func main() {
	fmt.Printf("Starting...\n")
	filesMap, err := files.ExtractJSONFiles(path)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%d Files extracted\n", len(filesMap))

	stepsCalendar, err := steps.TotalStepsByDate(filesMap)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Collated into %d different dates\n", len(stepsCalendar.DateToSteps))

	switch outputMode {
	case "print":
		output.PrintInOrder(stepsCalendar.DateToSteps, printOrder)
	case "json":
		fmt.Printf("Exporting as JSON file\n")
	case "csv":
		fmt.Printf("Exporting as CSV file\n")
	default:
		fmt.Println("Invalid output format, select from 'print', 'json', 'csv'")
	}

	steps.AdditionalStepsStats(stepsCalendar)
	fmt.Printf("Finished\n")
}
