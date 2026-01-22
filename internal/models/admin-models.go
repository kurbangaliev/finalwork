package models

type ImagePayload struct {
	Src    string `json:"src"`
	Name   string `json:"name"`
	Folder string `json:"folder"`
}

type UploadRequest struct {
	Images []ImagePayload `json:"images"`
}

type ImageInfo struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type NewsItem struct {
	Id      uint   `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Date    string `json:"date"`
	Image   string `json:"image"`
}
