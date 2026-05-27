package handlers

import "net/http"

func RegisterDeviceCheck(mux *http.ServeMux) {
	mux.HandleFunc("GET /rest/v1/subscribers/{phoneNumber}/device-check", func(w http.ResponseWriter, r *http.Request) {
		writeJSON(w, map[string]any{
			"reqId": "ad30594292f7959683a410bf1add088e",
			"cid":   "4d3e4995-fa0c-4a44-91e5-c71d7d447700",
			"resource": map[string]any{
				"imeiRes":   "COMPLETELY_MATCHED",
				"imeiCount": 1,
			},
		})
	})
}
