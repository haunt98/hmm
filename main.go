package main

import (
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

	// pre seed
	rand.Seed(time.Now().UTC().UnixNano())

	names, err := getValues("dict/names.txt")
	if err != nil {
		log.Println("get names failed", err)
		return
	}

	gods, err := getValues("dict/gods.txt")
	if err != nil {
		log.Println("get gods failed", err)
	}

	r := router{
		names: names,
		gods:  gods,
	}

	http.HandleFunc("/", r.routeHome)
	http.HandleFunc("/name", r.routeName)
	http.HandleFunc("/project", r.routeProject)
	http.HandleFunc("/god", r.routeGod)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
