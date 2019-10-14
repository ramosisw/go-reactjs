package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/jramos/golang-reactjs/models"
)

//CreateEvent create
func CreateEvent(w http.ResponseWriter, r *http.Request) {
	var newEvent models.Event
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Kindly enter data with the event title and description only in order to update")
	}
	fmt.Fprintf(w, "Kindly enter data with the event title and description only in order to update")

	json.Unmarshal(reqBody, &newEvent)
	// models.Events = append(models.Events, newEvent)
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(newEvent)
}

//GetOneEvent GetOneEvent
func GetOneEvent(w http.ResponseWriter, r *http.Request) {
	// eventID, _ := strconv.Atoi(mux.Vars(r)["id"])

	// for _, singleEvent := range models.Events {
	// 	if singleEvent.Id == eventID {
	// 		json.NewEncoder(w).Encode(singleEvent)
	// 	} else {
	// 		var newEvent models.Event
	// 		json.NewEncoder(w).Encode(newEvent)
	// 	}
	// }
}

//GetAllEvents GetAllEvents
func GetAllEvents(w http.ResponseWriter, r *http.Request) {
	events, _ := models.ListEvents()
	fmt.Println("Endpoint Hit: returnAllBookings")
	json.NewEncoder(w).Encode(events)
}

//HomeLink HomeLink
func HomeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: HomeLink")
	fmt.Fprintf(w, "Welcome home!")
}
