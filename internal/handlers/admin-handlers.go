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

// UploadHandler POST /upload
func UploadHandler(w http.ResponseWriter, r *http.Request) {
	HttpCounter.With(prometheus.Labels{"path": r.URL.Path}).Inc()

	var req models.UploadRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	for _, img := range req.Images {
		folderPath := filepath.Join(UploadDir, utils.SanitizeFolder(img.Folder))
		err := os.MkdirAll(folderPath, 0755)
		if err != nil {
			log.Println("❌", err)
			http.Error(w, "Failed to save image", http.StatusInternalServerError)
			return
		}
		if err := saveImage(img, folderPath); err != nil {
			log.Println("❌", err)
			http.Error(w, "Failed to save image", http.StatusInternalServerError)
			return
		}
	}

	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte("Images saved"))
	if err != nil {
		return
	}
}

// saveImage Save image to local file storage in upload folder
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

// ImagesHandler GET /images
func ImagesHandler(w http.ResponseWriter, r *http.Request) {
	HttpCounter.With(prometheus.Labels{"path": r.URL.Path}).Inc()

	folder := utils.SanitizeFolder(r.URL.Query().Get("folder"))
	if folder == "" {
		folder = "default"
	}
	folderPath := filepath.Join(UploadDir, folder)

	err := os.MkdirAll(folderPath, 0755)
	if err != nil {
		http.Error(w, "Failed to create directory", http.StatusInternalServerError)
		return
	}

	files, err := os.ReadDir(folderPath)
	if err != nil {
		http.Error(w, "Failed to read directory", http.StatusInternalServerError)
		return
	}

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

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(images)
	if err != nil {
		http.Error(w, "Failed to encode images", http.StatusInternalServerError)
		return
	}
}

/* ================= DELETE ================= */

// DeleteHandler DELETE /image/
func DeleteHandler(w http.ResponseWriter, r *http.Request) {
	HttpCounter.With(prometheus.Labels{"path": r.URL.Path}).Inc()

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
	_, err := w.Write([]byte("Deleted"))
	if err != nil {
		http.Error(w, "Failed to write status", http.StatusInternalServerError)
		return
	}
}

/* ================= GET FOLDERS ================= */

// FoldersHandler GET /folders
func FoldersHandler(w http.ResponseWriter, r *http.Request) {
	HttpCounter.With(prometheus.Labels{"path": r.URL.Path}).Inc()

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
	err = json.NewEncoder(w).Encode(folders)
	if err != nil {
		http.Error(w, "Failed to encode folders", http.StatusInternalServerError)
		return
	}
}
