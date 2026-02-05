package tests

import (
	"bytes"
	"encoding/json"
	"finalwork/internal/handlers"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAuthUser(t *testing.T) {
	type want struct {
		statusCode       int
		contentType      string
		responseValid    bool
		responseContains string
	}

	type requestBody struct {
		login    string
		password string
	}

	tests := []struct {
		name string
		body requestBody
		want want
	}{
		{
			name: "Test Auth User with valid input",
			body: requestBody{
				login:    "admin",
				password: "admin",
			},
			want: want{
				statusCode:       200, // Ожидаем Created
				contentType:      "application/json; charset=utf-8",
				responseValid:    false,
				responseContains: "ok",
			},
		},
		{
			name: "Test Auth User with invalid password",
			body: requestBody{
				login:    "admin",
				password: "1234",
			},
			want: want{
				statusCode:       401,
				responseValid:    false,
				responseContains: "Invalid username or password",
			},
		},
		{
			name: "Test Auth User with not valid user",
			body: requestBody{
				login:    "test",
				password: "1234",
			},
			want: want{
				statusCode:       401,
				responseValid:    false,
				responseContains: "Invalid username or password",
			},
		},
		{
			name: "Test Auth User with invalid request body",
			body: requestBody{},
			want: want{
				statusCode:       400,
				responseValid:    false,
				responseContains: "Invalid JSON",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var bodyBytes []byte
			if tt.body.login != "" {
				jsonBody := map[string]string{
					"login":    tt.body.login,
					"password": tt.body.password,
				}
				var err error
				bodyBytes, err = json.Marshal(jsonBody)
				require.NoError(t, err)
			} else {
				bodyBytes = []byte(`{invalid json`)
			}

			request := httptest.NewRequest(http.MethodPost, "/login",
				bytes.NewReader(bodyBytes))
			request.Header.Set("Content-Type", "application/json")

			writer := httptest.NewRecorder()

			handlers.HandleLogin(writer, request)

			res := writer.Result()
			defer res.Body.Close()

			assert.Equal(t, tt.want.statusCode, res.StatusCode,
				"Expected status %d, got %d", tt.want.statusCode, res.StatusCode)

			if tt.want.contentType != "" {
				assert.Equal(t, tt.want.contentType, res.Header.Get("Content-Type"),
					"Expected Content-Type %s, got %s", tt.want.contentType, res.Header.Get("Content-Type"))
			}

			resBody, err := io.ReadAll(res.Body)
			log.Printf("Request Body: %s", string(resBody))
			require.NoError(t, err)

			if tt.want.responseValid {
				var response map[string]interface{}
				err = json.Unmarshal(resBody, &response)
				assert.NoError(t, err, "Response should be valid JSON")
			}
			assert.Contains(t, string(resBody), tt.want.responseContains,
				"Response should contain email field")
		})
	}
}
