package main

import (
	"flag"
	"fmt"
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
	var (
		res                          Arguments = make(map[string]string)
		operFlag, itemFlag, fileFlag string
	)
	//args := os.Args[1:]
	// for i, arg := range args {
	// 	switch arg {
	// 	case "-operation":
	// 		if i != len(args)-1 && args[i+1][0] != '-' {
	// 			res[arg[1:]] = args[i+1]
	// 		}
	// 	case "-item":
	// 		if i != len(args)-1 && args[i+1][0] != '-' {
	// 			res[arg[1:]] = args[i+1]
	// 		}
	// 	case "-fileName":
	// 		if i != len(args)-1 && args[i+1][0] != '-' {
	// 			res[arg[1:]] = args[i+1]
	// 		}
	// 	}
	// }
	flag.StringVar(&operFlag, "operation", "", "value for the operation argument")
	flag.StringVar(&itemFlag, "item", "", "value for the item argument")
	flag.StringVar(&fileFlag, "fileName", "", "value for the fileName argument")
	flag.Parse()
	res["operation"] = operFlag
	res["item"] = itemFlag
	res["fileName"] = fileFlag
	fmt.Println(res)

	return res
}
