package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

const uploadDir = "web\\assets\\uploads"

type ImagePayload struct {
	Src  string `json:"src"`
	Name string `json:"name"`
}

type UploadRequest struct {
	Images []ImagePayload `json:"images"`
}

type ImageInfo struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

func main() {
	if err := os.MkdirAll(uploadDir, 0755); err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/upload", cors(uploadHandler))
	http.HandleFunc("/images", cors(imagesHandler))

	// Раздача файлов
	http.Handle("/uploads/",
		http.StripPrefix("/uploads/",
			http.FileServer(http.Dir(uploadDir)),
		),
	)

	log.Println("🚀 Server started on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

/* ================= UPLOAD ================= */

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req UploadRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	for _, img := range req.Images {
		if err := saveImage(img); err != nil {
			log.Println("❌", err)
			http.Error(w, "Failed to save image", http.StatusInternalServerError)
			return
		}
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Images saved"))
}

func saveImage(img ImagePayload) error {
	parts := strings.Split(img.Src, ",")
	if len(parts) != 2 {
		return fmt.Errorf("invalid base64 image")
	}

	data, err := base64.StdEncoding.DecodeString(parts[1])
	if err != nil {
		return err
	}

	filename := sanitizeFilename(img.Name)
	filename = fmt.Sprintf("%d_%s", time.Now().UnixNano(), filename)

	path := filepath.Join(uploadDir, filename)
	return os.WriteFile(path, data, 0644)
}

/* ================= LOAD ================= */

func imagesHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	files, err := os.ReadDir(uploadDir)
	if err != nil {
		http.Error(w, "Failed to read directory", http.StatusInternalServerError)
		return
	}

	var images []ImageInfo
	for _, file := range files {
		if file.IsDir() {
			continue
		}

		name := file.Name()
		images = append(images, ImageInfo{
			Name: name,
			URL:  "/uploads/" + name,
		})
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(images)
}

/* ================= UTILS ================= */

func sanitizeFilename(name string) string {
	name = filepath.Base(name)
	name = strings.ReplaceAll(name, "..", "")
	return name
}

func cors(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}
		next(w, r)
	}
}
