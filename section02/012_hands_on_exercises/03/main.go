package main

import (
	"log"
	"os"
	"text/template"
)

type Hotel struct {
	Name    string
	Address string
	City    string
	Zipcode string
}

type Hotels []Hotel

type Region struct {
	Name   string
	Hotels
}

type Regions []Region

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	h1 := Hotel{
		"Bleary",
		"23 Way",
		"LA",
		"0043",
	}
	h2 := Hotel{
		"Sizzle",
		"028v sxxet",
		"San foerkg",
		"1111111",
	}
	h3 := Hotel{
		"Captacha",
		"buuttss",
		"zedesert",
		"11032",
	}
	h4 := Hotel{
		"Whatsit",
		"333 half satan ln",
		"caofklwudoa",
		"000001",
	}
	h5 := Hotel{
		"Icayntevenriyghtnyow",
		"0003 cmn",
		"Los Angeles",
		"68584",
	}

	r1 := Region{
		"Southern",
		[]Hotel{h1, h2},
	}
	r2 := Region{
		"Central",
		[]Hotel{h3},
	}
	r3 := Region{
		"Northern",
		[]Hotel{h4, h5},
	}

	regions := Regions{r1, r2, r3}

	err := tpl.Execute(os.Stdout, regions)
	if err != nil {
		log.Fatalln(err)
	}

}
