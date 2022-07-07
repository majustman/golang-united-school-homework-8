package main

import (
	"io"
	"os"
)

type Arguments map[string]string

func Perform(args Arguments, writer io.Writer) error {
	return nil
}

func main() {
	err := Perform(parseArgs(), os.Stdout)
	if err != nil {
		panic(err)
	}
}

func parseArgs() Arguments {
	var res Arguments = make(map[string]string)
	args := os.Args[1:]
	for i, arg := range args {
		switch arg {
		case "-operation":
			if i != len(args)-1 && args[i+1][0] != '-' {
				res[arg[1:]] = args[i+1]
			}
		case "-item":
			if i != len(args)-1 && args[i+1][0] != '-' {
				res[arg[1:]] = args[i+1]
			}
		case "-fileName":
			if i != len(args)-1 && args[i+1][0] != '-' {
				res[arg[1:]] = args[i+1]
			}
		}
	}
	return res
}
