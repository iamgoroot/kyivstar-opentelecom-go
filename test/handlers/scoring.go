package handlers

import "net/http"

func RegisterScoring(mux *http.ServeMux) {
	mux.HandleFunc("GET /rest/v1/subscribers/{phoneNumber}/scoring", func(w http.ResponseWriter, _ *http.Request) {
		writeJSON(w, map[string]any{
			"reqId": "bd4553a32a7e238e33fa8245be4d3f6d",
			"cid":   "4ca11176-f947-4757-a1a8-09031d60217c",
			"resource": map[string]float64{
				"scoreBal": 0.08457,
			},
		})
	})
}
