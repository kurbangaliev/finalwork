package handlers

import (
	"encoding/json"
	"finalwork/internal/db"
	"finalwork/internal/utils"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

/* =================== Objects =================== */

// HandleGetObjects GET /news /managers
func HandleGetObjects[T comparable](w http.ResponseWriter, r *http.Request) {
	items, err := db.SelectAll[T]()
	if err != nil {
		log.Println(err)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(`{"error": "` + err.Error() + `"}`)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(items)
}

// HandleGetObject GET /news/{id}
func HandleGetObject[T comparable](w http.ResponseWriter, r *http.Request) {
	var item T

	vars := mux.Vars(r)
	strId := vars["id"]
	id, err := strconv.Atoi(strId)
	if err != nil {
		http.Error(w, "Invalid id", http.StatusBadRequest)
		return
	}

	utils.SetField(&item, "ID", uint(id))

	log.Printf("Get Item: %v\n", item)
	item, err = db.Select[T](item)
	if err != nil {
		http.Error(w, "Failed to get item", http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(item)
}

// HandleAddObject POST /news /managers
func HandleAddObject[T comparable](w http.ResponseWriter, r *http.Request) {
	var item T
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		log.Println(err)
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	//Валидацию сделать через интерфейс
	// Валидация
	//if item.Title == "" || item.Content == "" || item.Date == "" || item.Image == "" {
	//	http.Error(w, "All fields required", http.StatusBadRequest)
	//	return
	//}

	// Добавляем метку времени
	//item.Date = item.Date + "T" + time.Now().Format("15:04:05")

	err := db.SaveObject(item)
	if err != nil {
		log.Println(err)
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("News saved"))
}

// HandleEditObject PUT /news/{id} /managers/{id}
func HandleEditObject[T comparable](w http.ResponseWriter, r *http.Request) {
	var item T
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	err := db.UpdateObject(item)
	if err != nil {
		http.Error(w, "Failed to update item", http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("News updated"))
}

// HandleDeleteObject DELETE /news/{id} /managers/{id}
func HandleDeleteObject[T comparable](w http.ResponseWriter, r *http.Request) {
	var item T
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	err := db.DeleteObject(item)
	if err != nil {
		http.Error(w, "Failed to delete item", http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("News deleted"))
}
