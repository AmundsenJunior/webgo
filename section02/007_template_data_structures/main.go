/*
  Templates are populated by data passed from the application to *Template.ExecuteTemplate.
  It is useful and efficient and manageable to create and pass composite data structures that contain
  all data used in a template. text/template functionality in the template files themselves makes
  it possible to parse and act upon that provided data to populate the template files.
 */

package main

import (
	"text/template"
	"os"
	"log"
)

type Person struct {
	Fname string
	Lname string
}

type Actor struct {
	Person
}

type Director struct {
	Person
}

type Writer struct {
	Person
}

type Film struct {
	Name string
	Year int
	Actors []Actor
	Director
	Writers []Writer
}

var tpls *template.Template

func init() {
	tpls = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {
	a1 := Actor{Person{"Matt","Damon"}}
	a2 := Actor{Person{"Donald","Glover"}}
	a3 := Actor{Person{"Kristen","Wiig"}}
	a4 := Actor{Person{"Jeff","Daniels"}}
	a5 := Actor{Person{"Chiwetel","Ejiofor"}}
	a6 := Actor{Person{"Jessica","Chastain"}	}

	d1 := Director{Person{"Ridley","Scott"}}

	w1 := Writer{Person{"Andy","Weir"}}
	w2 := Writer{Person{"Drew","Goddard"}}

	f1 := Film {
		"The Martian",
		2015,
		[]Actor{a1, a2, a3, a4, a5, a6},
		d1,
		[]Writer{w1, w2},
	}

	err := tpls.ExecuteTemplate(os.Stdout, "film.gohtml", f1)
	if err != nil {
		log.Fatalln(err)
	}
}
