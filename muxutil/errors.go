package muxutil

import (
	log "github.com/sirupsen/logrus"
	"hack-change-api/metrics"
	"net/http"
)

// HandleBadRequest must be raised on 400 errors.
func HandleBadRequest(w http.ResponseWriter, err error) {
	log.Error(err)
	w.WriteHeader(http.StatusBadRequest)
	metrics.PromResponseBadRequest.Inc()
	Respond(w, Message(false, err.Error()))
}

// HandleUnauthorized must be raised on 401 errors.
func HandleUnauthorized(w http.ResponseWriter, err error) {
	log.Warn(err)
	w.WriteHeader(http.StatusUnauthorized)
	metrics.PromResponseUnauthorized.Inc()
	Respond(w, Message(false, err.Error()))
}

// HandleForbidden must be raised on 403 errors.
func HandleForbidden(w http.ResponseWriter, err error) {
	log.Warn(err)
	w.WriteHeader(http.StatusForbidden)
	metrics.PromResponseForbidden.Inc()
	Respond(w, Message(false, err.Error()))
}

// HandleNotFound must be raised on 404 errors.
func HandleNotFound(w http.ResponseWriter) {
	log.Info("Not found")
	w.WriteHeader(http.StatusNotFound)
	metrics.PromResponseNotFound.Inc()
	Respond(w, Message(false, "not found"))
}

// HandleInternalError must be raised on 500 errors.
func HandleInternalError(w http.ResponseWriter, err error) {
	log.Error(err)
	w.WriteHeader(http.StatusInternalServerError)
	metrics.PromResponseInternalError.Inc()
	Respond(w, Message(false, err.Error()))
}

