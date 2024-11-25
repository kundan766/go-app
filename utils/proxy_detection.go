package utils

import "strings"

func IsProxyTool(userAgent string) bool {
	proxyTools := []string{
		"Fiddler",
		"Burp Suite",
		"Charles",
		"Proxyman",
		"Reqable",
	}

	for _, tool := range proxyTools {
		if strings.Contains(userAgent, tool) {
			return true
		}
	}
	return false
}
