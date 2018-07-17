package main

import (
	"fmt"
	"net/http"
)

type macncheese struct {
	macaroni string
	cheese   string
}

func (m macncheese) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Mac 'n' cheese is, just, %s and %s, right?", m.macaroni, m.cheese)
}

func main() {
	x := macncheese{
		"elbows",
		"cheddar",
	}
	http.ListenAndServe(":8080", x)
}
