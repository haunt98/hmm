package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strings"
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

	names, err := getNames("names.txt")
	if err != nil {
		log.Println("get names failed", err)
		return
	}

	r := router{
		names: names,
	}

	http.HandleFunc("/", r.routeHome)
	http.HandleFunc("/name", r.routeName)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}

type router struct {
	names []string
}

func (r *router) routeHome(w http.ResponseWriter, req *http.Request) {
	if _, err := fmt.Fprintf(w, "GET /name\n"); err != nil {
		log.Println(err)
	}
}

func (r *router) routeName(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		if _, err := fmt.Fprintf(w, "wrong method %s", req.Method); err != nil {
			log.Println(err)
			return
		}
	}

	// rand name
	rand.Seed(time.Now().UTC().UnixNano())
	randIndex := rand.Intn(len(r.names))
	postFix := rand.Intn(postfixSize)

	// send name
	if _, err := fmt.Fprintf(w, "%s%d", r.names[randIndex], postFix); err != nil {
		log.Println(err)
		return
	}
}

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
