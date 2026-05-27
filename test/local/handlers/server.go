package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
)

func NewServer(registers ...func(*http.ServeMux)) *httptest.Server {
	mux := http.NewServeMux()
	for _, r := range registers {
		r(mux)
	}

	return httptest.NewServer(mux)
}

func writeJSON(w http.ResponseWriter, v any) {
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(v)
}
