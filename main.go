package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/jramos/golang-reactjs/handler"
	"github.com/jramos/golang-reactjs/models"

	"github.com/gorilla/mux"
)

func main() {
	models.Init()
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/event", handler.CreateEvent).Methods("POST")
	router.HandleFunc("/events", handler.GetAllEvents).Methods("GET")
	router.HandleFunc("/event/{id}", handler.GetOneEvent).Methods("GET")

	// Server frontend
	var pathFrontend string = os.Getenv("PATH_FRONTEND")
	if pathFrontend == "" {
		pathFrontend = "./static"
	}
	router.PathPrefix("/").Handler(http.FileServer(http.Dir(pathFrontend)))
	fmt.Printf("Frontend loaded from: [%v]\n", pathFrontend)
	fmt.Println("GO REST server running on http://localhost ")
	log.Fatal(http.ListenAndServe(":80", router))
}
