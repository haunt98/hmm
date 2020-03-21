package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"
)

const (
	namesSize   = 4096
	postfixSize = 10000
)

func main() {
	// PORT
	port := os.Getenv("PORT")
	if port == "" {
		log.Println("PORT is empty")
		return
	}
	log.Printf("PORT is %s\n", port)

	// open names.txt
	file, err := os.Open("names.txt")
	if err != nil {
		log.Println("failed to open file", err)
		return
	}
	log.Println("open names.txt OK")

	// scan each line then append to names
	names := make([]string, 0, namesSize)

	fileScanner := bufio.NewScanner(file)
	for fileScanner.Scan() {
		names = append(names, fileScanner.Text())
	}
	if err := fileScanner.Err(); err != nil {
		log.Println("failed to scan file", err)
		return
	}

	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		if _, err := fmt.Fprintf(w, "GET /name\n"); err != nil {
			log.Println(err)
		}
	})
	http.HandleFunc("/name", func(w http.ResponseWriter, req *http.Request) {
		if req.Method != http.MethodGet {
			if _, err := fmt.Fprintf(w, "wrong method %s", req.Method); err != nil {
				log.Println(err)
				return
			}
		}

		// rand name
		rand.Seed(time.Now().UTC().UnixNano())
		randIndex := rand.Intn(len(names))
		postFix := rand.Intn(postfixSize)

		if _, err := fmt.Fprintf(w, "%s%d", names[randIndex], postFix); err != nil {
			log.Println(err)
			return
		}
	})

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
