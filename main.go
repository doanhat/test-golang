package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"text/template"
	"time"
)

type Welcome struct {
	Name string
	Time string
}

func main() {

	welcome := Welcome{Name: "Anonymous", Time: time.Now().Format(time.Stamp)}

	templates := template.Must(template.ParseFiles("templates/welcome-template.html"))

	http.Handle("/static/", //final url can be anything
		http.StripPrefix("/static/",
			http.FileServer(http.Dir("static"))))

	http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}

		if name := r.FormValue("name"); name!= "" {
			welcome.Name = name
		}
		if err := templates.ExecuteTemplate(w, "welcome-template.html", welcome); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})
	port := os.Getenv("PORT")
	if port == "" {
		port = "8088"
		log.Printf("Defaulting to port %s", port)
	}

	log.Printf("Listening on port %s", port)
	log.Printf("Open http://localhost:%s in the browser", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}

