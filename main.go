package main

import (
	"fmt"
	"net/http"
)

func Home(rw http.ResponseWriter, r *http.Request) {
	fmt.Println(rw, r)
	fmt.Println("TEST")
}

func main() {
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		Home(rw, r)
	})

    fmt.Println("http://localhost:8080/")
}
