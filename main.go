package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"proyect/handler"
)

func main()  {
	router := mux.NewRouter()
	router.HandleFunc("/products/list", handler.GetAllProducts).Methods(http.MethodGet)
	router.HandleFunc("/products/{id}", handler.GetProducts).Methods(http.MethodGet)

	http.ListenAndServe(":12345", router)
}