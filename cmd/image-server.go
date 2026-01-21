package main

import (
	"encoding/base64"
	"encoding/json"
	"finalwork/internal/db"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

const (
	uploadDir = "web\\assets\\uploads"
	newsFile  = "news.json"
)

type ImagePayload struct {
	Src    string `json:"src"`
	Name   string `json:"name"`
	Folder string `json:"folder"`
}

var newsMutex sync.Mutex // для безопасной работы с файлом новостей

type UploadRequest struct {
	Images []ImagePayload `json:"images"`
}

type ImageInfo struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type NewsItem struct {
	Title   string `json:"title"`
	Content string `json:"content"`
	Date    string `json:"date"`
	Image   string `json:"image"`
}

func main() {
	if err := os.MkdirAll(uploadDir, 0755); err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/upload", cors(uploadHandler))
	http.HandleFunc("/images", cors(imagesHandler))
	http.HandleFunc("/image/", cors(deleteHandler))
	http.HandleFunc("/folders", cors(foldersHandler))
	http.HandleFunc("/news", cors(newsHandler)) // новый обработчик новостей

	http.Handle("/uploads/", http.StripPrefix("/uploads/", http.FileServer(http.Dir(uploadDir))))

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
		folderPath := filepath.Join(uploadDir, sanitizeFolder(img.Folder))
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

func saveImage(img ImagePayload, folderPath string) error {
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
	path := filepath.Join(folderPath, filename)

	err = db.SaveImage(data, filename, path, img.Folder)
	if err != nil {
		log.Printf("Erorr while saving image: %v\n", err)
	}
	return os.WriteFile(path, data, 0644)
}

/* ================= LOAD ================= */

func imagesHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	folder := sanitizeFolder(r.URL.Query().Get("folder"))
	if folder == "" {
		folder = "default"
	}
	folderPath := filepath.Join(uploadDir, folder)
	os.MkdirAll(folderPath, 0755)

	files, err := os.ReadDir(folderPath)
	if err != nil {
		http.Error(w, "Failed to read directory", http.StatusInternalServerError)
		return
	}

	var images []ImageInfo
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		images = append(images, ImageInfo{
			Name: file.Name(),
			URL:  "/uploads/" + folder + "/" + file.Name(),
		})
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(images)
}

/* ================= GET FOLDERS ================= */
func foldersHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	entries, err := os.ReadDir(uploadDir)
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

/* ================= DELETE ================= */

func deleteHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	parts := strings.SplitN(r.URL.Path[len("/image/"):], "?", 2)
	name := sanitizeFilename(parts[0])
	folder := sanitizeFolder(r.URL.Query().Get("folder"))
	if folder == "" {
		folder = "default"
	}
	path := filepath.Join(uploadDir, folder, name)

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

/* ================= UTILS ================= */

func sanitizeFilename(name string) string {
	name = filepath.Base(name)
	name = strings.ReplaceAll(name, "..", "")
	return name
}

func sanitizeFolder(folder string) string {
	folder = filepath.Base(folder)
	folder = strings.ReplaceAll(folder, "..", "")
	if folder == "" {
		folder = "default"
	}
	return folder
}

func cors(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, DELETE, OPTIONS")
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}
		next(w, r)
	}
}

/* =================== NEWS =================== */
func newsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		// обработка добавления новости (уже реализована)
		handleAddNews(w, r)
		return
	}
	if r.Method == http.MethodGet {
		// возвращаем список всех новостей
		newsMutex.Lock()
		defer newsMutex.Unlock()

		var newsList []NewsItem
		if _, err := os.Stat(newsFile); err == nil {
			data, _ := os.ReadFile(newsFile)
			json.Unmarshal(data, &newsList)
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(newsList)
		return
	}
	http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
}

func handleAddNews(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var item NewsItem
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// Валидация
	if item.Title == "" || item.Content == "" || item.Date == "" || item.Image == "" {
		http.Error(w, "All fields required", http.StatusBadRequest)
		return
	}

	// Добавляем метку времени
	item.Date = item.Date + "T" + time.Now().Format("15:04:05")

	// Сохраняем в файл с блокировкой
	newsMutex.Lock()
	defer newsMutex.Unlock()

	var newsList []NewsItem

	// Если файл существует, читаем существующие новости
	if _, err := os.Stat(newsFile); err == nil {
		data, _ := os.ReadFile(newsFile)
		json.Unmarshal(data, &newsList)
	}

	newsList = append(newsList, item)
	data, _ := json.MarshalIndent(newsList, "", "  ")
	if err := os.WriteFile(newsFile, data, 0644); err != nil {
		http.Error(w, "Failed to save news", http.StatusInternalServerError)
		return
	}

	err := db.SaveNews(item.Title, item.Content, item.Date, item.Image)
	if err != nil {
		log.Println(err)
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("News saved"))
}
