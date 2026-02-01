package models

// ImagePayload - структура сохранения изображений
type ImagePayload struct {
	Src    string `json:"src"`
	Name   string `json:"name"`
	Folder string `json:"folder"`
}

// UploadRequest - структура загрузки массива изображений
type UploadRequest struct {
	Images []ImagePayload `json:"images"`
}

// ImageInfo - Структура изображения
type ImageInfo struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

// LoginRequest - Структура аутентификации
type LoginRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}
