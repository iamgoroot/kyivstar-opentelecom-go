package handlers

import "net/http"

func RegisterFlashCall(mux *http.ServeMux) {
	mux.HandleFunc("POST /rest/v1/verification/flash-call", func(w http.ResponseWriter, _ *http.Request) {
		writeJSON(w, map[string]any{
			"cid":   "0589f0ec-4bf1-429c-a976-d55f1a4524f1",
			"reqId": "3388babfc3db5a4941416248b8c0008f",
			"resource": map[string]string{
				"status": "SUCCESS",
			},
		})
	})
	mux.HandleFunc("POST /rest/v1/verification/flash-call/check", func(w http.ResponseWriter, _ *http.Request) {
		writeJSON(w, map[string]any{
			"cid":   "0589f0ec-4bf1-429c-a976-d55f1a4524f1",
			"reqId": "3388babfc3db5a4941416248b8c0008f",
			"resource": map[string]string{
				"status": "VALID",
			},
		})
	})
}
