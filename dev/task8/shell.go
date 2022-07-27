package main

import (
	"bufio"
	"fmt"
	"github.com/shirou/gopsutil/load"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		command, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("cannot read command: %v\n", err)
			continue
		}
		command = command[:len(command)-1]
		commandParts := strings.Fields(command)
		switch commandParts[0] {
		case "exit":
			return
		case "cd":
			cd(commandParts)
		case "pwd":
			pwd()
		case "echo":
			echo(commandParts[1:])
		case "ps":
			ListProcesses()
		case "kill":
			pid, err := strconv.Atoi(commandParts[1])
			if err != nil {
				fmt.Println(err)
				continue
			}
			kill(pid)
		default:
			fmt.Printf("%s: command not found", commandParts[0])
		}
	}
}

func cd(args []string) {
	if len(args) < 2 {
		fmt.Println("Not enough arguments to change directory")
		return
	}
	err := os.Chdir(args[1])
	if err != nil {
		fmt.Println(err)
	}
}

func pwd() {
	path, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(path)
}

func echo(args []string) {
	for _, word := range args {
		fmt.Print(word)
		fmt.Print(" ")
	}
	fmt.Println()
}

func ListProcesses() {
	miscStat, _ := load.Misc()
	fmt.Printf("Running processes: %d\n", miscStat.ProcsRunning)
}

func kill(pid int) {
	process, err := os.FindProcess(pid)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = process.Kill()
	if err != nil {
		fmt.Println(err)
	}
}
