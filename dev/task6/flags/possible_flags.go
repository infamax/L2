package flags

type Flag struct {
	Name         string
	DefaultValue interface{}
	Description  string
}

var (
	FieldFlag = Flag{
		Name:         "f",
		DefaultValue: 0,
		Description: "select  only these fields;  " +
			"also print any line that contains no delimiter character, unless the -s option is specified",
	}
	DelimiterFlag = Flag{
		Name:         "d",
		DefaultValue: "    ",
		Description:  "use DELIM instead of TAB for field delimiter",
	}
	SeparatedFlag = Flag{
		Name:         "s",
		DefaultValue: false,
		Description:  "do not print lines not containing delimiters",
	}
)
