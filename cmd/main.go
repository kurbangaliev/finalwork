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
	staticDir := "/assets/"
	staticPath := http.Dir("./web/assets")
	fs := http.FileServer(staticPath)

	r.HandleFunc("/", showIndexPage).Methods("GET")
	r.HandleFunc("/services", showServicesPage).Methods("GET")
	r.HandleFunc("/sustainableDevelopment", showSustainableDevelopment).Methods("GET")
	r.HandleFunc("/news", showNews).Methods("GET")
	r.HandleFunc("/contacts", showContacts).Methods("GET")
	r.PathPrefix(staticDir).Handler(http.StripPrefix(staticDir, fs)).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", r))
}

func showContacts(writer http.ResponseWriter, request *http.Request) {
	tmpl, err := template.ParseFiles("web/templates/contacts.html")
	if err != nil {
		fmt.Printf("Error parsing contacts.html: %v \n", err)
	}

	err = tmpl.Execute(writer, nil)
	if err != nil {
		http.Error(writer, "Error rendering template", http.StatusInternalServerError)
		log.Printf("Error executing template: %v", err)
	}
}

func showNews(writer http.ResponseWriter, request *http.Request) {
	tmpl, err := template.ParseFiles("web/templates/news.html")
	if err != nil {
		fmt.Printf("Error parsing news.html: %v \n", err)
	}

	err = tmpl.Execute(writer, nil)
	if err != nil {
		http.Error(writer, "Error rendering template", http.StatusInternalServerError)
		log.Printf("Error executing template: %v", err)
	}
}

func showSustainableDevelopment(writer http.ResponseWriter, request *http.Request) {
	tmpl, err := template.ParseFiles("web/templates/sustainableDevelopment.html")
	if err != nil {
		fmt.Printf("Error parsing sustainableDevelopment.html: %v \n", err)
	}

	err = tmpl.Execute(writer, nil)
	if err != nil {
		http.Error(writer, "Error rendering template", http.StatusInternalServerError)
		log.Printf("Error executing template: %v", err)
	}
}

func showServicesPage(writer http.ResponseWriter, request *http.Request) {
	tmpl, err := template.ParseFiles("web/templates/services.html")
	if err != nil {
		fmt.Printf("Error parsing services.html: %v \n", err)
	}

	err = tmpl.Execute(writer, nil)
	if err != nil {
		http.Error(writer, "Error rendering template", http.StatusInternalServerError)
		log.Printf("Error executing template: %v", err)
	}
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
