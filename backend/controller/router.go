package controller

import (
	"net/http"

	"github.com/gorilla/mux"
)

func Register() *mux.Router {
	router := mux.NewRouter()

	// routes http
	// router.HandleFunc("/ping", ping).Methods(http.MethodGet)
	router.HandleFunc("/coupons", GetAllCoupons).Methods(http.MethodGet)
	// router.HandleFunc("/product/{id}", handlers.GetProduct).Methods(http.MethodGet)
	return router
}
