package main

import (
	"finalwork/internal/db"
	"finalwork/internal/handlers"
	"finalwork/internal/models"
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
	//"web/templates/index.html"
	//router.HandleFunc("/", handlers.ShowIndexPage).Methods("GET")
	router.Handle("/", handlers.ShowTemplatePageGeneric[models.News]("web/templates/index.html", http.HandlerFunc(handlers.HandlerTemplate), db.SelectAll)).Methods("GET")
	//"web/templates/services.html"
	//router.HandleFunc("/services", handlers.ShowServicesPage).Methods("GET")
	router.Handle("/services", handlers.ShowTemplatePage("web/templates/services.html", http.HandlerFunc(handlers.HandlerTemplate))).Methods("GET")
	//"web/templates/sustainableDevelopment.html"
	//router.HandleFunc("/sustainableDevelopment", handlers.ShowSustainableDevelopment).Methods("GET")
	router.Handle("/sustainableDevelopment", handlers.ShowTemplatePage("web/templates/sustainableDevelopment.html", http.HandlerFunc(handlers.HandlerTemplate))).Methods("GET")
	//	router.HandleFunc("/news", handlers.ShowNews).Methods("GET")
	//router.HandleFunc("/news", handlers.ShowNews).Methods("GET")
	//"web/templates/news.html"
	router.Handle("/news", handlers.ShowTemplatePageParams("web/templates/news.html", http.HandlerFunc(handlers.HandlerTemplate), handlers.SelectDataNews)).Methods("GET")
	//router.HandleFunc("/news/{id:[0-9]+}", handlers.ShowNews).Methods("GET")
	router.Handle("/news/{id:[0-9]+}", handlers.ShowTemplatePageParams("web/templates/news.html", http.HandlerFunc(handlers.HandlerTemplate), handlers.SelectDataNews)).Methods("GET")
	//"web/templates/contacts.html"
	//router.HandleFunc("/contacts", handlers.ShowContacts).Methods("GET")
	router.Handle("/contacts", handlers.ShowTemplatePageGeneric[models.Manager]("web/templates/contacts.html", http.HandlerFunc(handlers.HandlerTemplate), db.SelectAll)).Methods("GET")

	//Prometheus handler
	router.Handle("/metrics", promhttp.Handler())
	router.PathPrefix(staticDir).Handler(http.StripPrefix(staticDir, fs)).Methods("GET")

	// Authentification
	//	router.HandleFunc("/login", handlers.ShowLoginPage).Methods("GET")
	router.Handle("/login", handlers.ShowTemplatePage("web/templates/admin/auth.html", http.HandlerFunc(handlers.HandlerTemplate))).Methods("GET")
	router.HandleFunc("/logout", handlers.HandlerLogout).Methods("POST")

	// Private Routes
	protected := router.PathPrefix("/").Subrouter()
	protected.Use(handlers.JWTAuth)
	// Admin page handlers
	//"web/templates/admin/images.html"
	//protected.HandleFunc("/images", handlers.ShowImagesPage).Methods("GET")
	protected.Handle("/images", handlers.ShowTemplatePage("web/templates/admin/images.html", http.HandlerFunc(handlers.HandlerTemplate))).Methods("GET")
	//"web/templates/admin/newsAdd.html"
	//protected.HandleFunc("/newsAdd", handlers.NewsAddPage).Methods("GET")
	protected.Handle("/newsAdd", handlers.ShowTemplatePage("web/templates/admin/newsAdd.html", http.HandlerFunc(handlers.HandlerTemplate))).Methods("GET")
	//"web/templates/admin/newsBrowser.html"
	//protected.HandleFunc("/newsBrowser", handlers.NewsBrowserPage).Methods("GET")
	protected.Handle("/newsBrowser", handlers.ShowTemplatePage("web/templates/admin/newsBrowser.html", http.HandlerFunc(handlers.HandlerTemplate))).Methods("GET")
	//"web/templates/admin/managerAdd.html"
	//protected.HandleFunc("/managerAdd", handlers.ManagerAddPage).Methods("GET")
	protected.Handle("/managerAdd", handlers.ShowTemplatePage("web/templates/admin/managerAdd.html", http.HandlerFunc(handlers.HandlerTemplate))).Methods("GET")
	//"web/templates/admin/managerBrowser.html"
	//protected.HandleFunc("/managerBrowser", handlers.ManagerBrowserPage).Methods("GET")
	protected.Handle("/managerBrowser", handlers.ShowTemplatePage("web/templates/admin/managerBrowser.html", http.HandlerFunc(handlers.HandlerTemplate))).Methods("GET")

	log.Println("🚀 Frontend server started on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
