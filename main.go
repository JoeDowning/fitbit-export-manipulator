package main

import "fmt"

const path = "path/"

func main() {
	filesMap, err := extractJSONFiles(path)
	if err != nil {
		fmt.Println(err)
		return
	}

	stepsCalendar, err := totalStepsByDate(filesMap)
	if err != nil {
		fmt.Println(err)
		return
	}

	for date, steps := range stepsCalendar.DateToSteps {
		fmt.Printf("Date: %s, Steps: %d\n", date, steps)
	}

	printInOrder(stepsCalendar.DateToSteps)
}
