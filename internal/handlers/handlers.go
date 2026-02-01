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
)

var HttpCounter = prometheus.NewCounterVec(prometheus.CounterOpts{
	Name: "http_request_total",
	Help: "Total numbers of HTTP Requests"},
	[]string{"path"},
)

// ShowContacts - Отображение страницы Контакты
func ShowContacts(writer http.ResponseWriter, request *http.Request) {
	HttpCounter.With(prometheus.Labels{"path": request.URL.Path}).Inc()
	tmpl, err := template.ParseFiles("web/templates/contacts.html")
	if err != nil {
		fmt.Printf("Error parsing contacts.html: %v \n", err)
	}

	//managers := db.SelectAllManagers()
	managers, err := db.SelectAll[models.Manager]()
	if err != nil {
		http.Error(writer, "Error load managers", http.StatusInternalServerError)
		log.Fatal(err)
	}

	err = tmpl.Execute(writer, managers)
	if err != nil {
		http.Error(writer, "Error rendering template", http.StatusInternalServerError)
		log.Printf("Error executing template: %v", err)
	}
}

type DataNews struct {
	SelectedId int
	AllNews    []models.News
}

// ShowNews - отображение страницы Новости
func ShowNews(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	strId := vars["id"]
	var selectedId int

	HttpCounter.With(prometheus.Labels{"path": request.URL.Path}).Inc()
	tmpl, err := template.ParseFiles("web/templates/news.html")
	if err != nil {
		fmt.Printf("Error parsing news.html: %v \n", err)
	}

	//news := db.SelectAllNews()
	news, err := db.SelectAll[models.News]()
	if err != nil {
		http.Error(writer, "Error load news", http.StatusInternalServerError)
		log.Fatal(err)
	}

	selectedId, err = strconv.Atoi(strId)
	if err != nil {
		selectedId = 0
	}
	if selectedId > len(news) {
		selectedId = 0
	}

	var dataNews = DataNews{
		SelectedId: selectedId,
		AllNews:    news,
	}
	//log.Printf("dataNews=%v\n", dataNews)

	err = tmpl.Execute(writer, dataNews)
	if err != nil {
		http.Error(writer, "Error rendering template", http.StatusInternalServerError)
		log.Printf("Error executing template: %v", err)
	}
}

// ShowSustainableDevelopment Отображение страницы Устойчивое развитие
func ShowSustainableDevelopment(writer http.ResponseWriter, request *http.Request) {
	HttpCounter.With(prometheus.Labels{"path": request.URL.Path}).Inc()
	tmpl, err := template.ParseFiles("web/templates/sustainableDevelopment.html")
	if err != nil {
		fmt.Printf("Error parsing sustainableDevelopment.html: %v \n", err)
	}

	err = tmpl.Execute(writer, nil)
	if err != nil {
		http.Error(writer, "Error rendering template", http.StatusInternalServerError)
		log.Printf("Error executing template: %v", err)
	}
}

// ShowServicesPage Отображение страницы Услуги
func ShowServicesPage(writer http.ResponseWriter, request *http.Request) {
	HttpCounter.With(prometheus.Labels{"path": request.URL.Path}).Inc()
	tmpl, err := template.ParseFiles("web/templates/services.html")
	if err != nil {
		fmt.Printf("Error parsing services.html: %v \n", err)
	}

	err = tmpl.Execute(writer, nil)
	if err != nil {
		http.Error(writer, "Error rendering template", http.StatusInternalServerError)
		log.Printf("Error executing template: %v", err)
	}
}

// ShowIndexPage Отображение главной стрнаницы
func ShowIndexPage(writer http.ResponseWriter, r *http.Request) {
	HttpCounter.With(prometheus.Labels{"path": r.URL.Path}).Inc()
	tmpl, err := template.ParseFiles("web/templates/index.html")
	if err != nil {
		fmt.Printf("Error parsing index.html: %v \n", err)
	}

	//	news := db.SelectAllNews()
	news, _ := db.SelectAll[models.News]()

	err = tmpl.Execute(writer, news)
	if err != nil {
		http.Error(writer, "Error rendering template", http.StatusInternalServerError)
		log.Printf("Error executing template: %v", err)
	}
}

// ShowImagesPage Отображение страницы Администрирование-Изображения
func ShowImagesPage(writer http.ResponseWriter, r *http.Request) {
	HttpCounter.With(prometheus.Labels{"path": r.URL.Path}).Inc()
	tmpl, err := template.ParseFiles("web/templates/admin/images.html")
	if err != nil {
		fmt.Printf("Error parsing images.html: %v \n", err)
	}

	err = tmpl.Execute(writer, nil)
	if err != nil {
		http.Error(writer, "Error rendering template", http.StatusInternalServerError)
		log.Printf("Error executing template: %v", err)
	}
}

// NewsAddPage Отображение страницы Администрирование-Создание новости
func NewsAddPage(writer http.ResponseWriter, r *http.Request) {
	HttpCounter.With(prometheus.Labels{"path": r.URL.Path}).Inc()
	tmpl, err := template.ParseFiles("web/templates/admin/newsAdd.html")
	if err != nil {
		fmt.Printf("Error parsing newsAdd.html: %v \n", err)
	}

	err = tmpl.Execute(writer, nil)
	if err != nil {
		http.Error(writer, "Error rendering template", http.StatusInternalServerError)
		log.Printf("Error executing template: %v", err)
	}
}

// NewsBrowserPage Отображение страницы Администрирование-Новости
func NewsBrowserPage(writer http.ResponseWriter, r *http.Request) {
	HttpCounter.With(prometheus.Labels{"path": r.URL.Path}).Inc()
	tmpl, err := template.ParseFiles("web/templates/admin/newsBrowser.html")
	if err != nil {
		fmt.Printf("Error parsing newsBrowser.html: %v \n", err)
	}

	err = tmpl.Execute(writer, nil)
	if err != nil {
		http.Error(writer, "Error rendering template", http.StatusInternalServerError)
		log.Printf("Error executing template: %v", err)
	}
}

// ManagerAddPage Отображение страницы Администрирование-Добавление руководителей
func ManagerAddPage(writer http.ResponseWriter, r *http.Request) {
	HttpCounter.With(prometheus.Labels{"path": r.URL.Path}).Inc()
	tmpl, err := template.ParseFiles("web/templates/admin/managerAdd.html")
	if err != nil {
		fmt.Printf("Error parsing managerAdd.html: %v \n", err)
	}

	err = tmpl.Execute(writer, nil)
	if err != nil {
		http.Error(writer, "Error rendering template", http.StatusInternalServerError)
		log.Printf("Error executing template: %v", err)
	}
}

// ManagerBrowserPage Отображение страницы Администрирование-Руководители
func ManagerBrowserPage(writer http.ResponseWriter, r *http.Request) {
	HttpCounter.With(prometheus.Labels{"path": r.URL.Path}).Inc()
	tmpl, err := template.ParseFiles("web/templates/admin/managerBrowser.html")
	if err != nil {
		fmt.Printf("Error parsing managerBrowser.html: %v \n", err)
	}

	err = tmpl.Execute(writer, nil)
	if err != nil {
		http.Error(writer, "Error rendering template", http.StatusInternalServerError)
		log.Printf("Error executing template: %v", err)
	}
}

// ShowLoginPage Отображение страницы Аутентификация
func ShowLoginPage(writer http.ResponseWriter, r *http.Request) {
	HttpCounter.With(prometheus.Labels{"path": r.URL.Path}).Inc()
	tmpl, err := template.ParseFiles("web/templates/admin/auth.html")
	if err != nil {
		fmt.Printf("Error parsing auth.html: %v \n", err)
	}

	err = tmpl.Execute(writer, nil)
	if err != nil {
		http.Error(writer, "Error rendering template", http.StatusInternalServerError)
		log.Printf("Error executing template: %v", err)
	}
}
