package main

import (
	"go-rest-api/model"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", model.HomeLink)
	router.HandleFunc("/event", model.CreateEvent).Methods("POST")
	router.HandleFunc("/get_events", model.GetAllEvents).Methods("GET")
	router.HandleFunc("/event/{id}", model.GetOneEvent).Methods("GET")
	router.HandleFunc("/event/{id}", model.UpdateEvent).Methods("PATCH")
	router.HandleFunc("/event/{id}", model.DeleteEvent).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", router))
}
