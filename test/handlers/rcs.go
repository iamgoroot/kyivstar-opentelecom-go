package handlers

import "net/http"

func RegisterRCS(mux *http.ServeMux) {
	mux.HandleFunc("POST /rest/v1/rcs/text", func(w http.ResponseWriter, _ *http.Request) {
		writeJSON(w, map[string]string{
			"reqId": "ad30594292f7959683a410bf1add088e",
			"msgId": "20200000-0000-0000-0000-380670000200",
		})
	})
	mux.HandleFunc("POST /rest/v1/rcs/suggestion", func(w http.ResponseWriter, _ *http.Request) {
		writeJSON(w, map[string]string{
			"reqId": "ad30594292f7959683a410bf1add088e",
			"msgId": "20200000-0000-0000-0000-380670000200",
		})
	})
	mux.HandleFunc("POST /rest/v1/rcs/richcard", func(w http.ResponseWriter, _ *http.Request) {
		writeJSON(w, map[string]string{
			"reqId": "ad30594292f7959683a410bf1add088e",
			"msgId": "20200000-0000-0000-0000-380670000200",
		})
	})
	mux.HandleFunc("GET /rest/v1/rcs/{msgId}", func(w http.ResponseWriter, r *http.Request) {
		writeJSON(w, map[string]string{
			"reqId":  "ad30594292f7959683a410bf1add088e",
			"msgId":  r.PathValue("msgId"),
			"status": "delivered",
			"date":   "2025-01-01T12:00:00.000Z",
		})
	})
}
