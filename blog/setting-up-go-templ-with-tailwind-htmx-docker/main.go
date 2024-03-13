package main

import (
	"embed"
	"log"
	"net/http"
	"spendings/components"

	"github.com/a-h/templ"
)

//go:embed static/*
var static embed.FS

func main() {
	homePage := components.Index()
	pagesHandler := http.NewServeMux()
	pagesHandler.Handle("/", templ.Handler(homePage))
	pagesHandler.Handle("/static/", http.FileServer(http.FS(static)))

	log.Fatalln(http.ListenAndServe(":8080", pagesHandler))
}
