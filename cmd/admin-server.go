package main

import (
	"finalwork/internal/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

const (
	uploadPrefix = "/uploads/"
)

func main() {
	if err := os.MkdirAll(handlers.UploadDir, 0755); err != nil {
		log.Fatal(err)
	}

	r := mux.NewRouter()

	r.HandleFunc("/upload", handlers.Cors(handlers.UploadHandler)).Methods("POST")
	r.HandleFunc("/images", handlers.Cors(handlers.ImagesHandler)).Methods("GET")
	r.HandleFunc("/image/", handlers.Cors(handlers.DeleteHandler)).Methods("DELETE")
	r.HandleFunc("/folders", handlers.Cors(handlers.FoldersHandler)).Methods("GET")
	r.HandleFunc("/news/", handlers.Cors(handlers.HandleAddNews)).Methods("POST")
	r.HandleFunc("/news/", handlers.Cors(handlers.HandleEditNews)).Methods("PUT")
	r.HandleFunc("/news/", handlers.Cors(handlers.HandleGetNews)).Methods("GET")
	r.HandleFunc("/news/", handlers.Cors(handlers.HandleDeleteNews)).Methods("DELETE")

	uploadPath := http.Dir(handlers.UploadDir)
	fileServer := http.FileServer(uploadPath)
	r.PathPrefix(uploadPrefix).Handler(http.StripPrefix(uploadPrefix, fileServer)).Methods("GET")

	log.Println("🚀 Server started on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
