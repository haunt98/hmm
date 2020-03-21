package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
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

	names, err := getNames("dict/names.txt")
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
