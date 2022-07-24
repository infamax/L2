package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/infamax/l2/dev/task3/internal/flags"
	"github.com/infamax/l2/dev/task3/internal/sorts"
	"os"
)

func main() {
	flag.BoolVar(&flags.BoolFlags[flags.RevFlag],
		flags.ReverseFlag.Name, flags.ReverseFlag.DefaultValue.(bool), flags.ReverseFlag.Description)
	flag.BoolVar(&flags.BoolFlags[flags.NumFlag], flags.NumericFlag.Name, flags.NumericFlag.DefaultValue.(bool),
		flags.NumericFlag.Description)
	flag.BoolVar(&flags.BoolFlags[flags.UniqFlag], flags.UniqueFlag.Name, flags.UniqueFlag.DefaultValue.(bool),
		flags.UniqueFlag.Description)
	flag.BoolVar(&flags.BoolFlags[flags.CheckSortedFlag], flags.CheckedSortedFlag.Name,
		flags.CheckedSortedFlag.DefaultValue.(bool), flags.CheckedSortedFlag.Description)
	flag.BoolVar(&flags.BoolFlags[flags.MonthSortFlag], flags.MonthFlag.Name,
		flags.MonthFlag.DefaultValue.(bool), flags.MonthFlag.Description)
	flag.BoolVar(&flags.BoolFlags[flags.HumanSortFlag], flags.HumanReadableFlag.Name,
		flags.HumanReadableFlag.DefaultValue.(bool), flags.HumanReadableFlag.Description)
	flag.BoolVar(&flags.BoolFlags[flags.IgnoreFlag], flags.IgnoreLeadingBlanksFlag.Name,
		flags.IgnoreLeadingBlanksFlag.DefaultValue.(bool), flags.IgnoreLeadingBlanksFlag.Description)
	flag.IntVar(&flags.NumFlags[flags.ColFlag], flags.ColumnFlag.Name, flags.ColumnFlag.DefaultValue.(int),
		flags.ColumnFlag.Description)
	flag.Parse()

	if flag.Arg(0) == "" {
		fmt.Println("Enter file path after flags!")
		return
	}

	outputStream := os.Stdout
	defer outputStream.Close()
	var err error
	if flag.Arg(1) != "" {
		//fileOutput = true
		fmt.Println("file created!")
		outputStream, err = os.Create(flag.Arg(1))
		if err != nil {
			fmt.Printf("Cannot create file with name: %s\n", flag.Arg(1))
		}
	}

	f, err := os.Open(flag.Arg(0))
	defer f.Close()

	if err != nil {
		fmt.Printf("Cannot open file with name: %s", flag.Arg(0))
	}

	buf := bufio.NewScanner(f)
	strs := make([]string, 0)

	for buf.Scan() {
		strs = append(strs, buf.Text())
	}

	var res []string
	fl := false
	if !flags.BoolFlags[flags.CheckSortedFlag] {
		res = sorts.SortWithoutKeys(strs)
	}

	for i := 0; i < len(flags.BoolFlags); i++ {
		if flags.BoolFlags[i] {
			switch i {
			case flags.RevFlag:
				fmt.Println("RevFlag!")
				res, fl = sorts.SortReverseOrder(res)
			case flags.NumFlag:
				fmt.Println("NumFlag!")
				res = sorts.SortByNumericValue(res)
			case flags.UniqFlag:
				fmt.Println("UniqFlag!")
				res = sorts.SortUnique(res)
			case flags.CheckSortedFlag:
				fmt.Println("CheckFlag!")
				fl = sorts.CheckSortedSlice(res)
			case flags.MonthSortFlag:
				fmt.Println("MonthFlag!")
				res = sorts.SortMonth(res)
			case flags.HumanSortFlag:
				fmt.Println("HumanReadableFlag")
				res = sorts.SortHumanReadable(res)
			case flags.IgnoreFlag:
				fmt.Println("IgnoreLeadingBlanksFlag")
				res = sorts.SortIgnoreLeadingBlanks(res)
			}
		}
	}

	for i := 0; i < len(flags.NumFlags); i++ {
		if flags.NumFlags[i] != 1 {
			switch i {
			case flags.ColFlag:
				fmt.Println("ColFlag!")
				res, err = sorts.SortByColumn(res, flags.NumFlags[flags.ColFlag])
				if err != nil {
					fmt.Println(err)
					return
				}
			}
		}
	}

	if flags.BoolFlags[flags.CheckSortedFlag] {
		if fl {
			fmt.Println("Array sorted")
		} else {
			fmt.Println("Array not sorted")
		}
		return
	}

	for _, val := range res {
		_, err := outputStream.WriteString(val + "\n")
		if err != nil {
			return
		}
	}
}
