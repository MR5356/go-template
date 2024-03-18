package ginutil

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"testing"
)

func setupRoute() *gin.Engine {
	r := gin.Default()

	r.GET("/id", func(c *gin.Context) {
		c.String(http.StatusOK, GetId(c))
	})

	r.GET("/token", func(c *gin.Context) {
		c.String(http.StatusOK, GetToken(c))
	})

	return r
}

func TestGetId(t *testing.T) {
	r := setupRoute()

	tests := []struct {
		name string
		id   string
	}{
		{
			name: "test-id",
			id:   "test-id",
		},
		{
			name: "test-id-nil",
			id:   "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rr := httptest.NewRecorder()

			req, err := http.NewRequest("GET", "/id?id="+tt.id, nil)
			if err != nil {
				t.Fatal(err)
			}

			r.ServeHTTP(rr, req)

			if rr.Code != http.StatusOK {
				t.Errorf("GetId() = %v, want %v", rr.Code, http.StatusOK)
			}

			if rr.Body.String() != tt.id {
				t.Errorf("GetId() = %v, want %v", rr.Body.String(), tt.id)
			}
		})
	}
}

func TestGetToken(t *testing.T) {
	r := setupRoute()

	tests := []struct {
		name   string
		token  string
		header string
		cookie string
	}{
		{
			name:   "test-cookie",
			token:  "test-token",
			cookie: "test-token",
		},
		{
			name:   "test-header",
			token:  "test-token",
			header: "test-token",
		},
		{
			name:  "test-nil",
			token: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rr := httptest.NewRecorder()

			req, err := http.NewRequest("GET", "/token", nil)
			if err != nil {
				t.Fatal(err)
			}

			if len(tt.cookie) > 0 {
				req.AddCookie(&http.Cookie{
					Name:  "token",
					Value: tt.cookie,
				})
			}

			if len(tt.header) > 0 {
				req.Header.Set("Authorization", tt.header)
			}

			r.ServeHTTP(rr, req)

			if rr.Code != http.StatusOK {
				t.Errorf("GetToken() = %v, want %v", rr.Code, http.StatusOK)
			}

			if rr.Body.String() != tt.token {
				t.Errorf("GetToken() = %v, want %v", rr.Body.String(), tt.token)
			}
		})
	}
}
