package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jessevdk/go-flags"
	"log"
	"net/http"
	"proyect/handler"
)

type settings struct {
	Port string `short:"p" long:"port" default:"3001" description:"http server port"`
}

func main() {
	log.Print("server starting \n")

	var st settings
	parser := flags.NewParser(&st, flags.Default)

	_, err := parser.Parse()
	if errType, ok := err.(*flags.Error); ok && errType.Type == flags.ErrHelp {
		return
	}

	router := mux.NewRouter()
	router.HandleFunc("/v1/products/list", handler.GetAllProducts).Methods(http.MethodGet)
	router.HandleFunc("/v1/products/{id}", handler.GetProducts).Methods(http.MethodGet)

	err = http.ListenAndServe(fmt.Sprintf(":%s", st.Port), router)

	fmt.Printf("service terminated %s", err)
}