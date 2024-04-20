package main

import (
	"embed"
	"log"
	"net/http"
	"templpwa/components"

	"github.com/a-h/templ"
)

//go:embed static/*
var static embed.FS

//go:generate npx tailwindcss build -i static/css/style.css -o static/css/tailwind.css -m

func main() {
	http.Handle("/", templ.Handler(components.Index()))
	http.Handle("/static/", http.FileServer(http.FS(static)))
	log.Println("starting server on port 8080")
	log.Fatalln(http.ListenAndServe(":8080", nil))
}
