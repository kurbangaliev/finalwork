package handlers

import (
	"encoding/json"
	"finalwork/internal/db"
	"finalwork/internal/models"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

/* =================== NEWS =================== */
// GET /news
func HandleGetNews(w http.ResponseWriter, r *http.Request) {
	newsList := db.SelectAllNews()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newsList)
}

// PUT /news/{index}
func HandleEditNews(w http.ResponseWriter, r *http.Request) {
	indexStr := strings.TrimPrefix(r.URL.Path, "/news/")
	id, err := strconv.Atoi(indexStr)
	if err != nil {
		http.Error(w, "Invalid id", http.StatusBadRequest)
		return
	}

	var item models.NewsItem
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	err = db.UpdateNews(uint(id), item.Title, item.Content, item.Date, item.Image)
	if err != nil {
		http.Error(w, "Failed to update item", http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("News updated"))
}

// DELETE /news/{index}
func HandleDeleteNews(w http.ResponseWriter, r *http.Request) {
	indexStr := strings.TrimPrefix(r.URL.Path, "/news/")
	id, err := strconv.Atoi(indexStr)
	if err != nil {
		http.Error(w, "Invalid id", http.StatusBadRequest)
		return
	}
	log.Printf("delete %d", id)

	err = db.DeleteNews(uint(id))
	if err != nil {
		http.Error(w, "Failed to delete item", http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("News deleted"))
}

func HandleAddNews(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var item models.NewsItem
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		log.Println(err)
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

	err := db.SaveNews(item.Title, item.Content, item.Date, item.Image)
	if err != nil {
		log.Println(err)
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("News saved"))
}
