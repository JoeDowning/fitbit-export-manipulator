package output

import (
	"fmt"
	"sort"
)

func PrintInOrder(m map[string]int, order string) {
	var dates []string
	for date := range m {
		dates = append(dates, date)
	}

	switch order {
	case "newFirst":
		sort.Sort(sort.Reverse(sort.StringSlice(dates)))
	case "oldFirst":
		sort.Sort(sort.StringSlice(dates))
	default:
		sort.Sort(sort.StringSlice(dates))
	}

	for _, date := range dates {
		fmt.Printf("%s: %d\n", date, m[date])
	}
}
