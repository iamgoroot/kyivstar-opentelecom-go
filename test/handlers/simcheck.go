package handlers

import "net/http"

func RegisterSimCheck(mux *http.ServeMux) {
	mux.HandleFunc("GET /rest/v1/subscribers/{phoneNumber}/sim-check-antifraud", func(w http.ResponseWriter, _ *http.Request) {
		writeJSON(w, map[string]any{
			"reqId": "ad30594292f7959683a410bf1add088e",
			"cid":   "4d3e4995-fa0c-4a44-91e5-c71d7d447700",
			"resource": map[string]bool{
				"simChanged":     false,
				"simTypeChanged": false,
				"geoChanged":     false,
				"imeiChanged":    false,
				"callForwarding": false,
			},
		})
	})
}
