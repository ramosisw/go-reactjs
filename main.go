package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/jramos/go-reactjs/frontend"
	"github.com/jramos/go-reactjs/handler"
	"github.com/jramos/go-reactjs/models"

	assetfs "github.com/elazarl/go-bindata-assetfs"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/sqlite" //sqlite3
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

	router.Use(commonMiddleware)
	router.HandleFunc("/todo", handler.PostTodo).Methods("POST")
	router.HandleFunc("/todo", handler.GetTodos).Methods("GET")
	router.HandleFunc("/todo/{id}", handler.PutTodo).Methods("PUT")
	router.HandleFunc("/todo/{id}", handler.GetTodo).Methods("GET")

	router.PathPrefix("/").Handler(
		http.FileServer(assetFS()),
	)

	fmt.Println("GO REST server running on http://localhost ")
	log.Fatal(http.ListenAndServe(":80", router))
}

func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
