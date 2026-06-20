package handlers

import "net/http"

func RegisterLifetime(mux *http.ServeMux) {
	mux.HandleFunc("GET /rest/v1/subscribers/{phoneNumber}/lifetime", func(w http.ResponseWriter, _ *http.Request) {
		writeJSON(w, map[string]any{
			"reqId": "3388babfc3db5a4941416248b8c0008f",
			"cid":   "4d3e4995-fa0c-4a44-91e5-c71d7d447700",
			"resource": map[string]any{
				"lifetimeDuration": map[string]any{
					"from":     0,
					"to":       1,
					"timeUnit": "MONTHS",
				},
			},
		})
	})
}
