package handlers

import (
	"encoding/json"
	"finalwork/internal/db"
	"finalwork/internal/models"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

/* =================== Managers =================== */

// HandleEditManagers PUT /managers/{id}
func HandleEditManagers(w http.ResponseWriter, r *http.Request) {
	var item models.Manager
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	err := db.UpdateManager(item)
	if err != nil {
		http.Error(w, "Failed to update item", http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Manager updated"))
}

// HandleDeleteManager DELETE /managers/{id}
func HandleDeleteManager(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	strId := vars["id"]
	id, err := strconv.Atoi(strId)
	if err != nil {
		http.Error(w, "Invalid id", http.StatusBadRequest)
		return
	}
	log.Printf("delete %d", id)

	err = db.DeleteManager(uint(id))
	if err != nil {
		http.Error(w, "Failed to delete item", http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("News deleted"))
}

// HandleAddManager POST /managers
func HandleAddManager(w http.ResponseWriter, r *http.Request) {
	var item models.Manager
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		log.Println(err)
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// Валидация
	if item.Name == "" || item.JobTitle == "" || item.Address == "" || item.Phone == "" || item.Email == "" || item.Schedule == "" || item.Image == "" {
		http.Error(w, "All fields required", http.StatusBadRequest)
		return
	}

	err := db.SaveManager(item)
	if err != nil {
		log.Println(err)
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("News saved"))
}
