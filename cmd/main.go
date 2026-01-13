package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Server is running")
	r := mux.NewRouter()
	r.HandleFunc("/", showIndexPage).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", r))
}

func showIndexPage(writer http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("web/templates/index.html")
	if err != nil {
		fmt.Printf("Error parsing index.html: %v \n", err)
	}

	err = tmpl.Execute(writer, nil)
	if err != nil {
		http.Error(writer, "Error rendering template", http.StatusInternalServerError)
		log.Printf("Error executing template: %v", err)
	}

}
