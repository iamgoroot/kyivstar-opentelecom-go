package handlers

import "net/http"

func RegisterMultichannel(mux *http.ServeMux) {
	mux.HandleFunc("POST /rest/v1/messaging/multichannel", func(w http.ResponseWriter, r *http.Request) {
		writeJSON(w, map[string]string{
			"multiMsgId": "6badca00-2e05-42df-b7f1-4a5642e38af8",
		})
	})
	mux.HandleFunc("GET /rest/v1/messaging/multichannel/{multiMsgId}", func(w http.ResponseWriter, r *http.Request) {
		writeJSON(w, map[string]any{
			"date":       "2023-09-21T07:57:01Z",
			"multiMsgId": r.PathValue("multiMsgId"),
			"status":     "delivered",
			"bearerType": "viber",
			"reports": []map[string]string{
				{"bearerType": "sms", "date": "2023-09-21T07:55:32Z", "mid": "6badca00-2e05-42df-b7f1-4a5642e38af8", "state": "expired"},
				{"bearerType": "viber", "date": "2023-09-21T07:56:36Z", "mid": "d7c33333-749f-4027-a7bc-9ee66f48cad7", "state": "delivered"},
			},
		})
	})
}
