package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob(("templates/*.html")))
}

func main() {
	word := "go run run"
	f, err := os.Create("go.html")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	err = tpl.ExecuteTemplate(f, "index.html", word)
	if err != nil {
		log.Fatal(err)
	}

}
