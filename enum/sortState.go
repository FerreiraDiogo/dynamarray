package enum

type SortState int

const (
	SortAscending SortState = iota
	SortDescending
	Unknown
)

var stateName = map[SortState]string{
	SortAscending:  "Ascending",
	SortDescending: "Descending",
	Unknown:        "Unknown",
}

func (s SortState) String() string {
	if name, ok := stateName[s]; ok {
		return name
	}
	return "Unknown"
}
