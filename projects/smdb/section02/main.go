/*
  Templates are populated by data passed from the application to *Template.ExecuteTemplate.
  It is useful and efficient and manageable to create and pass composite data structures that contain
  all data used in a template. text/template functionality in the template files themselves makes
  it possible to parse and act upon that provided data to populate the template files.
*/

package main

import (
	"log"
	"os"
	"sort"
	"html/template"
	"strings"
	"fmt"
)

type Person struct {
	Fname string
	Lname string
}

type People []Person

type Film struct {
	Title    string
	Year     int
	Director Person
	Writers  People
	Actors   People
}

type Films []Film

func (p People) Len() int {
	return len(p)
}

func (p People) Less(i, j int) bool {
	return p[i].Lname < p[j].Lname
}

func (p People) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p People) sortNames() People {
	sort.Sort(p)
	return p
}

func (f Film) toPageName() string {
	return strings.Replace(strings.ToLower(f.Title), " ", "_", -1)
}

var tpls *template.Template

var fm = template.FuncMap{
	"sn": People.sortNames,
	"tpn": Film.toPageName,
}

func init() {
	tpls = template.Must(template.New("").Funcs(fm).ParseGlob("templates/*.gohtml"))
}

func main() {
	p1 := Person{"Ridley", "Scott"}
	p2 := Person{"Scott", "Derrickson"}
	p3 := Person{"Andy", "Weir"}
	p4 := Person{"Drew", "Goddard"}
	p5 := Person{"Jon", "Spaihts"}
	p6 := Person{"Matt", "Damon"}
	p7 := Person{"Donald", "Glover"}
	p8 := Person{"Kristen", "Wiig"}
	p9 := Person{"Jeff", "Daniels"}
	p10 := Person{"Chiwetel", "Ejiofor"}
	p11 := Person{"Jessica", "Chastain"}
	p12 := Person{"Michael", "Peña"}
	p13 := Person{"Kate", "Mara"}
	p14 := Person{"Benedict", "Wong"}
	p15 := Person{"Sean", "Bean"}
	p16 := Person{"Mackenzie", "Davis"}
	p17 := Person{"Noomi", "Rapace"}
	p18 := Person{"Michael", "Fassbender"}
	p19 := Person{"Charlize", "Theron"}
	p20 := Person{"Idris", "Elba"}
	p21 := Person{"Guy", "Pearce"}
	p22 := Person{"Benedict", "Cumberbatch"}
	p23 := Person{"Rachel", "MacAdams"}
	p24 := Person{"Mads", "Mikkelsen"}
	p25 := Person{"Tilda", "Swinton"}
	p26 := Person{"Benjamin", "Bratt"}
	p27 := Person{"Damon", "Lindelof"}


	f1 := Film{
		"The Martian",
		2015,
		p1,
		People{p3, p4},
		People{p6, p7, p8, p9, p10, p11, p12, p13, p14, p15, p16},
	}

	f2 := Film{
		"Doctor Strange",
		2016,
		p2,
		People{p2, p5},
		People{p10, p14, p22, p23, p24, p25, p26},
	}

	f3 := Film{
		"Prometheus",
		2012,
		p1,
		People{p5, p27},
		People{p14, p17, p18, p19, p20, p21},
	}

	films := Films{f1, f2, f3}

	data := struct{
		Title string
		Films
	}{
		"Scott's IMDb",
		films,
	}

	filename := "index.html"
	file, err := os.Create(filename)
	if err != nil {
		file.Close()
		log.Fatalf("error creating %s: %s", filename, err)
	}

	err = tpls.ExecuteTemplate(file, "index.gohtml", data)
	if err != nil {
		file.Close()
		log.Fatalf("error executing template: %s", err)
	}
	file.Close()

	for _, film := range films {
		filename := fmt.Sprintf("%s.html", Film.toPageName(film))

		file, err = os.Create(filename)
		if err != nil {
			file.Close()
			log.Fatalf("error creating %s: %s", filename, err)
		}

		err = tpls.ExecuteTemplate(file, "film.gohtml", film)
		if err != nil {
			file.Close()
			log.Fatalf("error executing template: %s", err)
		}
		file.Close()
	}
}
