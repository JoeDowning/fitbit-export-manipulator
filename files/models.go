package files

type Entry struct {
	DateTime string `json:"dateTime"`
	Value    string `json:"value"`
}

type Entries []Entry

type FilesMap map[string]Entries
