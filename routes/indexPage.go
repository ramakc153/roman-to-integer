package routes

import (
	"encoding/json"
	"html/template"
	"net/http"
	"roman-to-integer/lib"
	"strconv"
)

type Content struct {
	Title string
	Body  string
}

type Response struct {
	Message string `json:"message"`
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
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Failed to parse form data", http.StatusBadRequest)
			return
		}
		text := r.FormValue("string")
		convertedRoman := lib.RomanToInt(text)
		str_roman := strconv.Itoa(convertedRoman)
		data.Body = str_roman
		tmpl, err := template.ParseFiles("views/roman.html")
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

func RomanAjax(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Unable to parse form", http.StatusBadRequest)
			return
		}
		text := r.FormValue("string")
		convertedRoman := lib.RomanToInt(text)
		str_roman := strconv.Itoa(convertedRoman)
		response := Response{Message: str_roman}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}
