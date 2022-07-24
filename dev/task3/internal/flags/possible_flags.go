package flags

const (
	RevFlag = iota
	NumFlag
	UniqFlag
	CheckSortedFlag
	MonthSortFlag
	HumanSortFlag
	IgnoreFlag
)

const (
	ColFlag = iota
)

var (
	BoolFlags [7]bool
	NumFlags  = [1]int{0}
)

var (
	ReverseFlag = Flag{
		Name:         "r",
		DefaultValue: false,
		Description:  "reverse the result of comparisons",
	}
	NumericFlag = Flag{
		Name:         "n",
		DefaultValue: false,
		Description:  "compare according to string numerical value",
	}
	UniqueFlag = Flag{
		Name:         "u",
		DefaultValue: false,
		Description:  "with -c, check for strict ordering; without -c, output only the first of an equal run",
	}
	ColumnFlag = Flag{
		Name:         "k",
		DefaultValue: 1,
		Description:  "sort via a key; KEYDEF gives location and type",
	}
	CheckedSortedFlag = Flag{
		Name:         "c",
		DefaultValue: false,
		Description:  "check for sorted input; do not sort",
	}
	MonthFlag = Flag{
		Name:         "M",
		DefaultValue: false,
		Description:  " compare (unknown) < 'JAN' < ... < 'DEC'",
	}
	HumanReadableFlag = Flag{
		Name:         "h",
		DefaultValue: false,
		Description:  " compare human readable numbers (e.g., 2K 1G)",
	}
	IgnoreLeadingBlanksFlag = Flag{
		Name:         "b",
		DefaultValue: false,
		Description:  "ignore leading blanks",
	}
)
