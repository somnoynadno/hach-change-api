package muxutil

import (
	"encoding/json"
	"hack-change-api/db"
	"hack-change-api/metrics"
	"net/http"
)

// SetTotalCountHeader adds important headers for React Admin panel.
func SetTotalCountHeader(w http.ResponseWriter, count string) {
	w.Header().Add("Access-Control-Expose-Headers", "X-Total-Count")
	w.Header().Add("X-Total-Count", count)
}

// SetTotalCountHeader verifies for necessary parameters existence.
func CheckOrderAndSortParams(order *string, sort *string) {
	if *order != "ASC" && *order != "DESC" {
		*order = "ASC"
	}
	if *sort == "" {
		*sort = "ID"
	}
}

func Message(status bool, message string) map[string]interface{} {
	return map[string]interface{}{"status": status, "message": message}
}

func Respond(w http.ResponseWriter, data map[string]interface{}) {
	metrics.PromResponseOK.Inc()
	_ = json.NewEncoder(w).Encode(data)
}

func RespondJSON(w http.ResponseWriter, data []byte) {
	metrics.PromResponseOK.Inc()
	_, _ = w.Write(data)
}

var HandleOptions = func(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

var HandlePing = func(w http.ResponseWriter, r *http.Request) {
	err := db.GetDB().DB().Ping()
	if err != nil {
		HandleInternalError(w, err)
		return
	}

	_ = json.NewEncoder(w).Encode(Message(true, "pong"))
}

