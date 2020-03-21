package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func getNames(filename string) ([]string, error) {
	// open names.txt
	file, err := os.Open("names.txt")
	if err != nil {
		return nil, err
	}
	log.Println("open names.txt OK")

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
