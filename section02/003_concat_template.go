/*
  Before using text/template or html/template, think like a programmer:
  how do I get data into a body of text?

  Generate an HTML file from command-line args and template text.
*/

package main

import (
	"io"
	"log"
	"os"
	"strings"
)

const (
	indexTemplate = `
	<html>
	  <head>
	    <metadata />
	  </head>
	  <body>
	    <p>Hello, {{name}}.</p>
	    <p>Welcome to the abyss.</p>
	    <p>Eternity awaits.</p>
	    <p>Oh, by the way, {{name}}, {{color}} is my favorite color, too!</p>
	  </body>
	</html>`
)

type data struct {
	name  string
	color string
}

var filename string

func main() {
	args := os.Args
	if len(args) < 4 {
		log.Fatalln("Missing arguments, bro. I need a name, a color, and a filename.")
	}

	userData := data{
		strings.Title(args[1]),
		args[2],
	}
	filename = args[3]

	indexPopulated := strings.Replace(indexTemplate, "{{name}}", userData.name, -1)
	indexPopulated = strings.Replace(indexPopulated, "{{color}}", userData.color, -1)

	file, err := os.Create(filename)
	if err != nil {
		log.Fatalln("Error opening file to create.")
	}
	defer file.Close()

	io.Copy(file, strings.NewReader(indexPopulated))
}
