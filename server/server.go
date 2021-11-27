package server

import (
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	log "github.com/sirupsen/logrus"
	u "hack-change-api/muxutil"
	"hack-change-api/server/middleware"
	"net/http"
	"os"
)

func initRouter() *mux.Router {
	r := mux.NewRouter()

	router := r.PathPrefix("/api").Subrouter()
	v1 := router.PathPrefix("/v1").Subrouter()
	//auth := router.PathPrefix("/auth").Subrouter()

	// middleware usage
	// do NOT modify the order
	router.Use(middleware.CORS)    // enable CORS headers
	router.Use(middleware.LogPath) // log IP, path and method
	router.Use(middleware.LogBody) // log HTTP body
	v1.Use(middleware.JwtAuth)     // attach JWT auth middleware

	// handle ping
	router.HandleFunc("/ping", u.HandlePing).Methods(http.MethodGet, http.MethodPost)

	// set up prometheus
	r.Handle("/metrics", promhttp.Handler())

	return r
}

func RunForever() {
	r := initRouter()

	port := os.Getenv("PORT")
	if port == "" {
		port = "9898" // localhost
	}

	log.Info("listening on: ", port)
	err := http.ListenAndServe(":"+port, r)
	if err != nil {
		log.Panic(err)
	}
}

