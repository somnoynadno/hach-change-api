package api

import (
	"github.com/gorilla/mux"
	"hack-change-api/controller/auth"
	"net/http"
)

func InitAuth(router *mux.Router) {
	router.HandleFunc("/login", auth.Login).Methods(http.MethodPost, http.MethodOptions)
	router.HandleFunc("/register", auth.Register).Methods(http.MethodPost, http.MethodOptions)
}
