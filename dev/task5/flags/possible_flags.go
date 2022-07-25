package flags

type Flag struct {
	Name         string
	DefaultValue interface{}
	Description  string
}

var (
	AfterFlag = Flag{
		Name:         "A",
		DefaultValue: 0,
		Description: " Print NUM lines of trailing context after matching lines.  " +
			"Places a line containing a group separator (--) between contiguous groups of matches. " +
			" With the -o or --only-matching option, this has no effect and a warning is given.",
	}
	BeforeFlag = Flag{
		Name:         "B",
		DefaultValue: 0,
		Description: " Print NUM lines of leading context before matching lines. " +
			"Places a line containing a group separator (--) between contiguous groups of matches. " +
			"With the -o or --only-matching option, this has no effect and a warning is given",
	}
	ContextFlag = Flag{
		Name:         "C",
		DefaultValue: 0,
		Description: "Print NUM lines of output context. " +
			"Places a line containing a group separator (--) between contiguous groups of matches." +
			" With the -o or --only-matching option, this  has  no  effect  and a warning is given.",
	}
	CountFlag = Flag{
		Name:         "c",
		DefaultValue: false,
		Description: "Suppress normal output; instead print a count of matching lines for each input file. " +
			"With the -v, --invert-match option (see below), count non-matching lines",
	}
	IgnoreCaseFlag = Flag{
		Name:         "i",
		DefaultValue: false,
		Description: "Ignore case distinctions in patterns and input data, " +
			"so that characters that differ only in case match each other.",
	}
	InvertFlag = Flag{
		Name:         "v",
		DefaultValue: false,
		Description:  " Invert the sense of matching, to select non-matching lines.",
	}
	FixedFlag = Flag{
		Name:         "F",
		DefaultValue: false,
		Description:  "Interpret PATTERNS as fixed strings, not regular expressions.",
	}
	LineNumFlag = Flag{
		Name:         "n",
		DefaultValue: false,
		Description:  "Prefix each line of output with the 1-based line number within its input file.",
	}
)
