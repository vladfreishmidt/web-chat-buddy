package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

type application struct {
	templateCache map[string]*template.Template
}

func main() {
	// routes
	// web server

	templateCache, err := newTemplateCahce()
	if err != nil {
		log.Fatal(err)
		fmt.Println(err)
	}

	app := &application{
		templateCache: templateCache,
	}

	log.Println("Starting channel listener")
	go ListenToWsChannel()

	log.Printf("Starting server on %s", ":4000")
	err = http.ListenAndServe(":8080", app.routes())
	log.Fatal(err)
}
