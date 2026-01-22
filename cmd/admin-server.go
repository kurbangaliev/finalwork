package main

import (
	"finalwork/internal/handlers"
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
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token"},
		Debug:            false,
	})

	r := mux.NewRouter()

	r.HandleFunc("/upload", handlers.UploadHandler).Methods("POST")
	r.HandleFunc("/images", handlers.ImagesHandler).Methods("GET")
	r.HandleFunc("/image/", handlers.DeleteHandler).Methods("DELETE")
	r.HandleFunc("/folders", handlers.FoldersHandler).Methods("GET")
	r.HandleFunc("/news/", handlers.HandleAddNews).Methods("POST")
	r.HandleFunc("/news/{id}", handlers.HandleEditNews).Methods("PUT")
	r.HandleFunc("/news/", handlers.HandleGetNews).Methods("GET")
	r.HandleFunc("/news/{id}", handlers.HandleDeleteNews).Methods("DELETE")

	uploadPath := http.Dir(handlers.UploadDir)
	fileServer := http.FileServer(uploadPath)
	r.PathPrefix(uploadPrefix).Handler(http.StripPrefix(uploadPrefix, fileServer)).Methods("GET")

	log.Println("🚀 Server started on http://localhost:8081")

	handler := c.Handler(r)
	log.Fatal(http.ListenAndServe(":8081", handler))
}
