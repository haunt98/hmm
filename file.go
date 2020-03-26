package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func getValues(filename string) ([]string, error) {
	// open names.txt
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	log.Printf("open %s OK\n", filename)

	// scan each line then append to names
	names := make([]string, 0, namesSize)

	fileScanner := bufio.NewScanner(file)
	for fileScanner.Scan() {
		name := fileScanner.Text()
		name = strings.TrimSpace(name)
		name = strings.ReplaceAll(name, " ", "")
		name = strings.ReplaceAll(name, "\t", "")
		names = append(names, name)
	}
	if err := fileScanner.Err(); err != nil {
		return nil, err
	}

	return names, nil
}
