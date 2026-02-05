package tests

import (
	"encoding/json"
	"finalwork/internal/handlers"
	"finalwork/internal/models"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHandleGetObjectsNewsHandlers(t *testing.T) {
	type want struct {
		statusCode  int
		contentType string
		url         string
	}

	tests := []struct {
		name string
		want want
	}{
		{
			name: "Test News Api",
			want: want{
				statusCode:  200,
				contentType: "application/json",
				url:         "/news/",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			request := httptest.NewRequest(http.MethodGet, "/news/", nil)
			request.Header.Set("Content-Type", "application/json")

			writer := httptest.NewRecorder()
			handlers.HandleGetObjects[models.News](writer, request)

			res := writer.Result()
			defer res.Body.Close()

			assert.Equal(t, tt.want.statusCode, res.StatusCode,
				"Expected status %d, got %d", tt.want.statusCode, res.StatusCode)

			if tt.want.contentType != "" {
				assert.Equal(t, tt.want.contentType, res.Header.Get("Content-Type"),
					"Expected Content-Type %s, got %s", tt.want.contentType, res.Header.Get("Content-Type"))
			}

			body, err := io.ReadAll(res.Body)
			defer res.Body.Close()
			if err != nil {
				t.Fatal(err)
			}

			if string(body) == "" {
				t.Errorf("Expected body, but result %s", string(body))
			}

			var news []models.News
			err = json.Unmarshal(body, &news)
			if err != nil {
				assert.NoError(t, err, "Response should be valid JSON")
			}
			if len(news) == 0 {
				t.Errorf("Expected length more than zero, got %d", len(news))
			}
			log.Printf("Count news: %d", len(news))
		})
	}
}

func TestHandleGetObjectsManagersHandlers(t *testing.T) {
	type want struct {
		statusCode  int
		contentType string
		url         string
	}

	tests := []struct {
		name string
		want want
	}{
		{
			name: "Test Managers Api",
			want: want{
				statusCode:  200,
				contentType: "application/json",
				url:         "/managers/",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			request := httptest.NewRequest(http.MethodGet, "/managers/", nil)
			request.Header.Set("Content-Type", "application/json")

			writer := httptest.NewRecorder()
			handlers.HandleGetObjects[models.Manager](writer, request)

			res := writer.Result()
			defer res.Body.Close()

			assert.Equal(t, tt.want.statusCode, res.StatusCode,
				"Expected status %d, got %d", tt.want.statusCode, res.StatusCode)

			if tt.want.contentType != "" {
				assert.Equal(t, tt.want.contentType, res.Header.Get("Content-Type"),
					"Expected Content-Type %s, got %s", tt.want.contentType, res.Header.Get("Content-Type"))
			}

			body, err := io.ReadAll(res.Body)
			defer res.Body.Close()
			if err != nil {
				t.Fatal(err)
			}

			if string(body) == "" {
				t.Errorf("Expected body, but result %s", string(body))
			}

			var news []models.Manager
			err = json.Unmarshal(body, &news)
			if err != nil {
				assert.NoError(t, err, "Response should be valid JSON")
			}
			if len(news) == 0 {
				t.Errorf("Expected length more than zero, got %d", len(news))
			}
			log.Printf("Count news: %d", len(news))
		})
	}
}
