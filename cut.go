package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	// Step 1: Parse command line arguments
	var delimiter string
	var fields string
	flag.StringVar(&delimiter,"d", "\t", "Delimiter between fields")
	flag.StringVar(&fields, "f", "2", "Fields to print")
	flag.Parse()

	// Step 2: 	Check for the stdin or file inpute
	args := flag.Args()
	var input io.Reader

	if len(args) == 0 || args[0] == "-" {
		input = os.Stdin
	} else {
		file, err := os.Open(args[0])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error opening file: %s\n", err)
			os.Exit(1)
		}
		defer file.Close()
		input = file
	}

	//Step 3: Process inpute and print selected fields
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := scanner.Text()
		fieldsArray := strings.Split(line, delimiter)

		// Step 4: Parse fields and print
		selectedFields := parseFields(fields, fieldsArray)
		fmt.Println(strings.Join(selectedFields, "\t"))
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error reading input: %s\n", err)
		os.Exit(1)
	}
}


// parseFields parses the provided fields string and returns the corresponding field from the input line.

func parseFields(fields string, fieldsArray []string) []string {
	var selectedFields []string

	//Parse fields string to get individual field numbers 
	fieldNumbers := strings.FieldsFunc(fields, func(r rune) bool {
		return r == ',' || r == ' '
	})

	for _, field := range fieldNumbers {
		index := parseFieldIndex(field)
		if index >= 1 && index <=len(fieldsArray) {
			selectedFields = append(selectedFields, fieldsArray[index-1])
		}
	}

	return selectedFields
}

// parseFieldIndex parses the provided field string and returns the field index as an integer.

func parseFieldIndex(field string) int {
	index := 0
	fmt.Sscanf(field, "%d", &index)
	return index
}