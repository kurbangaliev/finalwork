package handlers

import (
	"finalwork/internal/db"
	"finalwork/internal/models"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus"
	"gorm.io/gorm"
)

var HttpCounter = prometheus.NewCounterVec(prometheus.CounterOpts{
	Name: "http_request_total",
	Help: "Total numbers of HTTP Requests"},
	[]string{"path"},
)

type DataNews struct {
	SelectedId int
	AllNews    []models.News
	Selected   models.News
}

func HandlerTemplate(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

type GenericFunc[T comparable] func() ([]T, error)
type DataFuncParams func(map[string]string) (any, error)

func SelectDataNews(vars map[string]string) (any, error) {
	strId := vars["id"]
	var selectedId int

	news, err := db.SelectAll[models.News]()
	if err != nil {
		log.Fatal(err)
	}

	selectedId, err = strconv.Atoi(strId)
	if err != nil {
		selectedId = 0
	}

	var selectedNews = models.News{
		Model: gorm.Model{
			ID: uint(selectedId),
		},
	}

	if selectedNews, err = db.Select[models.News](selectedNews); err != nil {
		log.Fatal(err)
	}

	var dataNews = DataNews{
		AllNews:  news,
		Selected: selectedNews,
	}
	return dataNews, nil
}

func ShowTemplatePage(templatePage string, handler http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		log.Printf("ShowTemplatePage [%s] \n]...", templatePage)
		HttpCounter.With(prometheus.Labels{"path": request.URL.Path}).Inc()
		tmpl, err := template.ParseFiles(templatePage)
		if err != nil {
			fmt.Printf("Error parsing services.html: %v \n", err)
		}

		err = tmpl.Execute(writer, nil)
		if err != nil {
			http.Error(writer, "Error rendering template page "+templatePage, http.StatusInternalServerError)
			log.Printf("Error executing template: [%s] %v", templatePage, err)
		}
		handler.ServeHTTP(writer, request)
	})
}

func ShowTemplatePageGeneric[T comparable](templatePage string, handler http.Handler, dataFunc GenericFunc[T]) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		log.Printf("ShowTemplatePageGeneric [%s]...", templatePage)
		HttpCounter.With(prometheus.Labels{"path": request.URL.Path}).Inc()
		tmpl, err := template.ParseFiles(templatePage)
		if err != nil {
			fmt.Printf("Error parsing services.html: %v \n", err)
		}

		data, err := dataFunc()
		if err != nil {
			http.Error(writer, "Error rendering template page "+templatePage, http.StatusInternalServerError)
			log.Printf("Error executing template: [%s] %v", templatePage, err)
		}

		err = tmpl.Execute(writer, data)
		if err != nil {
			http.Error(writer, "Error rendering template page "+templatePage, http.StatusInternalServerError)
			log.Printf("Error executing template: [%s] %v", templatePage, err)
		}
		handler.ServeHTTP(writer, request)
	})
}

func ShowTemplatePageParams(templatePage string, handler http.Handler, dataFunc DataFuncParams) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		log.Printf("ShowTemplatePageParams [%s]...", templatePage)
		HttpCounter.With(prometheus.Labels{"path": request.URL.Path}).Inc()
		vars := mux.Vars(request)
		tmpl, err := template.ParseFiles(templatePage)
		if err != nil {
			fmt.Printf("Error parsing services.html: %v \n", err)
		}

		data, err := dataFunc(vars)

		if err != nil {
			http.Error(writer, "Error rendering template page "+templatePage, http.StatusInternalServerError)
			log.Printf("Error executing template: [%s] %v", templatePage, err)
		}

		err = tmpl.Execute(writer, data)
		if err != nil {
			http.Error(writer, "Error rendering template page "+templatePage, http.StatusInternalServerError)
			log.Printf("Error executing template: [%s] %v", templatePage, err)
		}
		handler.ServeHTTP(writer, request)
	})
}
