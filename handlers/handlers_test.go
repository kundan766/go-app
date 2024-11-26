package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRedirectHandler(t *testing.T) {
	tests := []struct {
		name         string
		userAgent    string
		customHeader string
		expectedURL  string
	}{
		{"Normal Request", "", "", "https://www.google.com"},
		{"Fiddler User-Agent", "Fiddler", "", "https://www.yahoo.com"},
		{"Burp Suite User-Agent", "Burp Suite", "", "https://www.yahoo.com"},
		{"Custom Proxy Header", "", "X-Proxy-Detection: true", "https://www.yahoo.com"},
		{"X-Forwarded-For Header", "", "X-Forwarded-For: 192.168.1.1", "https://www.yahoo.com"},
		{"Malformed User-Agent", "InvalidUser Agent", "", "https://www.google.com"},
		{"Multiple Headers", "Fiddler", "X-Proxy-Detection: true", "https://www.yahoo.com"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest("GET", "http://localhost:8080", nil)
			if tt.userAgent != "" {
				req.Header.Set("User-Agent", tt.userAgent)
			}
			if tt.customHeader != "" {
				req.Header.Set("X-Proxy-Detection", "true")
			}

			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(RedirectHandler)

			handler.ServeHTTP(rr, req)

			if status := rr.Code; status != http.StatusFound {
				t.Errorf("handler returned wrong status code: got %v want %v",
					status, http.StatusFound)
			}

			if location := rr.Header().Get("Location"); location != tt.expectedURL {
				t.Errorf("handler returned wrong redirect location: got %v want %v",
					location, tt.expectedURL)
			}
		})
	}
}
