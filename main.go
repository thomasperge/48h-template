package main

import (
	"fmt"
	"net/http"
	"text/template"
)

func Home(rw http.ResponseWriter, r *http.Request) {
	template, _ := template.ParseFiles("index.html")
	fmt.Println(rw, r)
	fmt.Println("TEST")

	template.Execute(rw, r)
}

func main() {
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		Home(rw, r)
	})

	fmt.Println("http://localhost:8080/")
	http.ListenAndServe(":8080", nil)
}
