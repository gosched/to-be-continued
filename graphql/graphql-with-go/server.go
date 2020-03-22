package main

import (
	"net/http"
	"os"

	"graphql-with-go/view/template"
)

func init() {
	wd, _ := os.Getwd()
	template.InitTemplate(wd + "/view/template/")
	_ = importJSONDataFromFile("data.json", &data)
}

func page(w http.ResponseWriter, r *http.Request) {
	template.RenderTemplate(w, "index.html", nil)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", page)
	mux.HandleFunc("/graphql", handler)
	http.ListenAndServe(":8080", mux)
}
