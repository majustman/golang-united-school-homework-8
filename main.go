package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"io/fs"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type Arguments map[string]string

type user struct {
	ID    int    `json:"id"`
	Email string `json:"email"`
	Age   int    `json:"age"`
}

func Perform(args Arguments, writer io.Writer) error {

	oper, ok := args["operation"]
	if !ok {
		return errors.New("-operation flag has to be specified")
	}

	fileName, ok := args["fileName"]
	if !ok {
		return errors.New("-fileName flag has to be specified")
	}

	switch oper {
	case "add":
		err := addItem(fileName, args)
		if err != nil {
			return err
		}
	case "list":
		err := list(fileName, writer)
		if err != nil {
			return err
		}
	case "remove":
		id, ok := args["id"]
		if !ok {
			return errors.New("-id flag has to be specified")
		}
		err := removeUser(id, fileName)
		if err != nil {
			return err
		}
	default:
		return fmt.Errorf("operation %s not allowed", oper)
	}
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
		res                                  Arguments = make(map[string]string)
		operFlag, itemFlag, fileFlag, idFlag string
	)
	flag.StringVar(&operFlag, "operation", "", "value for the operation argument")
	flag.StringVar(&itemFlag, "item", "", "value for the item argument")
	flag.StringVar(&fileFlag, "fileName", "", "value for the fileName argument")
	flag.StringVar(&idFlag, "id", "", "value for the fileName argument")
	flag.Parse()
	if operFlag != "" {
		res["operation"] = operFlag
	}
	if itemFlag != "" {
		res["item"] = replaceChar(itemFlag)
	}
	if fileFlag != "" {
		res["fileName"] = fileFlag
	}
	if idFlag != "" {
		res["id"] = idFlag
	}
	fmt.Println(res)
	return res
}

// replaceChar replaces "«" or "»" on '"'
func replaceChar(input string) (output string) {
	for _, char := range input {
		if char == '«' || char == '»' {
			output += string('"')
		} else {
			output += string(char)
		}
	}
	return
}

func addItem(fileName string, args Arguments) error {

	list, err := readFile(fileName)
	if err != nil {
		return fmt.Errorf("adding operation error: %v", err)
	}

	u, err := createUserFromArg(args)
	if err != nil {
		return fmt.Errorf("adding operation error: %v", err)
	}
	list = append(list, u)

	err = writeToFile(fileName, list)
	if err != nil {
		return fmt.Errorf("adding operation error: %v", err)
	}

	return nil
}

func readFile(fileName string) ([]user, error) {
	var list []user
	content, err := ioutil.ReadFile(fileName)
	if errors.Is(err, fs.ErrNotExist) {
		return list, nil
	}
	if err != nil {
		return nil, fmt.Errorf("reading file error: %v", err)
	}
	json.Unmarshal(content, &list)
	return list, nil
}

func writeToFile(fileName string, list []user) error {
	content, err := json.Marshal(list)
	if err != nil {
		return fmt.Errorf("writing to file error: %v", err)
	}

	err = ioutil.WriteFile(fileName, content, 0644)
	if err != nil {
		return fmt.Errorf("writing to file error: %v", err)
	}
	return nil
}

func createUserFromArg(args Arguments) (user, error) {
	content, ok := args["item"]
	if !ok {
		return user{}, errors.New("-item flag has to be specified")
	}

	dec := json.NewDecoder(strings.NewReader(content))
	var u user
	err := dec.Decode(&u)
	if err != nil {
		return user{}, fmt.Errorf("creating user error: %v", err)
	}
	return u, nil
}

func list(fileName string, writer io.Writer) error {
	content, err := ioutil.ReadFile(fileName)
	if errors.Is(err, fs.ErrNotExist) {
		return nil
	}
	if err != nil {
		return fmt.Errorf("listing file error: %v", err)
	}
	_, err = writer.Write(content)
	if err != nil {
		return fmt.Errorf("listing file error: %v", err)
	}
	return nil
}

func removeUser(id string, fileName string) error {
	users, err := readFile(fileName)
	if err != nil {
		return fmt.Errorf("removing user error: %v", err)
	}
	n, err := strconv.Atoi(id)
	if err != nil {
		return fmt.Errorf("removing user error: %v", err)
	}
	var res []user
	for _, u := range users {
		if u.ID != n {
			res = append(res, u)
		}
	}
	err = writeToFile(fileName, res)
	if err != nil {
		return fmt.Errorf("removing user error: %v", err)
	}
	return nil
}
