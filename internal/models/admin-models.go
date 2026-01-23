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

type ManagerItem struct {
	Id       uint   `json:"id"`
	Name     string `json:"name"`
	JobTitle string `json:"job_title"`
	Address  string `json:"address"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
	Schedule string `json:"schedule"`
	Image    string `json:"image"`
}

type PartnerItem struct {
	Id    uint   `json:"id"`
	Name  string `json:"name"`
	Image string `json:"image"`
}
