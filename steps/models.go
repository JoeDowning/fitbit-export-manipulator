package steps

type StepsCalendar struct {
	DateToSteps map[string]int `json:"dateToSteps"`
}

type extremeStat struct {
	date  string
	steps int
}
