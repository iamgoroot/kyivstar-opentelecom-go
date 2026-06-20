package handlers

import (
	"net/http"
)

func RegisterSMS(mux *http.ServeMux) {
	mux.HandleFunc("POST /rest/v1/sms", func(w http.ResponseWriter, _ *http.Request) {
		writeJSON(w, map[string]any{
			"reqId":               "ad30594292f7959683a410bf1add088e",
			"msgId":               "20200000-0000-0000-0000-380670000200",
			"reservedSmsSegments": 1,
		})
	})
	mux.HandleFunc("GET /rest/v1/sms/{msgId}", func(w http.ResponseWriter, r *http.Request) {
		writeJSON(w, map[string]string{
			"reqId":  "ad30594292f7959683a410bf1add088e",
			"msgId":  r.PathValue("msgId"),
			"status": "delivered",
			"date":   "2025-01-01T12:00:00.000Z",
		})
	})
	mux.HandleFunc("POST /rest/v1/sms/batch", func(w http.ResponseWriter, _ *http.Request) {
		writeJSON(w, map[string]any{
			"reqId": "f57b32f728056be66cfdb1eee74ea1ac",
			"data": map[string]any{
				"uniqueMsgKey1": map[string]any{
					"msgId":               "20200000-0000-0000-0000-380670000200",
					"reservedSmsSegments": 1,
				},
				"uniqueMsgKey2": map[string]any{
					"msgId":               "20200000-0000-0000-0000-380670000201",
					"reservedSmsSegments": 1,
				},
			},
		})
	})
	mux.HandleFunc("POST /rest/v1/sms/status/batch", func(w http.ResponseWriter, _ *http.Request) {
		writeJSON(w, map[string]any{
			"reqId": "f57b32f728056be66cfdb1eee74ea1ac",
			"data": map[string]any{
				"20200000-0000-0000-0000-380670000200": map[string]string{
					"msgId":  "20200000-0000-0000-0000-380670000200",
					"status": "delivered",
					"date":   "2025-01-01T12:00:00.000Z",
				},
			},
		})
	})
}

func RegisterSMSErrors(mux *http.ServeMux) {
	mux.HandleFunc("POST /rest/v1/sms", func(w http.ResponseWriter, _ *http.Request) {
		writeJSONStatus(w, http.StatusBadRequest, map[string]any{
			"reqId":     "err-req-id",
			"errorCode": 40001,
			"errorMsg":  "Invalid phone number format",
		})
	})
	mux.HandleFunc("GET /rest/v1/sms/{msgId}", func(w http.ResponseWriter, _ *http.Request) {
		writeJSONStatus(w, http.StatusNotFound, map[string]any{
			"reqId":     "err-req-id",
			"errorCode": 40401,
			"errorMsg":  "Message not found",
		})
	})
	mux.HandleFunc("POST /rest/v1/sms/batch", func(w http.ResponseWriter, _ *http.Request) {
		writeJSONStatus(w, http.StatusUnauthorized, map[string]any{
			"reqId":     "err-req-id",
			"errorCode": 40101,
			"errorMsg":  "Invalid credentials",
		})
	})
	mux.HandleFunc("POST /rest/v1/sms/status/batch", func(w http.ResponseWriter, _ *http.Request) {
		writeJSONStatus(w, http.StatusInternalServerError, map[string]any{
			"reqId":     "err-req-id",
			"errorCode": 50001,
			"errorMsg":  "Internal server error",
		})
	})
}
