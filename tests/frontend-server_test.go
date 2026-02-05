package tests

import (
	"finalwork/internal/db"
	"finalwork/internal/handlers"
	"finalwork/internal/models"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNewsPageHandlers(t *testing.T) {
	type want struct {
		statusCode   int
		contentType  string
		url          string
		templatePage string
	}

	tests := []struct {
		name string
		want want
	}{
		{
			name: "Test Index Page",
			want: want{
				statusCode:   200,
				url:          "/",
				templatePage: "../web/templates/index.html",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := httptest.NewServer(handlers.ShowTemplatePageGeneric[models.News](tt.want.templatePage, http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
				if req.URL.Path != tt.want.url {
					t.Errorf("Expected path '%s', got '%s'", tt.want.url, req.URL.Path)
				}
			}), db.SelectAll))

			defer server.Close()

			resp, err := http.Get(server.URL + tt.want.url)
			if err != nil {
				t.Fatal(err)
			}

			if resp.StatusCode != http.StatusOK {
				t.Errorf("Expected status OK, got %d", resp.StatusCode)
			}

			body, err := io.ReadAll(resp.Body)
			defer resp.Body.Close()
			if err != nil {
				t.Fatal(err)
			}

			if string(body) == "" {
				t.Errorf("Expected body, but result %s", string(body))
			}
		})
	}
}

func TestManagersPageHandlers(t *testing.T) {
	type want struct {
		statusCode   int
		contentType  string
		url          string
		templatePage string
	}

	tests := []struct {
		name string
		want want
	}{
		{
			name: "Test Contacts Page",
			want: want{
				statusCode:   200,
				url:          "/contacts",
				templatePage: "../web/templates/contacts.html",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := httptest.NewServer(handlers.ShowTemplatePageGeneric[models.Manager](tt.want.templatePage, http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
				if req.URL.Path != tt.want.url {
					t.Errorf("Expected path '%s', got '%s'", tt.want.url, req.URL.Path)
				}
			}), db.SelectAll))

			defer server.Close()

			resp, err := http.Get(server.URL + tt.want.url)
			if err != nil {
				t.Fatal(err)
			}

			if resp.StatusCode != http.StatusOK {
				t.Errorf("Expected status OK, got %d", resp.StatusCode)
			}

			body, err := io.ReadAll(resp.Body)
			defer resp.Body.Close()
			if err != nil {
				t.Fatal(err)
			}

			if string(body) == "" {
				t.Errorf("Expected body, but result %s", string(body))
			}
		})
	}
}

func TestTemplatePageHandlers(t *testing.T) {
	type want struct {
		statusCode   int
		contentType  string
		url          string
		templatePage string
	}

	tests := []struct {
		name string
		want want
	}{
		{
			name: "Test Services Page",
			want: want{
				statusCode:   200,
				url:          "/services",
				templatePage: "../web/templates/services.html",
			},
		},
		{
			name: "Test SustainableDevelopment Page",
			want: want{
				statusCode:   200,
				url:          "/sustainableDevelopment",
				templatePage: "../web/templates/sustainableDevelopment.html",
			},
		},
		{
			name: "Test Login Page",
			want: want{
				statusCode:   200,
				url:          "/login",
				templatePage: "../web/templates/admin/auth.html",
			},
		},
		{
			name: "Test images Page",
			want: want{
				statusCode:   200,
				url:          "/images",
				templatePage: "../web/templates/admin/images.html",
			},
		},
		{
			name: "Test images Page",
			want: want{
				statusCode:   200,
				url:          "/newsAdd",
				templatePage: "../web/templates/admin/newsAdd.html",
			},
		},
		{
			name: "Test newsBrowser Page",
			want: want{
				statusCode:   200,
				url:          "/newsBrowser",
				templatePage: "../web/templates/admin/newsBrowser.html",
			},
		},
		{
			name: "Test managerAdd Page",
			want: want{
				statusCode:   200,
				url:          "/managerAdd",
				templatePage: "../web/templates/admin/managerAdd.html",
			},
		},
		{
			name: "Test managerBrowser Page",
			want: want{
				statusCode:   200,
				url:          "/managerBrowser",
				templatePage: "../web/templates/admin/managerBrowser.html",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := httptest.NewServer(handlers.ShowTemplatePage(tt.want.templatePage, http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
				if req.URL.Path != tt.want.url {
					t.Errorf("Expected path '%s', got '%s'", tt.want.url, req.URL.Path)
				}
			})))

			defer server.Close()

			resp, err := http.Get(server.URL + tt.want.url)
			if err != nil {
				t.Fatal(err)
			}

			if resp.StatusCode != http.StatusOK {
				t.Errorf("Expected status OK, got %d", resp.StatusCode)
			}

			body, err := io.ReadAll(resp.Body)
			defer resp.Body.Close()
			if err != nil {
				t.Fatal(err)
			}

			if string(body) == "" {
				t.Errorf("Expected body, but result %s", string(body))
			}
		})
	}
}

func TestTemplatePageParamsHandlers(t *testing.T) {
	type want struct {
		statusCode   int
		contentType  string
		url          string
		templatePage string
	}

	tests := []struct {
		name string
		want want
	}{
		{
			name: "Test News Page",
			want: want{
				statusCode:   200,
				url:          "/news",
				templatePage: "../web/templates/news.html",
			},
		},
		{
			name: "Test News/1  Page",
			want: want{
				statusCode:   200,
				url:          "/news/1",
				templatePage: "../web/templates/news.html",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := httptest.NewServer(handlers.ShowTemplatePageParams(tt.want.templatePage, http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
				if req.URL.Path != tt.want.url {
					t.Errorf("Expected path '%s', got '%s'", tt.want.url, req.URL.Path)
				}
			}), handlers.SelectDataNews))

			defer server.Close()

			resp, err := http.Get(server.URL + tt.want.url)
			if err != nil {
				t.Fatal(err)
			}

			if resp.StatusCode != http.StatusOK {
				t.Errorf("Expected status OK, got %d", resp.StatusCode)
			}

			body, err := io.ReadAll(resp.Body)
			defer resp.Body.Close()
			if err != nil {
				t.Fatal(err)
			}

			if string(body) == "" {
				t.Errorf("Expected body, but result %s", string(body))
			}
		})
	}
}
