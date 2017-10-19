package main

import (
	"log"
	"encoding/csv"
	"text/template"
	"io"
	"os"
)

type Field struct {
	Index int
	Name string
}

type Fields []Field

type Record struct {
	Open string
	Close string
}

type Records []Record

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	file, err := os.Open("data.csv")
	if err != nil {
		log.Fatalln("Unable to open data file", err)
	}

	csvData := csv.NewReader(io.Reader(file))

	header, err := csvData.Read()
	if err != nil {
		log.Fatalln("Error reading fields", err)
	}

	var fields Fields
	for idx, name := range header {
		if name == "Open" || name == "Close" {
			fields = append(fields, Field {idx, name})
		}
	}

	allRows, err := csvData.ReadAll()
	if err != nil {
		log.Fatalln("Error reading all rows", err)
	}

	var records Records
	for _, row := range allRows {
		records = append(records, Record{row[fields[0].Index], row[fields[1].Index]})
	}

	displayData := struct {
		Fields
		Records
	}{
		fields,
		records,
	}
	err = tpl.Execute(os.Stdout, displayData)
	if err != nil {
		log.Fatalln(err)
	}
}