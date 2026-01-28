package main

import (
	"finalwork/internal/handlers"
	"finalwork/internal/models"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

const (
	uploadPrefix = "/uploads/"
)

func main() {
	if err := os.MkdirAll(handlers.UploadDir, 0755); err != nil {
		log.Fatal(err)
	}

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:8080"},
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Content-Type", "Authorization", "Content-Length", "Accept-Encoding", "X-CSRF-Token"},
		Debug:            false,
	})

	r := mux.NewRouter()
	//r.Use(handlers.CORS)

	//images handlers
	r.HandleFunc("/upload", handlers.UploadHandler).Methods("POST")
	r.HandleFunc("/images", handlers.ImagesHandler).Methods("GET")
	r.HandleFunc("/image/{filename}", handlers.DeleteHandler).Methods("DELETE")
	r.HandleFunc("/folders", handlers.FoldersHandler).Methods("GET")
	//news handlers
	//r.HandleFunc("/news/", handlers.HandleAddNews).Methods("POST")
	r.HandleFunc("/news/", handlers.HandleAddObject[models.News]).Methods("POST")
	//r.HandleFunc("/news/{id}", handlers.HandleEditNews).Methods("PUT")
	r.HandleFunc("/news/{id}", handlers.HandleEditObject[models.News]).Methods("PUT")
	//r.HandleFunc("/news/", handlers.HandleGetNews).Methods("GET")
	r.HandleFunc("/news/", handlers.HandleGetObjects[models.News]).Methods("GET")
	//r.HandleFunc("/news/{id}", handlers.HandleDeleteNews).Methods("DELETE")
	r.HandleFunc("/news/{id}", handlers.HandleDeleteObject[models.News]).Methods("DELETE")
	//manager handlers
	r.HandleFunc("/managers/", handlers.HandleAddManager).Methods("POST")
	r.HandleFunc("/managers/{id}", handlers.HandleEditManagers).Methods("PUT")
	//r.HandleFunc("/managers/", handlers.HandleGetManagers).Methods("GET")
	r.HandleFunc("/managers/", handlers.HandleGetObjects[models.Manager]).Methods("GET")
	r.HandleFunc("/managers/{id}", handlers.HandleDeleteManager).Methods("DELETE")
	//security handlers
	r.HandleFunc("/login", handlers.HandleLogin).Methods("POST")
	uploadPath := http.Dir(handlers.UploadDir)
	fileServer := http.FileServer(uploadPath)
	r.PathPrefix(uploadPrefix).Handler(http.StripPrefix(uploadPrefix, fileServer)).Methods("GET")

	log.Println("🚀 Server started on http://localhost:8081")

	handler := c.Handler(r)
	log.Fatal(http.ListenAndServe(":8081", handler))
}
