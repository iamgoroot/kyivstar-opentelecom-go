package handlers

import "net/http"

func RegisterProfile(mux *http.ServeMux) {
	mux.HandleFunc("POST /rest/v1/subscribers/profile", func(w http.ResponseWriter, r *http.Request) {
		writeJSON(w, map[string]any{
			"data": map[string]any{
				"profile": map[string]string{
					"gender": "MALE",
					"age":    "18-24",
				},
			},
			"dataPresent": true,
			"errors":      []string{},
			"extensions":  nil,
		})
	})
}
