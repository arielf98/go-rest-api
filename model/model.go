package model

import (
	"encoding/json"
	"fmt"
	"go-rest-api/dummy"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

func HomeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Halo dunia")
}

func CreateEvent(w http.ResponseWriter, r *http.Request) {
	var newEvent = dummy.Events
	reqBody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		fmt.Fprintf(w, "Kindly enter data with the event title and description only in order to update")
	}

	json.Unmarshal(reqBody, &newEvent)
	dummy.Events = append(dummy.Events, newEvent...)
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(newEvent)
}

func GetOneEvent(w http.ResponseWriter, r *http.Request) {
	eventID := mux.Vars(r)["id"]
	for _, singleEvent := range dummy.Events {
		if singleEvent.ID == eventID {
			json.NewEncoder(w).Encode(singleEvent)
		}
	}
}

func GetAllEvents(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(dummy.Events)
}

func UpdateEvent(w http.ResponseWriter, r *http.Request) {
	eventID := mux.Vars(r)["id"]
	var updateEvent dummy.Event

	reqBody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		fmt.Fprintf(w, "Kindly enter data with the event title and description only in order to update")

	}

	json.Unmarshal(reqBody, &updateEvent)

	for i, singleEvent := range dummy.Events {
		if singleEvent.ID == eventID {
			singleEvent.Title = updateEvent.Title
			singleEvent.Description = updateEvent.Description
			dummy.Events = append(dummy.Events[:i], singleEvent)
			json.NewEncoder(w).Encode(singleEvent)
		}
	}
}

func DeleteEvent(w http.ResponseWriter, r *http.Request) {
	eventID := mux.Vars(r)["id"]

	for i, singleEvent := range dummy.Events {
		if singleEvent.ID == eventID {
			dummy.Events = append(dummy.Events[:i], dummy.Events[i+1:]...)
			fmt.Fprintf(w, "The event with ID %v has been deleted successfully", eventID)
		}
	}
}
