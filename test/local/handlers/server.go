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

func addGatewayHeaders(w http.ResponseWriter) {
	w.Header().Set("X-Request-Id", "d06d47e31cffcd60f3fd38d9718c9026")
	w.Header().Set("X-Rate-Limit-Limit", "1000")
	w.Header().Set("X-Rate-Limit-Period-Sec", "1")
	w.Header().Set("X-Rate-Limit-Remaining", "999")
	w.Header().Set("X-Rate-Limit-Reset", "1")
	w.Header().Set("X-Reserved-Tarification-Units", "1")
}

func writeJSONStatus(w http.ResponseWriter, status int, v any) {
	addGatewayHeaders(w)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(v)
}

func writeJSON(w http.ResponseWriter, v any) {
	writeJSONStatus(w, http.StatusOK, v)
}
