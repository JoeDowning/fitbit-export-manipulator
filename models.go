package main

type Entry struct {
	DateTime string `json:"dateTime"`
	Value    string `json:"value"`
}

type Entries []Entry

type FilesMap map[string]Entries

type StepsCalendar struct {
	DateToSteps map[string]int `json:"dateToSteps"`
}
