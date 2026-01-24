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
	//	r.Use(handlers.CORS)

	staticPath := http.Dir("./web/assets")
	fs := http.FileServer(staticPath)

	prometheus.MustRegister(handlers.HttpCounter)

	// Public Routes
	r.HandleFunc("/", handlers.ShowIndexPage).Methods("GET")
	r.HandleFunc("/services", handlers.ShowServicesPage).Methods("GET")
	r.HandleFunc("/sustainableDevelopment", handlers.ShowSustainableDevelopment).Methods("GET")
	r.HandleFunc("/news", handlers.ShowNews).Methods("GET")
	r.HandleFunc("/contacts", handlers.ShowContacts).Methods("GET")

	//Prometey handler
	r.Handle("/metrics", promhttp.Handler())
	r.PathPrefix(staticDir).Handler(http.StripPrefix(staticDir, fs)).Methods("GET")

	// Authentification
	r.HandleFunc("/login", handlers.ShowLoginPage).Methods("GET")
	r.HandleFunc("/logout", handlers.HandlerLogout).Methods("POST")

	// Private Routes
	protected := r.PathPrefix("/").Subrouter()
	protected.Use(handlers.JWTAuth)
	// Admin page handlers
	protected.HandleFunc("/images", handlers.ShowImagesPage).Methods("GET")
	protected.HandleFunc("/newsAdd", handlers.NewsAddPage).Methods("GET")
	protected.HandleFunc("/newsBrowser", handlers.NewsBrowserPage).Methods("GET")
	protected.HandleFunc("/managerAdd", handlers.ManagerAddPage).Methods("GET")
	protected.HandleFunc("/managerBrowser", handlers.ManagerBrowserPage).Methods("GET")

	log.Println("🚀 Server started on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
