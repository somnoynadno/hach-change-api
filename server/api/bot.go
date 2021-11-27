package api

import (
	"github.com/gorilla/mux"
	"hack-change-api/controller/bot"
	"net/http"
)

func InitBot(router *mux.Router) {
	router.HandleFunc("/balaboba", bot.TalkToBalaboba).Methods(http.MethodGet, http.MethodOptions)
}

