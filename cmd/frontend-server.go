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

// main Главная функция запуска web-сервера frontend-server
func main() {
	fmt.Println("Frontend server is starting...")
	fmt.Println("Database auto migrate...")
	err := db.AutoMigrate()
	if err != nil {
		log.Println(err)
	}
	err = db.CreateDefaultUser()
	if err != nil {
		log.Println(err)
	}

	fmt.Println("Server handling requests...")
	router := mux.NewRouter()
	//	router.Use(handlers.CORS)

	staticPath := http.Dir("./web/assets")
	fs := http.FileServer(staticPath)

	prometheus.MustRegister(handlers.HttpCounter)

	// Public Routes
	router.HandleFunc("/", handlers.ShowIndexPage).Methods("GET")
	router.HandleFunc("/services", handlers.ShowServicesPage).Methods("GET")
	router.HandleFunc("/sustainableDevelopment", handlers.ShowSustainableDevelopment).Methods("GET")
	//	router.HandleFunc("/news", handlers.ShowNews).Methods("GET")
	router.HandleFunc("/news", handlers.ShowNews).Methods("GET")
	router.HandleFunc("/news/{id:[0-9]+}", handlers.ShowNews).Methods("GET")
	router.HandleFunc("/contacts", handlers.ShowContacts).Methods("GET")

	//Prometheus handler
	router.Handle("/metrics", promhttp.Handler())
	router.PathPrefix(staticDir).Handler(http.StripPrefix(staticDir, fs)).Methods("GET")

	// Authentification
	router.HandleFunc("/login", handlers.ShowLoginPage).Methods("GET")
	router.HandleFunc("/logout", handlers.HandlerLogout).Methods("POST")

	// Private Routes
	protected := router.PathPrefix("/").Subrouter()
	protected.Use(handlers.JWTAuth)
	// Admin page handlers
	protected.HandleFunc("/images", handlers.ShowImagesPage).Methods("GET")
	protected.HandleFunc("/newsAdd", handlers.NewsAddPage).Methods("GET")
	protected.HandleFunc("/newsBrowser", handlers.NewsBrowserPage).Methods("GET")
	protected.HandleFunc("/managerAdd", handlers.ManagerAddPage).Methods("GET")
	protected.HandleFunc("/managerBrowser", handlers.ManagerBrowserPage).Methods("GET")

	log.Println("🚀 Frontend server started on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
