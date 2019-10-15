package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/jramos/go-reactjs/frontend"
	"github.com/jramos/go-reactjs/handler"
	"github.com/jramos/go-reactjs/models"

	"github.com/elazarl/go-bindata-assetfs"
	"github.com/gorilla/mux"
)

const assetPrefix = "build"

func assetFS() *assetfs.AssetFS {
	return &assetfs.AssetFS{
		Asset:     frontend.Asset,
		AssetDir:  frontend.AssetDir,
		AssetInfo: frontend.AssetInfo,
		Prefix:    assetPrefix,
	}
}

func main() {
	models.Init()
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/event", handler.CreateEvent).Methods("POST")
	router.HandleFunc("/events", handler.GetAllEvents).Methods("GET")
	router.HandleFunc("/event/{id}", handler.GetOneEvent).Methods("GET")

	router.PathPrefix("/").Handler(
		http.FileServer(assetFS()),
	)
	// http.Handle("/", http.FileServer(assetFS()))

	// fmt.Printf("Frontend loaded from: [%v]\n", pathFrontend)
	fmt.Println("GO REST server running on http://localhost ")
	log.Fatal(http.ListenAndServe(":80", router))
}
