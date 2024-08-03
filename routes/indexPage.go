package routes

import (
	"fmt"
	"html/template"
	"net/http"
	"roman-to-integer/lib"
	"strconv"
)

type Content struct {
	Title string
	Body  string
}

func HelloPage(w http.ResponseWriter, r *http.Request) {
	// tmpl, err := template.ParseFiles(path.Join("views", "helloworld.html"))
	data := Content{
		Title: "Hello world",
		Body:  "isi content halo dunia",
	}
	tmpl, err := template.ParseFiles("views/helloworld.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, data)
}

func ConverterPage(w http.ResponseWriter, r *http.Request) {
	data := Content{
		Title: "page from converter",
		Body:  "this is roman converter page",
	}
	if r.Method == "POST" {
		fmt.Println("haiiii from post")
		text := r.FormValue("string")
		convertedRoman := lib.RomanToInt(text)
		str_roman := strconv.Itoa(convertedRoman)
		data.Body = str_roman
		tmpl, err := template.ParseFiles("views/helloworld.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		tmpl.Execute(w, data)
	} else if r.Method == "GET" {
		tmpl, err := template.ParseFiles("views/roman.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		tmpl.Execute(w, "")

	}
}
