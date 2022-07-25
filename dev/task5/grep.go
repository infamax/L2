package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/infamax/l2/dev/task5/flags"
	"github.com/infamax/l2/dev/task5/grep"
	"log"
	"os"
)

func printSlice(words []string) {
	for _, word := range words {
		fmt.Println(word)
	}
}

func main() {
	after := flag.Int(flags.AfterFlag.Name, flags.AfterFlag.DefaultValue.(int), flags.AfterFlag.Description)
	before := flag.Int(flags.BeforeFlag.Name, flags.BeforeFlag.DefaultValue.(int), flags.BeforeFlag.Description)
	context := flag.Int(flags.ContextFlag.Name, flags.ContextFlag.DefaultValue.(int), flags.ContextFlag.Description)
	count := flag.Bool(flags.CountFlag.Name, flags.CountFlag.DefaultValue.(bool), flags.CountFlag.Description)
	ignoreCase := flag.Bool(flags.IgnoreCaseFlag.Name, flags.IgnoreCaseFlag.DefaultValue.(bool), flags.IgnoreCaseFlag.Description)
	invert := flag.Bool(flags.InvertFlag.Name, flags.InvertFlag.DefaultValue.(bool), flags.InvertFlag.Description)
	fixed := flag.Bool(flags.FixedFlag.Name, flags.FixedFlag.DefaultValue.(bool), flags.FixedFlag.Description)
	lineNum := flag.Bool(flags.LineNumFlag.Name, flags.LineNumFlag.DefaultValue.(bool), flags.LineNumFlag.Description)
	flag.Parse()

	if flag.NArg() < 2 {
		log.Fatal("required pattern(regular expression) and name file!")
	}

	args := flag.Args()
	pattern := args[0]
	filename := args[1]

	var file *os.File
	var err error
	if filename != "-" {
		file, err = os.Open(filename)
		if err != nil {
			log.Fatalf("grep: %s: No such file or directory", filename)
		}
	}

	scanner := bufio.NewScanner(file)
	input := make([]string, 0)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	if *after == 0 && *before == 0 &&
		*context == 0 && !*count && !*ignoreCase &&
		!*invert && !*fixed && !*lineNum {
		res, err := grep.StandardGrep(input, pattern)
		if err != nil {
			log.Fatalf("invalid regular expression: %v", err)
		}
		printSlice(res)
		return
	}

	if *after == 0 && *before == 0 &&
		*context == 0 && *count && !*ignoreCase &&
		!*fixed && !*lineNum {
		c, err := grep.CountGrep(input, pattern, *invert)
		if err != nil {
			log.Fatalf("invalid regular expression: %v", err)
		}
		fmt.Println(c)
		return
	}

	if *after == 0 && *before == 0 &&
		*context == 0 && !*count && *lineNum {
		res, err := grep.LineGrep(input, pattern, *invert, *ignoreCase)
		if err != nil {
			log.Fatalf("invalid regular expression: %v", err)
		}
		printSlice(res)
		return
	}

	if *after != 0 || *before != 0 || *context != 0 {
		res, err := grep.ContextGrep(input, pattern, *after, *before, *context, *invert, *ignoreCase)
		if err != nil {
			log.Fatalf("invalid regular expression: %v", err)
		}
		printSlice(res)
		return
	}
}
