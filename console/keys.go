package main

import (
	"encoding/json"
	"net/http"
	"strings"
	"sync"
	"time"
)

type keyCache struct {
	mu      sync.RWMutex
	keys    map[string]bool
	expires time.Time
}

var apiKeyCache = &keyCache{
	keys: make(map[string]bool),
}

type APIKey struct {
	ID        string `json:"id,omitempty"`
	Key       string `json:"key"`
	Name      string `json:"name"`
	Active    bool   `json:"active"`
	CreatedAt string `json:"created_at,omitempty"`
}

func isValidAPIKey(key string) bool {
	apiKeyCache.mu.RLock()
	if time.Now().Before(apiKeyCache.expires) {
		valid := apiKeyCache.keys[key]
		apiKeyCache.mu.RUnlock()
		return valid
	}
	apiKeyCache.mu.RUnlock()

	return checkKeyInSupabase(key)
}

func checkKeyInSupabase(key string) bool {
	data, status, err := supabaseServiceRequest("GET", "api_keys", "key=eq."+key+"&active=eq.true&select=key", nil)
	if err != nil || status != 200 {
		return false
	}
	var results []map[string]interface{}
	if err := json.Unmarshal(data, &results); err != nil {
		return false
	}
	found := len(results) > 0

	if found {
		apiKeyCache.mu.Lock()
		apiKeyCache.keys[key] = true
		if apiKeyCache.expires.IsZero() {
			apiKeyCache.expires = time.Now().Add(60 * time.Second)
		}
		apiKeyCache.mu.Unlock()
	}

	return found
}

func refreshKeyCache() {
	data, status, err := supabaseServiceRequest("GET", "api_keys", "active=eq.true&select=key", nil)
	if err != nil || status != 200 {
		return
	}
	var results []map[string]string
	if err := json.Unmarshal(data, &results); err != nil {
		return
	}
	apiKeyCache.mu.Lock()
	apiKeyCache.keys = make(map[string]bool)
	for _, r := range results {
		apiKeyCache.keys[r["key"]] = true
	}
	apiKeyCache.expires = time.Now().Add(60 * time.Second)
	apiKeyCache.mu.Unlock()
}

func authHeader(r *http.Request) string {
	auth := r.Header.Get("Authorization")
	if strings.HasPrefix(auth, "Bearer ") {
		return strings.TrimPrefix(auth, "Bearer ")
	}
	return r.Header.Get("x-api-key")
}
