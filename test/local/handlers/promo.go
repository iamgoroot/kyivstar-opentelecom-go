package handlers

import (
	"io"
	"net/http"
)

func RegisterPromo(mux *http.ServeMux) {
	mux.HandleFunc("POST /rest/v1/promo", func(w http.ResponseWriter, r *http.Request) {
		writeJSON(w, map[string]any{
			"id":                    "00000000-0000-0000-0000-000000000200",
			"authorUsername":        "user@user.com",
			"status":                "DRAFT",
			"startDate":             "2023-09-21T07:57:01",
			"endDate":               "2023-09-22T07:57:01",
			"textToSend":            "text",
			"nextAvailableStatuses": []string{"WAITING"},
			"messageContent": map[string]any{
				"messageParamCount": 0,
			},
		})
	})
	mux.HandleFunc("GET /rest/v1/promo", func(w http.ResponseWriter, r *http.Request) {
		writeJSON(w, map[string]any{
			"reqId":         "ad30594292f7959683a410bf1add088e",
			"totalPages":    1,
			"totalElements": 1,
			"number":        1,
			"size":          1,
			"promos": []map[string]any{
				{
					"id":             "00000000-0000-0000-0000-000000000200",
					"authorUsername": "user@user.com",
					"status":         "DRAFT",
				},
			},
		})
	})
	mux.HandleFunc("GET /rest/v1/promo/{promoUUID}", func(w http.ResponseWriter, r *http.Request) {
		writeJSON(w, map[string]string{
			"id":             r.PathValue("promoUUID"),
			"authorUsername": "user@user.com",
			"status":         "DRAFT",
		})
	})
	mux.HandleFunc("POST /rest/v1/promo/{promoUUID}/audience", func(w http.ResponseWriter, r *http.Request) {
		writeJSON(w, map[string]string{
			"reqId": "ad30594292f7959683a410bf1add088e",
			"name":  "audience-list-1",
		})
	})
	mux.HandleFunc("POST /rest/v1/promo/{promoUUID}/image", func(w http.ResponseWriter, r *http.Request) {
		file, _, err := r.FormFile("file")
		if err != nil {
			http.Error(w, "missing file", http.StatusBadRequest)
			return
		}

		_, _ = io.Copy(io.Discard, file)
		file.Close()

		writeJSON(w, map[string]any{
			"reqId":   "ad30594292f7959683a410bf1add088e",
			"success": true,
		})
	})
	mux.HandleFunc("PUT /rest/v1/promo/{promoUUID}/status/{status}", func(w http.ResponseWriter, r *http.Request) {
		writeJSON(w, map[string]string{
			"id":     r.PathValue("promoUUID"),
			"status": r.PathValue("status"),
		})
	})
	mux.HandleFunc("GET /rest/v1/promo/{promoUUID}/statistics", func(w http.ResponseWriter, r *http.Request) {
		writeJSON(w, map[string]any{
			"sentCount":                        10,
			"deliveriesCount":                  8,
			"unmatchedCount":                   0,
			"deliveriesPortionsCount":          0,
			"deliveriesUnmatchedPortionsCount": 0,
			"deliveriesInternalPortionsCount":  0,
			"deliveriesExternalPortionsCount":  0,
			"undeliveredCount":                 0,
			"unknownStatusCount":               0,
			"canceledByContactPolicyCount":     0,
			"seenCount":                        0,
			"blacklistedCount":                 0,
			"declinedCount":                    0,
			"expiredCount":                     0,
			"wasNotSent":                       20,
		})
	})
}
