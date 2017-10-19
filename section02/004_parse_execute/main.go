/*
  Parsing template files with text/template, usually as a Glob of all files, only once at the start of program execution.
 */

package main

import (
	"text/template"
	"os"
	"log"
)

var tpls *template.Template

func init() {
	tpls = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {
	err := tpls.ExecuteTemplate(os.Stdout, "tpl01.gohtml", nil)
	if err != nil {
		log.Fatal(err)
	}

	for _, tpl := range tpls.Templates() {
		err := tpl.Execute(os.Stdout, nil)
		if err != nil {
			log.Fatalln(err)
		}
	}
}