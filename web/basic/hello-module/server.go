package main

import (
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	filesHandler := http.FileServer(http.Dir("./public"))
	mux.Handle("/static/", http.StripPrefix("/static/", filesHandler))
	http.ListenAndServe(":8080", mux)
	// http://127.0.0.1:8080/static/
}
