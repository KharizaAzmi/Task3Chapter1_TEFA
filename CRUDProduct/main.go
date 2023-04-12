package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	route := mux.NewRouter()
	s := route.PathPrefix("/").Subrouter() //Base Path

	//Routes
	s.HandleFunc("/addProduct", addProduct).Methods("POST")
	s.HandleFunc("/getAllProduct", getAllProduct).Methods("GET")
	s.HandleFunc("/getProductProfile", getProductProfile).Methods("POST")
	s.HandleFunc("/updateProduct", updateProduct).Methods("PUT")
	s.HandleFunc("/deleteProduct/{id}", deleteProduct).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", s)) // Run Server
}