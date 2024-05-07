package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"html/template"
	"log"
	"net/http"
)

type Application struct {
	page *template.Template
	db   *sql.DB
}

func main() {
	// Serving static files (css, javascript, images and other assets)
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	// Parse the main page as a template
	page, errpage := template.ParseFiles("./ui/html/index.html")
	if errpage != nil {
		log.Fatal(errpage)
	}
	db, errdb := sql.Open("sqlite3", "./db/db.sqlite3")
	if errdb != nil {
		log.Fatal(errdb)
	}
	app := Application{page, db}
	// Creating a server and giving a route to it
	mux := http.NewServeMux()
	mux.HandleFunc("/", app.homepage)
	mux.HandleFunc("/api/meetups", app.getMeetupsApi)
	mux.HandleFunc("/api/meetup", app.getMeetupInfoApi)
	// Giving a route to a static file server handler
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	// Launching a server and logging its actions and errors
	log.Println("Starting a server on a port 4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
