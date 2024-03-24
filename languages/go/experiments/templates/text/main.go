package main

import (
	"os"
	"strings"
	"text/template"
)

type Pet struct {
	Name   string
	Sex    string
	Intact bool
	Age    string
	Breed  string
}

func main() {
	dogs := []Pet{
		{
			Name:   "<script>alert(\"Gotcha!\");</script>Jujube",
			Sex:    "Female",
			Intact: false,
			Age:    "10 months",
			Breed:  "German Shepherd/Pitbull",
		},
		{
			Name:   "Zephyr",
			Sex:    "Male",
			Intact: true,
			Age:    "13 years, 3 months",
			Breed:  "German Shepherd/Border Collie",
		},
		{
			Name:   "Bruce Wayne",
			Sex:    "Male",
			Intact: false,
			Age:    "3 years, 8 months",
			Breed:  "Chihuahua",
		},
	}

	// all dogs template
	var tmplFile = "pets.tmpl"
	tmpl, err := template.New(tmplFile).ParseFiles(tmplFile)
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(os.Stdout, dogs)
	if err != nil {
		panic(err)
	}

	// now for the lastPet template
	println("\n---------- Last Pet")

	functionMap := template.FuncMap{
		"dec":     func(i int) int { return i - 1 },
		"replace": strings.ReplaceAll,
	}

	tmplFile = "lastPet.tmpl"
	tmpl, err = template.New(tmplFile).Funcs(functionMap).ParseFiles(tmplFile)
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(os.Stdout, dogs)
	if err != nil {
		panic(err)
	}

	// now write HTML using the text/template structure to show it's unsafe
	tmplFile = "petsHtml.tmpl"
	tmpl, err = template.New(tmplFile).Funcs(functionMap).ParseFiles(tmplFile)
	if err != nil {
		panic(err)
	}
	var f *os.File
	f, err = os.Create("pets.html")
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(f, dogs)
	if err != nil {
		panic(err)
	}
	err = f.Close()
	if err != nil {
		panic(err)
	}
}
