package main

import (
	"finalwork/internal/db"
	"finalwork/internal/handlers"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

const (
	staticDir = "/assets/"
)

func main() {
	fmt.Println("Server is starting...")
	fmt.Println("Database auto migrate...")
	err := db.AutoMigrate()
	if err != nil {
		log.Println(err)
	}

	fmt.Println("Server handling requests...")
	r := mux.NewRouter()

	staticPath := http.Dir("./web/assets")
	fs := http.FileServer(staticPath)

	prometheus.MustRegister(handlers.HttpCounter)

	r.HandleFunc("/", handlers.ShowIndexPage).Methods("GET")
	r.HandleFunc("/services", handlers.ShowServicesPage).Methods("GET")
	r.HandleFunc("/sustainableDevelopment", handlers.ShowSustainableDevelopment).Methods("GET")
	r.HandleFunc("/news", handlers.ShowNews).Methods("GET")
	r.HandleFunc("/contacts", handlers.ShowContacts).Methods("GET")

	// Admin page handlers
	r.HandleFunc("/images", handlers.ShowImagesPage).Methods("GET")
	r.HandleFunc("/newsAdd", handlers.NewsAddPage).Methods("GET")
	r.HandleFunc("/newsBrowser", handlers.NewsBrowserPage).Methods("GET")
	r.HandleFunc("/managerAdd", handlers.ManagerAddPage).Methods("GET")
	r.HandleFunc("/managerBrowser", handlers.ManagerBrowserPage).Methods("GET")

	//Prometey handler
	r.Handle("/metrics", promhttp.Handler())
	r.PathPrefix(staticDir).Handler(http.StripPrefix(staticDir, fs)).Methods("GET")

	fmt.Println("Server is running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", r))
}
