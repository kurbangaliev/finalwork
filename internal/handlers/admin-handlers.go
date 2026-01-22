package handlers

import (
	"encoding/base64"
	"encoding/json"
	"finalwork/internal/db"
	"finalwork/internal/models"
	"finalwork/internal/utils"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

const (
	UploadDir = "./web/assets/uploads"
)

/* ================= UPLOAD ================= */

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	HttpCounter.With(prometheus.Labels{"path": r.URL.Path}).Inc()
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req models.UploadRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	for _, img := range req.Images {
		folderPath := filepath.Join(UploadDir, utils.SanitizeFolder(img.Folder))
		os.MkdirAll(folderPath, 0755)
		if err := saveImage(img, folderPath); err != nil {
			log.Println("❌", err)
			http.Error(w, "Failed to save image", http.StatusInternalServerError)
			return
		}
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Images saved"))
}

func saveImage(img models.ImagePayload, folderPath string) error {
	parts := strings.Split(img.Src, ",")
	if len(parts) != 2 {
		return fmt.Errorf("invalid base64 image")
	}
	data, err := base64.StdEncoding.DecodeString(parts[1])
	if err != nil {
		return err
	}
	filename := utils.SanitizeFilename(img.Name)
	filename = fmt.Sprintf("%d_%s", time.Now().UnixNano(), filename)
	path := filepath.Join(folderPath, filename)

	err = db.SaveImage(data, filename, path, img.Folder)
	if err != nil {
		log.Printf("Erorr while saving image: %v\n", err)
	}
	return os.WriteFile(path, data, 0644)
}

/* ================= LOAD ================= */

func ImagesHandler(w http.ResponseWriter, r *http.Request) {
	HttpCounter.With(prometheus.Labels{"path": r.URL.Path}).Inc()
	log.Println("ImagesHandler")
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	folder := utils.SanitizeFolder(r.URL.Query().Get("folder"))
	if folder == "" {
		folder = "default"
	}
	folderPath := filepath.Join(UploadDir, folder)
	log.Println("Folder:", folderPath)
	os.MkdirAll(folderPath, 0755)

	files, err := os.ReadDir(folderPath)
	if err != nil {
		http.Error(w, "Failed to read directory", http.StatusInternalServerError)
		return
	}
	log.Println("Files:", len(files))
	var images []models.ImageInfo
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		images = append(images, models.ImageInfo{
			Name: file.Name(),
			URL:  "/uploads/" + folder + "/" + file.Name(),
		})
	}

	log.Println("Images:", images)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(images)
}

/* ================= DELETE ================= */

func DeleteHandler(w http.ResponseWriter, r *http.Request) {
	HttpCounter.With(prometheus.Labels{"path": r.URL.Path}).Inc()
	if r.Method != http.MethodDelete {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	parts := strings.SplitN(r.URL.Path[len("/image/"):], "?", 2)
	name := utils.SanitizeFilename(parts[0])
	folder := utils.SanitizeFolder(r.URL.Query().Get("folder"))
	if folder == "" {
		folder = "default"
	}
	path := filepath.Join(UploadDir, folder, name)

	if _, err := os.Stat(path); os.IsNotExist(err) {
		http.Error(w, "File not found", http.StatusNotFound)
		return
	}

	if err := os.Remove(path); err != nil {
		http.Error(w, "Failed to delete file", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Deleted"))
}

/* ================= GET FOLDERS ================= */

func FoldersHandler(w http.ResponseWriter, r *http.Request) {
	HttpCounter.With(prometheus.Labels{"path": r.URL.Path}).Inc()
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	entries, err := os.ReadDir(UploadDir)
	if err != nil {
		http.Error(w, "Failed to read upload directory", http.StatusInternalServerError)
		return
	}

	var folders []string
	for _, entry := range entries {
		if entry.IsDir() {
			folders = append(folders, entry.Name())
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(folders)
}

/* ================= SET HEADERS ================= */

func Cors(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}
		next(w, r)
	}
}
