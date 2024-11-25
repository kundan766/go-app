package handlers

import (
	"go-app/utils"
	"net/http"
)

// RedirectHandler handles the redirect logic
func RedirectHandler(w http.ResponseWriter, r *http.Request) {
	// Check for common proxy tool User-Agents
	userAgent := r.Header.Get("User -Agent")
	if utils.IsProxyTool(userAgent) {
		// Redirect to Yahoo if a proxy tool is detected
		http.Redirect(w, r, "https://www.yahoo.com", http.StatusFound)
		return
	}

	// Check for a custom header that might indicate a proxy
	if r.Header.Get("X-Proxy-Detection") != "" {
		http.Redirect(w, r, "https://www.yahoo.com", http.StatusFound)
		return
	}

	// Redirect to Google if the connection is clean
	http.Redirect(w, r, "https://www.google.com", http.StatusFound)
}
