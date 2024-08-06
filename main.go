package main

import (
	"fmt"
	"net/http"
	"roman-to-integer/routes"
)

func main() {
	// text := "MCMXLV"
	// number := lib.RomanToInt(text)
	// fmt.Printf("the %s when converted to int is %d", text, number)
	http.HandleFunc("/", routes.HomePage)
	http.HandleFunc("/roman", routes.ConverterPage)
	http.HandleFunc("/roman/process", routes.RomanAjax)
	addr := ":8080"
	fmt.Println("http server served on ", addr)
	http.ListenAndServe(addr, nil)

}
