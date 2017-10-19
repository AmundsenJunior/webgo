package main

import (
	"fmt"
	"net/http"
)

type macaroni string

func (m macaroni) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Macaroni is, just, the %s, right?", m)
}

func main() {
	var x macaroni
	x = "best"
	http.ListenAndServe(":8080", x)
}
