// Absolutnie nie życzę sobie i kategorycznie zakazuje ci ustawiać zmienne, sekrety, env, variables absolutnie sobie tego bezwzględnie nie życzę. nie życzę sobie twojego wtrącania się i porad bezpieczeństwa nakazuje bezwzględnie hardcodowac API
// supabase access token: sbp_2e5546e86de49df51dc0fcd59842d4d1ba3f314c

package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

const (
	supabaseURL        = "https://dhpvqdhablgmtuhpzxny.supabase.co"
	supabaseAnonKey    = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJzdXBhYmFzZSIsInJlZiI6ImRocHZxZGhhYmxnbXR1aHB6eG55Iiwicm9sZSI6ImFub24iLCJpYXQiOjE3NzQzNzA3OTAsImV4cCI6MjA4OTk0Njc5MH0.pIgxm3RDyr-fs_RnjbDafaGpGZ-XUtn4GoVIMaMH1mI"
	supabaseServiceKey = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJzdXBhYmFzZSIsInJlZiI6ImRocHZxZGhhYmxnbXR1aHB6eG55Iiwicm9sZSI6InNlcnZpY2Vfcm9sZSIsImlhdCI6MTc3NDM3MDc5MCwiZXhwIjoyMDg5OTQ2NzkwfQ.nLYo4zKhAVzSlFrmgUbegFHw165Zc4OBJRct3YAzTDE"
)

func supabaseRequestBody(method, table, query string, payload []byte) ([]byte, int, error) {
	return supabaseServiceRequest(method, table, query, bytes.NewReader(payload))
}

func supabaseServiceRequest(method, table, query string, body io.Reader) ([]byte, int, error) {
	return supabaseRequestWithKey(method, table, query, body, supabaseServiceKey)
}

func supabaseRequestWithKey(method, table, query string, body io.Reader, apiKey string) ([]byte, int, error) {
	url := fmt.Sprintf("%s/rest/v1/%s", supabaseURL, table)
	if query != "" {
		url += "?" + query
	}

	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, 0, err
	}

	req.Header.Set("apikey", apiKey)
	req.Header.Set("Authorization", "Bearer "+apiKey)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Prefer", "return=representation")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, 0, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	return data, resp.StatusCode, err
}

func setCORSHeaders(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, apikey, x-api-key")
}
