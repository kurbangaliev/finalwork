package handlers

import (
	"encoding/json"
	"finalwork/internal/db"
	"finalwork/internal/models"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

/* =================== NEWS =================== */

// Deprecated: HandleEditNews PUT /news/{id}
func HandleEditNews(w http.ResponseWriter, r *http.Request) {
	var item models.News
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	err := db.UpdateNews(item)
	if err != nil {
		http.Error(w, "Failed to update item", http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("News updated"))
}

// Deprecated: HandleDeleteNews DELETE /news/{id}
func HandleDeleteNews(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	strId := vars["id"]
	id, err := strconv.Atoi(strId)
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

// Deprecated: HandleAddNews POST /news
func HandleAddNews(w http.ResponseWriter, r *http.Request) {
	var item models.News
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

	err := db.SaveNews(item)
	if err != nil {
		log.Println(err)
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("News saved"))
}
