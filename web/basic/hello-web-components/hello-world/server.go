package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	filesHandler := http.FileServer(http.Dir("./public"))
	mux.Handle("/static/", http.StripPrefix("/static/", filesHandler))
	fmt.Println("http://127.0.0.1:8080/static/")
	http.ListenAndServe(":8080", mux)
}
