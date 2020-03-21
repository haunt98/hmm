package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"
)

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
