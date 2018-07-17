package main

import (
	"html/template"
	"log"
	"net/http"
)

type name string

// method on a type that invokes the http.Handler interface
func (n name) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}

	// could instead call r.PostForm here and only get POSTed key-values
	tpl.ExecuteTemplate(w, "index.gohtml", r.Form)
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}

func main() {
	var n name
	http.ListenAndServe(":8080", n)
}
