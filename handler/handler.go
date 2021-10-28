package handler

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"proyect/service"
	"strconv"
)

func GetAllProducts(w http.ResponseWriter, r *http.Request) {
	products, err := service.GetAllProducts()
	if err != nil {
		return
	}

	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(products)
}

func GetProducts(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	isID := searchByID(id)

	/*
	if !validateParams(id) && !isID {
		return
	}

	 */
	
	products, err := service.GetProducts(id, isID)
	if err != nil {
		return
	}

	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(products)
}

func searchByID(value string) bool {
	_, err := strconv.Atoi(value)
	if err != nil {
		return false
	}

	return true
}

func validateParams(id string) bool {
	if len(id) > 3 {
		return true
	}

	return false
}