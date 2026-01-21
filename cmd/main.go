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

func main() {
	fmt.Println("Server is starting...")
	fmt.Println("Database auto migrate...")
	db.AutoMigrate()

	fmt.Println("Server handling requests...")
	r := mux.NewRouter()
	staticDir := "/assets/"
	staticPath := http.Dir("./web/assets")
	fs := http.FileServer(staticPath)

	prometheus.MustRegister(handlers.HttpCounter)

	r.HandleFunc("/", handlers.ShowIndexPage).Methods("GET")
	r.HandleFunc("/services", handlers.ShowServicesPage).Methods("GET")
	r.HandleFunc("/sustainableDevelopment", handlers.ShowSustainableDevelopment).Methods("GET")
	r.HandleFunc("/news", handlers.ShowNews).Methods("GET")
	r.HandleFunc("/contacts", handlers.ShowContacts).Methods("GET")
	r.HandleFunc("/images", handlers.ShowImagesPage).Methods("GET")
	r.Handle("/metrics", promhttp.Handler())
	r.PathPrefix(staticDir).Handler(http.StripPrefix(staticDir, fs)).Methods("GET")

	fmt.Println("Server is running on port 8000...")
	log.Fatal(http.ListenAndServe(":8000", r))
}
