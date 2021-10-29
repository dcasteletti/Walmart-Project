package handler

import (
	"github.com/gorilla/mux"
	"net/http"
	"proyect/server"
	"proyect/service"
	"strconv"
)

func GetAllProducts(w http.ResponseWriter, r *http.Request) {
	response := server.NewResponse(w)

	products, err := service.GetAllProducts()
	if err != nil {
		return
	}

	response.JSON(http.StatusOK, products)
}

func GetProducts(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	isID := searchByID(id)
	 response := server.NewResponse(w)

	/*
	if !validateParams(id) && !isID {
		return
	}

	 */
	
	products, err := service.GetProducts(id, isID)
	if err != nil {
		response.Error(http.StatusBadRequest, "ERROR")

		return
	}

	if products == nil || len(products.Product) < 1 {
		response.Error(http.StatusNotFound, "ERROR")

		return
	}

	response.JSON(http.StatusOK, products)
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