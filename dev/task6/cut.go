package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/infamax/l2/dev/task6/cut"
	"github.com/infamax/l2/dev/task6/flags"
	"log"
	"os"
)

func main() {
	sep := flag.Bool(flags.SeparatedFlag.Name, flags.SeparatedFlag.DefaultValue.(bool),
		flags.SeparatedFlag.Description)
	field := flag.Int(flags.FieldFlag.Name, flags.FieldFlag.DefaultValue.(int),
		flags.FieldFlag.Description)
	delimiter := flag.String(flags.DelimiterFlag.Name,
		flags.DelimiterFlag.DefaultValue.(string), flags.DelimiterFlag.Description)
	flag.Parse()
	args := flag.Args()

	if *field == 0 {
		log.Fatal("cut: you must specify a list of bytes, characters," +
			" or fields. Try 'cut --help' for more information.")
	}

	res := make([]string, 0)

	for _, arg := range args {
		f, err := os.Open(arg)
		if err != nil {
			fmt.Printf("cut: %s: No such file or directory", arg)
			continue
		}
		buf := bufio.NewScanner(f)
		inputData := make([]string, 0)
		for buf.Scan() {
			inputData = append(inputData, buf.Text())
		}
		res = append(res, inputData...)
	}

	if *sep {
		res = cut.DeleteNotSeparatedStrings(res, *delimiter)
	}

	res = cut.DefaultCut(res, *field, *delimiter)

	for _, str := range res {
		fmt.Println(str)
	}
}
