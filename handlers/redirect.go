package handlers

import (
	"go-app/utils"
	"net/http"
)

func RedirectHandler(w http.ResponseWriter, r *http.Request) {

	userAgent := r.Header.Get("User-Agent")
	if utils.IsProxyTool(userAgent) {
		// Redirect to Yahoo if a proxy tool is detected
		http.Redirect(w, r, "https://www.yahoo.com", http.StatusFound)
		return
	}

	if r.Header.Get("X-Proxy-Detection") != "" {
		http.Redirect(w, r, "https://www.yahoo.com", http.StatusFound)
		return
	}

	// Redirect to Google if the connection is clean
	http.Redirect(w, r, "https://www.google.com", http.StatusFound)
}
