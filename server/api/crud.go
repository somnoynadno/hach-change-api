package api

import (
	"github.com/gorilla/mux"
	"hack-change-api/controller/crud"
	"net/http"
)

func InitCRUD(router *mux.Router) {
	router.HandleFunc("/blog_post", crud.BlogPostQuery).Methods(http.MethodGet, http.MethodOptions)
	router.HandleFunc("/blog_post", crud.BlogPostCreate).Methods(http.MethodPost, http.MethodOptions)
	router.HandleFunc("/blog_post/{id}", crud.BlogPostRetrieve).Methods(http.MethodGet, http.MethodOptions)
	router.HandleFunc("/blog_post/{id}", crud.BlogPostUpdate).Methods(http.MethodPut, http.MethodOptions)
	router.HandleFunc("/blog_post/{id}", crud.BlogPostDelete).Methods(http.MethodDelete, http.MethodOptions)

	router.HandleFunc("/chat_message", crud.ChatMessageQuery).Methods(http.MethodGet, http.MethodOptions)
	router.HandleFunc("/chat_message", crud.ChatMessageCreate).Methods(http.MethodPost, http.MethodOptions)
	router.HandleFunc("/chat_message/{id}", crud.ChatMessageRetrieve).Methods(http.MethodGet, http.MethodOptions)
	router.HandleFunc("/chat_message/{id}", crud.ChatMessageUpdate).Methods(http.MethodPut, http.MethodOptions)
	router.HandleFunc("/chat_message/{id}", crud.ChatMessageDelete).Methods(http.MethodDelete, http.MethodOptions)

	router.HandleFunc("/comment", crud.CommentQuery).Methods(http.MethodGet, http.MethodOptions)
	router.HandleFunc("/comment", crud.CommentCreate).Methods(http.MethodPost, http.MethodOptions)
	router.HandleFunc("/comment/{id}", crud.CommentRetrieve).Methods(http.MethodGet, http.MethodOptions)
	router.HandleFunc("/comment/{id}", crud.CommentUpdate).Methods(http.MethodPut, http.MethodOptions)
	router.HandleFunc("/comment/{id}", crud.CommentDelete).Methods(http.MethodDelete, http.MethodOptions)

	router.HandleFunc("/financial_instrument", crud.FinancialInstrumentQuery).Methods(http.MethodGet, http.MethodOptions)
	router.HandleFunc("/financial_instrument", crud.FinancialInstrumentCreate).Methods(http.MethodPost, http.MethodOptions)
	router.HandleFunc("/financial_instrument/{id}", crud.FinancialInstrumentRetrieve).Methods(http.MethodGet, http.MethodOptions)
	router.HandleFunc("/financial_instrument/{id}", crud.FinancialInstrumentUpdate).Methods(http.MethodPut, http.MethodOptions)
	router.HandleFunc("/financial_instrument/{id}", crud.FinancialInstrumentDelete).Methods(http.MethodDelete, http.MethodOptions)

	router.HandleFunc("/instrument_type", crud.InstrumentTypeQuery).Methods(http.MethodGet, http.MethodOptions)
	router.HandleFunc("/instrument_type", crud.InstrumentTypeCreate).Methods(http.MethodPost, http.MethodOptions)
	router.HandleFunc("/instrument_type/{id}", crud.InstrumentTypeRetrieve).Methods(http.MethodGet, http.MethodOptions)
	router.HandleFunc("/instrument_type/{id}", crud.InstrumentTypeUpdate).Methods(http.MethodPut, http.MethodOptions)
	router.HandleFunc("/instrument_type/{id}", crud.InstrumentTypeDelete).Methods(http.MethodDelete, http.MethodOptions)

	router.HandleFunc("/subscription", crud.SubscriptionQuery).Methods(http.MethodGet, http.MethodOptions)
	router.HandleFunc("/subscription", crud.SubscriptionCreate).Methods(http.MethodPost, http.MethodOptions)
	router.HandleFunc("/subscription/{id}", crud.SubscriptionRetrieve).Methods(http.MethodGet, http.MethodOptions)
	router.HandleFunc("/subscription/{id}", crud.SubscriptionUpdate).Methods(http.MethodPut, http.MethodOptions)
	router.HandleFunc("/subscription/{id}", crud.SubscriptionDelete).Methods(http.MethodDelete, http.MethodOptions)

	router.HandleFunc("/thread_comment", crud.ThreadCommentQuery).Methods(http.MethodGet, http.MethodOptions)
	router.HandleFunc("/thread_comment", crud.ThreadCommentCreate).Methods(http.MethodPost, http.MethodOptions)
	router.HandleFunc("/thread_comment/{id}", crud.ThreadCommentRetrieve).Methods(http.MethodGet, http.MethodOptions)
	router.HandleFunc("/thread_comment/{id}", crud.ThreadCommentUpdate).Methods(http.MethodPut, http.MethodOptions)
	router.HandleFunc("/thread_comment/{id}", crud.ThreadCommentDelete).Methods(http.MethodDelete, http.MethodOptions)

	router.HandleFunc("/user", crud.UserQuery).Methods(http.MethodGet, http.MethodOptions)
	router.HandleFunc("/user", crud.UserCreate).Methods(http.MethodPost, http.MethodOptions)
	router.HandleFunc("/user/{id}", crud.UserRetrieve).Methods(http.MethodGet, http.MethodOptions)
	router.HandleFunc("/user/{id}", crud.UserUpdate).Methods(http.MethodPut, http.MethodOptions)
	router.HandleFunc("/user/{id}", crud.UserDelete).Methods(http.MethodDelete, http.MethodOptions)
}

